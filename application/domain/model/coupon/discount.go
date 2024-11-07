package coupon

import "errors"

type Discount struct {
	amount int
}

func NewDiscount(amount int) (Discount, error) {
	if amount <= 1 {
		return Discount{}, errors.New("discount acount should be over 1")
	}
	if amount > 100000 {
		return Discount{}, errors.New("discount acount should be under 100000")
	}

	return Discount{amount: amount}, nil
}
