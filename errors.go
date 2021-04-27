package httpin

import (
	"fmt"
	"reflect"
)

type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e UnsupportedTypeError) Error() string {
	return "httpin: unsupported type: " + e.Type.String()
}

type InvalidField struct {
	// Field is the name of the field.
	Field string `json:"field"`

	// Source is the tag indicates where to extract the value of the field.
	// e.g. query.name, header.bearer_token, body.file
	Source string `json:"source"`

	// Value of the source, who caused the error.
	Value interface{} `json:"value"`

	// InternalError
	InternalError error `json:"error"`
}

func (f *InvalidField) Error() string {
	return fmt.Sprintf("httpin: invalid field %q: %v", f.Field, f.InternalError)
}

func (f *InvalidField) Unwarp() error {
	return f.InternalError
}
