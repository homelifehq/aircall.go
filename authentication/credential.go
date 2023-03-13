package authentication

import (
	"encoding/base64"
	"strings"
)

type Credential []byte

func (credential Credential) Base64Decode() (BasicAuth, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(credential)))
	if _, err := base64.StdEncoding.Decode(dst, credential); err != nil {
		return BasicAuth{}, err
	}

	buffer := strings.Split(string(dst), ":")

	return BasicAuth{
		ID:    buffer[0],
		Token: buffer[1],
	}, nil
}

func (credential Credential) String() string {
	return string(credential)
}
