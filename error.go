package bitly

import (
	"encoding/json"
)

func unmarshalError(d *json.Decoder) *Error {
	err := new(Error)
	_ = d.Decode(err)
	return err
}

type Error struct {
	Description string        `json:"description"`
	Errors      []*FieldError `json:"errors"`
	Message     string        `json:"message"`
	Resource    string        `json:"resource"`
}

func (e *Error) Error() string {
	v, _ := json.MarshalIndent(e, "", "\t")
	return string(v)
}

type FieldError struct {
	ErrorCode string `json:"error_code"`
	Field     string `json:"field"`
	Message   string `json:"message"`
}
