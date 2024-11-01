package enum

type Currency string

func (c Currency) String() string {
	return string(c)
}

// Currency describes the currency of a transaction
const (
	// NGN
	NGN Currency = "NGN"
)
