package promotion

import (
	"errors"
	"html"
	"unicode/utf8"
)

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	if value == "" {
		return Name{}, errors.New("name should not be blank")
	}

	if utf8.RuneCountInString(value) > 50 {
		return Name{}, errors.New("name length should not be over 50")
	}

	return Name{value: html.EscapeString(value)}, nil
}
