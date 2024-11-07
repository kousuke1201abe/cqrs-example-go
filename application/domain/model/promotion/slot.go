package promotion

import "errors"

type Slot struct {
	remainingAmount int
}

func NewSlot(remainingAmount int) (Slot, error) {
	if remainingAmount < 0 {
		return Slot{}, errors.New("Slot remaining amount should be over 0")
	}
	if remainingAmount > 100000 {
		return Slot{}, errors.New("Slot remaining amount should be under 100000")
	}

	return Slot{remainingAmount: remainingAmount}, nil
}

func (s Slot) isFull() bool {
	return s.remainingAmount == 0
}
