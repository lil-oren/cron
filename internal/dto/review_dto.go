package dto

type (
	RateOfProductModel struct {
		RateCount float64 `db:"rate_count"`
		RateSum   float64 `db:"rate_sum"`
	}
)
