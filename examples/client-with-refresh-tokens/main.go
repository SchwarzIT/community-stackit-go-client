package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	stackit "github.com/SchwarzIT/community-stackit-go-client"
	serviceaccounts "github.com/SchwarzIT/community-stackit-go-client/pkg/services/service-accounts/v2.0"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var userToken = ``
var serviceAccountEmail = `odj-sa-siwk-3gxhsy1@sa.stackit.cloud`
var projectID = `7deb0137-2539-4571-bb0d-18e9cf7b6b77`

func main() {
	ctx := context.Background()
	// c, _ := stackit.NewClientWithConfig(ctx, stackit.Config{
	// 	ServiceAccountEmail: serviceAccountEmail,
	// 	ServiceAccountToken: userToken,
	// })

	// b, err := os.ReadFile("public_key.pem")
	// if err != nil {
	// 	panic(err)
	// }
	// pk := string(b)
	// res, err := c.ServiceAccounts.CreateKeys(ctx, projectID, types.Email(c.GetConfig().ServiceAccountEmail), serviceaccounts.CreateKeysJSONRequestBody{
	// 	PublicKey: &pk,
	// })

	// if err = validate.Response(res, err, "JSON201"); err != nil {
	// 	panic(err)
	// }

	// b, err = json.MarshalIndent(res.JSON201, "", "  ")
	// if err != nil {
	// 	panic(err)
	// }

	// if err := os.WriteFile("sa_key.json", b, 0644); err != nil {
	// 	panic(err)
	// }

	c := stackit.NewClient(ctx)
	b, err := os.ReadFile("sa_key.json")
	if err != nil {
		panic(err)
	}

	var v serviceaccounts.ServiceAccountKeyPrivateResponse
	if err := json.Unmarshal(b, &v); err != nil {
		panic(err)
	}

	b, err = os.ReadFile("private_key.pem")
	if err != nil {
		panic(err)
	}

	selfSignedJWT, err := GenerateSelfSignedJWT(v, b)
	if err != nil {
		panic(err)
	}
	fmt.Println("Generated a self signed JWT:")
	fmt.Println(selfSignedJWT)
	fmt.Println("")

	res, err := c.ServiceAccounts.CreateTokenWithFormdataBody(ctx, serviceaccounts.TokenRequestBody{
		Assertion: &selfSignedJWT,
		GrantType: "urn:ietf:params:oauth:grant-type:jwt-bearer",
	})
	if err != nil {
		panic(err)
	}
	if res.JSON200 == nil {
		panic("empty JSON200")
	}

	b, err = json.MarshalIndent(*res.JSON200, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

type Credentials struct {
	KeyId      string `json:"kid"`
	Issuer     string `json:"iss"`
	Subject    string `json:"sub"`
	Audience   string `json:"aud"`
	PrivateKey string `json:"privateKey"`
}

func GenerateSelfSignedJWT(key serviceaccounts.ServiceAccountKeyPrivateResponse, privateKeyBytes []byte) (string, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("Error parsing private key: %v", err)
	}
	claims := jwt.MapClaims{
		"iss": key.Credentials.Iss,
		"sub": key.Credentials.Sub,
		"jti": uuid.New(),
		"aud": key.Credentials.Aud,
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	token.Header["kid"] = key.Credentials.Kid
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
