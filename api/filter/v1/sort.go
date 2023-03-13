package filter

type OrderDirection string

const (
	Asc  OrderDirection = "asc"
	Desc OrderDirection = "desc"
)

// IsValid return true if value is 'asc' or 'desc'
func (receiver OrderDirection) IsValid() bool {
	isValid := false
	for _, cursor := range []OrderDirection{Asc, Desc} {
		if receiver != cursor {
			continue
		}

		isValid = true
		break
	}

	return isValid
}
