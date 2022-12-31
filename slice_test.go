package validate_test

import (
	"strings"
	"testing"
	"validate"
)

func TestSliceOf(t *testing.T) {
	tableTest := []struct {
		name           string
		value          []string
		wantError      bool
		wantConstraint map[string]string
		validator      validate.StringValidator
	}{
		{
			name:      "it should fail when any string is blank",
			value:     []string{"", "bob"},
			wantError: true,
			wantConstraint: map[string]string{
				"0": "string_not_blank",
			},
			validator: validate.String().NotBlank().Compose(),
		},
		{
			name:      "it should fail when any string is not an email",
			value:     []string{"bob@mail.com", "bob"},
			wantError: true,
			wantConstraint: map[string]string{
				"1": "string_email",
			},
			validator: validate.String().NotBlank().Email().Compose(),
		},
		{
			name:      "when all items have different errors",
			value:     []string{"", "valid@mail.com", strings.Repeat("a", 21), "bob@mail.com", "not@mail"},
			wantError: true,
			wantConstraint: map[string]string{
				"0": "string_not_blank",
				"2": "string_len",
				"4": "string_email",
			},
			validator: validate.String().NotBlank().Len(4, 20).Email().Compose(),
		},
		{
			name:      "when all items have the same error",
			value:     []string{"edo", "bob"},
			wantError: true,
			wantConstraint: map[string]string{
				"0": "string_len",
				"1": "string_len",
			},
			validator: validate.String().Len(4, 10).Compose(),
		},
		{
			name:           "when all items are valid",
			value:          []string{"edo@mail.com", "bob@mail.com"},
			wantError:      false,
			wantConstraint: map[string]string{},
			validator:      validate.String().NotBlank().Len(4, 100).Email().Compose(),
		},
	}

	for _, tt := range tableTest {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := validate.SliceOf(tt.value).With(tt.validator)

			if tt.wantError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}

				switch et := err.(type) {
				default:
					t.Errorf("expected MapError, got %T", et)
				case validate.MapError:
					t.Logf("got error: %s", et.Error())
					for index, constraint := range tt.wantConstraint {
						if et[index].Constraint != constraint {
							t.Errorf("expected constraint %s, got %s", constraint, et[index].Constraint)
						}
					}
				}
			} else {
				if err != nil {
					t.Errorf("expected nil, got %v", err)
				}
			}
		})
	}
}
