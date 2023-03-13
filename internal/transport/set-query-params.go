package transport

import (
	"github.com/homelifehq/aircall.go/api/filter/v1"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// SetQueryParams
func SetQueryParams(rawURL *string, filterParams *filter.Params) error {
	if filterParams == nil {
		return nil
	}

	URL, parseErr := url.Parse(*rawURL)
	if parseErr != nil {
		return parseErr
	}

	params := URL.Query()
	reflectionType := reflect.TypeOf(*filterParams)
	reflectionValue := reflect.ValueOf(*filterParams)

	for i := range make([]string, reflectionType.NumField()) {
		attribute := reflectionType.Field(i).Name
		field := reflectionValue.FieldByName(attribute)
		if !field.IsNil() {
			ptr := unsafe.Pointer(field.Pointer())
			switch field.Type().String() {
			case "*int64":
				{
					params.Add(strings.ToLower(attribute), strconv.Itoa(int(*(*int64)(ptr))))
				}
			case "*filter.OrderDirection":
				params.Add(strings.ToLower(attribute), string(*(*filter.OrderDirection)(ptr)))
			}
		}
	}

	URL.RawQuery = params.Encode()
	*rawURL = URL.String()

	return nil
}
