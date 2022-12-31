package validate

import (
	"fmt"
)

type SliceValidator[T any] struct {
	slices []T
}

func SliceOf[T any](s []T) *SliceValidator[T] {
	return &SliceValidator[T]{
		slices: s,
	}
}

func (v *SliceValidator[T]) With(validator func(T) error) error {
	errs := make(MapError, 0)
	for index, item := range v.slices {
		if err := validator(item); err != nil {
			switch et := err.(type) {
			default:
				return err
			case *Error:
				errs[fmt.Sprint(index)] = et
			}
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
