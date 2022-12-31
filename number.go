package validate

import "fmt"

type Number interface {
	Integers | Floats
}

func minimum[T Number](constraint string, min T) func(T) error {
	return func(n T) error {
		args := map[string]any{
			"min": min,
			"val": n,
		}

		if n < min {
			return NewError(constraint, fmt.Sprintf("must be greater than or equal to %v", min), args)
		}

		return nil
	}
}

func maximum[T Number](constraint string, max T) func(T) error {
	return func(n T) error {
		args := map[string]any{
			"max": max,
			"val": n,
		}

		if max < n {
			return NewError(constraint, fmt.Sprintf("must be less than or equal to %v", max), args)
		}

		return nil
	}
}

func choose[T Number](constraint string, choices []T) func(T) error {
	return func(n T) error {
		args := map[string]any{
			"choices": choices,
			"val":     n,
		}

		for _, choice := range choices {
			if choice == n {
				return nil
			}
		}

		return NewError(constraint, fmt.Sprintf("must be one of %v", choices), args)
	}
}
