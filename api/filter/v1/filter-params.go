package filter

import "fmt"

type Options func(*Params)

type Params struct {
	From  *int64
	To    *int64
	Order *OrderDirection
}

// WithFrom Set a minimal creation date for Users (UNIX timestamp)
func WithFrom(from int64) Options {
	return func(f *Params) {
		f.From = &from
	}
}

// WithTo Set a maximal creation date for Users (UNIX timestamp).
func WithTo(to int64) Options {
	return func(f *Params) {
		f.To = &to
	}
}

// WithSort Reorder entries by created_at. Can be asc or desc.
func WithSort(direction OrderDirection) Options {
	return func(f *Params) {
		f.Order = &direction
	}
}

func (s *Params) apply(opts []Options) {
	for _, opt := range opts {
		opt(s)
	}
}

func NewFilterParams(opts ...Options) *Params {
	params := Params{}
	params.apply(opts)

	if validationErr := params.validate(); validationErr != nil {
		fmt.Println(fmt.Errorf(validationErr.Error()))
	}

	return &params
}

// BACK
/*func NewFilterParams(opts ...Options) (filter *Params, err error) {
	params := Params{}
	params.apply(opts)
	if validationErr := params.validate(); validationErr != nil {
		err = validationErr
		return
	}

	filter = &params

	return
}*/
