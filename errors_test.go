package convsus_test

import (
	"testing"

	"github.com/acdgbrasil/convsus"
	"github.com/stretchr/testify/require"
)

// Since ErrInvalidError cannot be declared outside constructor due to generating a cycle. Here is a constant with its module and identifier
var ErrInvalidError = convsus.NewError(1, 1, "")

func TestNewError(t *testing.T) {
	testCases := []struct {
		purpose     string
		module      int
		id          int
		message     string
		extra       []interface{}
		expectedErr error
	}{
		{
			purpose:     "Should work",
			module:      999,
			id:          999,
			expectedErr: convsus.NewError(999, 999, ""),
		},
		{
			purpose:     "Invalid module",
			module:      0,
			id:          1,
			expectedErr: ErrInvalidError,
		},
		{
			purpose:     "Invalid identifier",
			module:      1,
			id:          0,
			expectedErr: ErrInvalidError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.purpose, func(t *testing.T) {
			require.NotNil(t, tc.expectedErr, "should have throwed an error, but did not")
			e, ok := tc.expectedErr.(*convsus.Error)
			require.True(t, ok, "expected error should be a system error, not generic")
			require.True(t, e.Equals(convsus.NewError(convsus.ERROR_MODULE(tc.module), convsus.ERROR_ID(tc.id), tc.message)), "errors should be equal, but weren't")
		})
	}
}

func TestErrorFormat(t *testing.T) {
	testCases := []struct {
		purpose  string
		message  string
		extra    []any
		expected string
	}{
		{
			purpose:  "Should work with single string variable",
			message:  "Field: %s",
			extra:    []any{"value"},
			expected: "Field: value",
		},
		{
			purpose:  "Should replace empty format symbol",
			message:  "Field: %s",
			extra:    nil,
			expected: "Field: [ missing ]",
		},
		{
			purpose:  "Should correctly apply different types",
			message:  "Field: %s, Field %d, Field %.2f",
			extra:    []any{"string", 0, 0.2},
			expected: "Field: string, Field 0, Field 0.20",
		},
		{
			purpose:  "Should ignore extra fields",
			message:  "Field",
			extra:    []any{"string", 0},
			expected: "Field",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.purpose, func(t *testing.T) {
			// Create error with message provided
			require.Equal(t, "[ERROR-099099] "+tc.expected, convsus.NewError(99, 99, tc.message, tc.extra...).Error())
		})
	}
}
