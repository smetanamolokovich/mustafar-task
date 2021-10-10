package data

import (
	"time"

	"github.com/smetanamolokovich/mustafar_task/pkg/validator"
)

type Data struct {
	Key     string    `json:"key"`
	Value   []byte    `json:"value"`
	Expires time.Time `json:"expires"`
}

func ValidateData(v *validator.Validator, data *Data) {
	v.Check(data.Key != "", "key", "must be provided")

	v.Check(data.Value != nil, "value", "must be provided")

	v.Check(!data.Expires.IsZero(), "expires", "must be provided")
	v.Check(int(data.Expires.Nanosecond()) <= int(time.Now().Nanosecond()), "expires", "must not be in the future")
}
