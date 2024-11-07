package model

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

const layout = time.RFC3339

func UnmarshalISO8601DateTime(v any) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		return time.Time{}, errors.New("invalid datetime format")
	}

	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}, errors.New("invalid datetime format")
	}

	return t.In(jst).Truncate(time.Second), nil
}

func MarshalISO8601DateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(t.In(jst).Format(layout)))
	})
}
