package exerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	errs := []error{
		errors.New("0"),
		errors.New("1"),
		errors.New("2"),
		errors.New("3"),
	}

	testCases := []struct {
		desc      string
		expected  []error
		arguments []error
	}{
		{
			desc:      "empty",
			expected:  nil,
			arguments: []error{},
		},
		{
			desc:      "one",
			expected:  []error{errs[0]},
			arguments: []error{errs[0]},
		},
		{
			desc:      "multiple",
			expected:  errs,
			arguments: errs,
		},
		{
			desc:      "nesting",
			expected:  errs,
			arguments: []error{errors.Join(errs[0], errors.Join(errs[1], errs[2], errors.Join(errs[3])))},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expected, Flatten(tC.arguments...))
		})
	}
}
