package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	UGX = "UGX"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
