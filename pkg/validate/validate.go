package validate

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	parse "github.com/karrick/tparse/v2"
	"github.com/oleiade/reflections"
	"github.com/pkg/errors"
)

// Response validates a response interface and error
// if requestError has an error, it is returned
// if resp.Error is defined and not nil, it is returned
// if one of the field namess provided in []checkNullFields are nil, an error is returned
func Response(resp interface{}, requestError error, checkNullFields ...string) error {
	// check request error
	if requestError != nil {
		return requestError
	}

	if err := ResponseObject(resp); err != nil {
		return err
	}

	return Fields(resp, checkNullFields...)
}

// Field validates that given field in the response object is set (not nil)
func Field(resp interface{}, field string) error {
	sl := strings.Split(field, ".")
	res := resp
	for _, f := range sl {
		a, err := reflections.GetField(res, f)
		if err != nil {
			return err
		}
		if a == nil || (reflect.ValueOf(a).Kind() == reflect.Ptr && reflect.ValueOf(a).IsNil()) {
			return fmt.Errorf("field %s in response is nil", field)
		}
		res = a
	}
	return nil
}

// Fields validates that given fields in the response object are set (not nil)
func Fields(resp interface{}, fields ...string) error {
	for _, field := range fields {
		if err := Field(resp, field); err != nil {
			return err
		}
	}
	return nil
}

// ResponseObject validates the response response and checks if it has an Error field
// that's set
func ResponseObject(resp interface{}) error {
	if resp == nil {
		return errors.New("response interface is nil")
	}

	// check Error field exists
	// if not return err (unless the resp is a non-struct, err will be nil)
	if ok, err := reflections.HasField(resp, "Error"); !ok {
		return err
	}

	value, _ := reflections.GetField(resp, "Error")
	if v, ok := value.(error); ok {
		if v != nil {
			return v
		}
	}
	return nil
}

type ResponseInterface interface {
	StatusCode() int
}

// StatusEquals returns true if interface.StatusCode() equals a given http code
// if more than one status code is provided, the function will return true if one of them matches
func StatusEquals(a ResponseInterface, statusCode ...int) bool {
	if a == nil || (reflect.ValueOf(a).Kind() == reflect.Ptr && reflect.ValueOf(a).IsNil()) {
		return false
	}
	for _, code := range statusCode {
		if a.StatusCode() == code {
			return true
		}
	}
	return false
}

// UUID validates a given UUID
func UUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return err
	}
	return nil
}

// ProjectID validates a given project ID
func ProjectID(projectID string) error {
	if err := UUID(projectID); err != nil {
		return errors.Wrap(err, "invalid UUID for project")
	}
	return nil
}

// ProjectName validates a given project name
func ProjectName(name string) error {
	exp := `^[a-zA-Z][ a-zA-Z0-9_-]{1,39}$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(name) {
		return fmt.Errorf("invalid project name. valid name is of: %s", exp)
	}
	return nil
}

// BillingRef validates a given billing reference
func BillingRef(billingRef string) error {
	exp := `^[a-zA-Z][a-zA-Z0-9_-]{1,29}$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(billingRef) {
		return fmt.Errorf("invalid billing reference. valid reference is of: %s", exp)
	}
	return nil
}

// SemVer validates a given version
func SemVer(version string) error {
	if version == "" {
		return errors.New("version is empty")
	}
	exp := `^\d+\.\d+(?:\.\d+)?$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(version) {
		return fmt.Errorf("invalid version. valid version is of: %s", exp)
	}
	return nil
}

// DefaultResponseErrorHandler is the default error handler used to check
// if a giving STACKIT API response returned an error
func DefaultResponseErrorHandler(resp *http.Response) error {
	if resp.StatusCode < 400 {
		return nil
	}

	var b []byte
	if resp.Body != nil {
		b, _ = io.ReadAll(resp.Body)
	}
	return fmt.Errorf(
		"call error:\nHTTP status code: %d\nHTTP status message: %s\nServer response: %s\nURL: %s\nTrace: %s\n",
		resp.StatusCode,
		http.StatusText(resp.StatusCode),
		string(b),
		resp.Request.URL.String(),
		resp.Request.Header.Get("Traceparent"),
	)
}

// ISO8601 Validates that given time is formatted as ISO 8601
func ISO8601(t string) error {
	isoFmt := "2006-01-02T15:04:05.999Z"
	_, err := time.Parse(isoFmt, t)
	if err != nil {
		return errors.Wrap(err, "couldn't parse given time as ISO8601")
	}
	return nil
}

// RFC3339 Validates that given time is formatted as RFC3339
func RFC3339(t string) error {
	_, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return errors.Wrap(err, "couldn't parse given time as RFC3339")
	}
	return nil
}

// Duration validates that a given string can be parsed as duration
// i.e. 5m, 60s, 1h
func Duration(s string) (time.Duration, error) {
	if s == "" {
		return 0, errors.New("can't parse empty string as duration")
	}
	return parse.AbsoluteDuration(time.Now(), s)
}

// ErrorIsOneOf checks checks if a given error message
// has one of the provided sub strings
func ErrorIsOneOf(err error, msgs ...string) bool {
	if err == nil {
		return false
	}
	for _, m := range msgs {
		if strings.Contains(err.Error(), m) {
			return true
		}
	}
	return false
}

// NetworkName validates a given network name
func NetworkName(name string) error {
	exp := `^[A-Za-z0-9]+((-|_|\s|\.)[A-Za-z0-9]+)*$`
	r := regexp.MustCompile(exp)
	if !r.MatchString(name) {
		return fmt.Errorf("invalid network name: %s. valid name must patch the expression: %s", name, exp)
	}
	if len(name) > 63 {
		return fmt.Errorf("invalid network name. The length of the name must contains maximum 63 characters")
	}
	return nil
}

// NameServer validates a given server name
func NameServer(name string) error {
	exp := "((^\\s*((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))\\s*$)|(^\\s*((([0-9a-f]{1,4}:){7}([0-9a-f]{1,4}|:))|(([0-9a-f]{1,4}:){6}(:[0-9a-f]{1,4}|((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3})|:))|(([0-9a-f]{1,4}:){5}(((:[0-9a-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3})|:))|(([0-9a-f]{1,4}:){4}(((:[0-9a-f]{1,4}){1,3})|((:[0-9a-f]{1,4})?:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9a-f]{1,4}:){3}(((:[0-9a-f]{1,4}){1,4})|((:[0-9a-f]{1,4}){0,2}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9a-f]{1,4}:){2}(((:[0-9a-f]{1,4}){1,5})|((:[0-9a-f]{1,4}){0,3}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9a-f]{1,4}:){1}(((:[0-9a-f]{1,4}){1,6})|((:[0-9a-f]{1,4}){0,4}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(:(((:[0-9a-f]{1,4}){1,7})|((:[0-9a-f]{1,4}){0,5}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:)))(%.+)?\\s*$))"
	r := regexp.MustCompile(exp)
	if !r.MatchString(name) {
		return fmt.Errorf("invalid server name: %s. valid name must match the expression: %s", name, exp)
	}
	return nil
}

// PrefixLengthV4 validates a given server name.
func PrefixLengthV4(prefix int64) error {
	if prefix < 22 || prefix > 29 {
		return fmt.Errorf("invalid prefix length: %d. The length must be between 22 and 29", prefix)
	}
	return nil
}

// Prefix validates a given prefix
func Prefix(pr string) error {
	exp := "^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\\/(3[0-2]|2[0-9]|1[0-9]|[0-9]))$|^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))(\\/((1(1[0-9]|2[0-8]))|([0-9][0-9])|([0-9])))?$"
	r := regexp.MustCompile(exp)
	if !r.MatchString(pr) {
		return fmt.Errorf("invalid prefix: %s. valid prefix shuld match expression: %s", pr, exp)
	}
	return nil
}

// PublicIP validates a given ID address
func PublicIP(ip string) error {
	exp := "((^\\s*((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))\\s*$)|(^\\s*((([0-9a-f]{1,4}:){7}([0-9a-f]{1,4}|:))|(([0-9a-f]{1,4}:){6}(:[0-9a-f]{1,4}|((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3})|:))|(([0-9a-f]{1,4}:){5}(((:[0-9a-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3})|:))|(([0-9a-f]{1,4}:){4}(((:[0-9a-f]{1,4}){1,3})|((:[0-9a-f]{1,4})?:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9a-f]{1,4}:){3}(((:[0-9a-f]{1,4}){1,4})|((:[0-9a-f]{1,4}){0,2}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9a-f]{1,4}:){2}(((:[0-9a-f]{1,4}){1,5})|((:[0-9a-f]{1,4}){0,3}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(([0-9a-f]{1,4}:){1}(((:[0-9a-f]{1,4}){1,6})|((:[0-9a-f]{1,4}){0,4}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:))|(:(((:[0-9a-f]{1,4}){1,7})|((:[0-9a-f]{1,4}){0,5}:((25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)(\\.(25[0-5]|2[0-4]\\d|1\\d\\d|[1-9]?\\d)){3}))|:)))(%.+)?\\s*$))"
	r := regexp.MustCompile(exp)
	if !r.MatchString(ip) {
		return fmt.Errorf("invalid public IP: %s. valid IP must match the expression: %s", ip, exp)
	}
	return nil
}

func NetworkID(id string) error {
	exp := "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	r := regexp.MustCompile(exp)
	if !r.MatchString(id) {
		return fmt.Errorf("invalid networkID: %s. valid ID must match the expression: %s", id, exp)
	}
	if len(id) != 36 {
		return fmt.Errorf("invalid networkID: %s. Valid ID must be exactly 36 characters", id)
	}

	return nil
}