package users_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client"
	"github.com/SchwarzIT/community-stackit-go-client/internal/clients"
	u "github.com/SchwarzIT/community-stackit-go-client/pkg/api/v1/users"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
	"github.com/SchwarzIT/community-stackit-go-client/pkg/validate"
)

var (
	skipAcceptanceTestGetUser = true
)

func TestUsersService_ValidateUserOrigin(t *testing.T) {
	type args struct {
		origin string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test fail",
			args: args{
				origin: "something",
			},
			wantErr: true,
		},
		{
			name: "test success",
			args: args{
				origin: consts.SCHWARZ_AUTH_ORIGIN,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validate.UserOrigin(tt.args.origin); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUserOrigin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsersService_Get(t *testing.T) {
	c, mux, teardown, _ := client.MockServer()
	defer teardown()
	users := u.New(c)

	want := u.User{
		Email:          "some@one.com",
		Origin:         consts.SCHWARZ_AUTH_ORIGIN,
		UUID:           "some-id",
		OrganizationID: consts.SCHWARZ_ORGANIZATION_ID,
	}

	mux.HandleFunc("/ucp-shadow-user-management/v1/createcuaashadowuser/user",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			b, err := json.Marshal(u.ShadowUsersResBody{
				Username: want.Email,
				UUID:     want.UUID,
				Origin:   want.Origin,
			})
			if err != nil {
				log.Fatalf("json response marshal: %v", err)
			}
			fmt.Fprint(w, string(b))
		})

	got, err := users.Get(context.Background(), want.Email, want.Origin)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestAccUsersService_Get(t *testing.T) {
	if skipAcceptanceTestGetUser {
		t.Skip()
	}

	ac, err := clients.LocalAuthClient()
	if err != nil {
		t.Error(err)
	}

	res, err := ac.GetToken(context.Background())
	if err != nil {
		t.Fatalf("failed to get token: %v", err)
	}

	c, err := clients.LocalClient()
	if err != nil {
		t.Error(err)
	}
	c.SetToken(res.AccessToken)
	users := u.New(c)
	user, err := users.Get(context.Background(), "deangili.oren@mail.schwarz", "schwarz-federation")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(user.Email, user.UUID)
}
