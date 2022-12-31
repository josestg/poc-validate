package validate

import "encoding/json"

type Error struct {
	Constraint string
	Message    string
	Args       map[string]any
}

func NewError(constraint string, message string, args map[string]any) *Error {
	return &Error{Constraint: constraint, Message: message, Args: args}
}

func (e *Error) Error() string { return stringify(e) }

type MapError map[string]*Error

func (e MapError) Error() string { return stringify(e) }

func stringify(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

func Identity[F any](f F) F { return f }

func nop[T any]() func(T) error { return func(_ T) error { return nil } }