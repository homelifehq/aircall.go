package filter

import (
	"fmt"
	"time"
)

/*func validateUnixDate(key string, stringUnix string) error {
	if _, err := strconv.ParseInt(stringUnix, 10, 64); err != nil {
		stringErr := "validation for %s failed, %s is not a valid unix timestamp"
		return fmt.Errorf(stringErr, key, stringUnix)
	}

	return nil
}*/

func validateUnixDate(key string, timeUnix int64) error {
	tm := time.Unix(timeUnix, 0)
	if len(tm.String()) == 0 {
		stringErr := "validation for %s failed, %d is not a valid unix timestamp"
		return fmt.Errorf(stringErr, key, timeUnix)
	}

	return nil
}

// validate returns an error if prerequisites failed the validation
func (s *Params) validate() error {
	if s.To != nil {
		if err := validateUnixDate("to", *s.To); err != nil {
			return err
		}
	}

	if s.From != nil {
		if err := validateUnixDate("from", *s.From); err != nil {
			return err
		}
	}

	if s.Order != nil {
		order := *s.Order
		if !order.IsValid() {
			return fmt.Errorf("%s is not a valid sort order", order)
		}
	}

	return nil
}
