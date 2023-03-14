package v1

type Resource string

const ErrInvalidResource = "invalid resource"

const (
	NumberResource  Resource = "number"
	UserResource    Resource = "user"
	ContactResource Resource = "contact"
	CallResource    Resource = "call"
)

// String returns the string representation of the resource
func (r Resource) String() string {
	return string(r)
}

// IsValid returns true if the resource is valid
func (r Resource) IsValid() bool {
	switch r {
	case NumberResource, UserResource, ContactResource, CallResource:
		return true
	}
	return false
}
