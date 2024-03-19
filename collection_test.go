package exerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errOdd = errors.New("odd")

func mapIteratee(value int, _ int) (int, error) {
	if value%2 == 0 {
		return value * 2, nil
	}

	return value, errOdd
}

var mapTestCases = []struct {
	desc       string
	collection []int
	expected   []int
	err        error
}{
	{
		desc:       "empty",
		collection: []int{},
		expected:   []int{},
		err:        nil,
	},
	{
		desc:       "even",
		collection: []int{2, 4, 6},
		expected:   []int{4, 8, 12},
		err:        nil,
	},
	{
		desc:       "odd",
		collection: []int{1, 3, 5},
		expected:   []int{1, 3, 5},
		err:        errors.Join(errOdd, errOdd, errOdd),
	},
}

func TestMap(t *testing.T) {
	for _, tC := range mapTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := Map(tC.collection, mapIteratee)
			assert.Equal(t, tC.expected, actual)
			assert.Equal(t, tC.err, err)
		})
	}
}

func TestParallelMap(t *testing.T) {
	for _, tC := range mapTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := ParallelMap(tC.collection, mapIteratee)
			assert.Equal(t, tC.expected, actual)
			assert.Equal(t, tC.err, err)
		})
	}
}
