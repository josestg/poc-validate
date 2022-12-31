package validate

import "encoding/json"

type Validator[T any] func(T) error

func (f Validator[T]) Evaluate(n T) error { return f(n) }

func mergeValidator[T any](validators ...Validator[T]) Validator[T] {
	return func(v T) error {
		for _, validator := range validators {
			if err := validator(v); err != nil {
				return err
			}
		}
		return nil
	}
}

type Composer[T any] func(f Validator[T]) Validator[T]

func compose[T any](composer Composer[T], next Validator[T]) Composer[T] {
	return func(validator Validator[T]) Validator[T] {
		return mergeValidator(composer(validator), next)
	}
}

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

func nop[T any]() Validator[T] { return func(_ T) error { return nil } }
