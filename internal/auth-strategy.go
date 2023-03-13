package internal

import (
	"net/http"
)

type AuthStrategy interface {
	String() string
	SetAuth(r *http.Request)
}
