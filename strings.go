package validate

import (
	"fmt"
	"regexp"
	"strings"
)

type StringComposer Composer[string]

func String() StringComposer { return Identity[Validator[string]] }

func (f StringComposer) and(next Validator[string]) StringComposer {
	return StringComposer(compose(Composer[string](f), next))
}

func (f StringComposer) Compose() Validator[string]          { return f(nop[string]()) }
func (f StringComposer) NotBlank() StringComposer            { return f.and(notBlank()) }
func (f StringComposer) NotBlankTrim() StringComposer        { return f.and(notBlankTrim()) }
func (f StringComposer) Email() StringComposer               { return f.and(email()) }
func (f StringComposer) Len(min int, max int) StringComposer { return f.and(strLen(min, max)) }

func notBlank() Validator[string] {
	return func(s string) error {
		if "" == s {
			return NewError("string_not_blank", "must not be blank", nil)
		}

		return nil
	}
}

func notBlankTrim() Validator[string] {
	return func(s string) error {
		return notBlank().Evaluate(strings.TrimSpace(s))
	}
}

func strLen(min int, max int) Validator[string] {
	return func(s string) error {
		args := map[string]any{
			"min": min,
			"max": max,
			"len": len(s),
		}

		// skip min check if min is negative.
		if min > 0 && min > len(s) {
			return NewError("string_len", fmt.Sprintf("must be at least %d characters", min), args)
		}

		// skip max check if max is negative.
		if max > 0 && max < len(s) {
			return NewError("string_len", fmt.Sprintf("must be at most %d characters", max), args)
		}

		return nil
	}
}

var regexEmail = regexp.MustCompile("^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")

func email() Validator[string] {
	return func(s string) error {
		if !regexEmail.MatchString(s) {
			return NewError("string_email", "must be a valid email address", nil)
		}

		return nil
	}
}
