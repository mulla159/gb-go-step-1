package fibb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name     string
	input    int
	expected int
}{
	{
		name:     "case - 0",
		input:    0,
		expected: 0,
	},
	{
		name:     "case - 1",
		input:    1,
		expected: 1,
	},
	{
		name:     "case - 2",
		input:    2,
		expected: 1,
	},
	{
		name:     "case - 30",
		input:    30,
		expected: 832040,
	},
}

func TestGetFibonacciByCache(t *testing.T) {
	for _, cse := range testCases {
		t.Run(cse.name, func(t *testing.T) {
			fibb := GetFibonacciByCache(cse.input)
			assert.Equal(t, cse.expected, fibb)
		})
	}
}

func TestGetFibonacciRecur(t *testing.T) {
	for _, cse := range testCases {
		t.Run(cse.name, func(t *testing.T) {
			fibb := GetFibonacciRecur(cse.input)
			assert.Equal(t, cse.expected, fibb)
		})
	}
}
