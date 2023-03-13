package authentication

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type BasicAuth struct {
	ID    string
	Token string
}

func (auth BasicAuth) baseEncode() Credential {
	base64Representation := base64.StdEncoding.EncodeToString(
		[]byte(auth.ID + ":" + auth.Token),
	)

	return []byte(base64Representation)
}

func (auth BasicAuth) SetAuth(r *http.Request) {
	authorization := fmt.Sprintf("Basic %s", auth.baseEncode())
	r.Header.Add("Authorization", authorization)
}

func (auth BasicAuth) String() string {
	return fmt.Sprintf("%s:%s", auth.ID, auth.Token)
}

func NewBasicAuth(id string, token string) BasicAuth {
	return BasicAuth{
		ID:    id,
		Token: token,
	}
}
