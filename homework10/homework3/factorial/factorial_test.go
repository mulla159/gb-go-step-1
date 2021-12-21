package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	var number, expected float64 = 9, 362880

	f := Factorial(number)

	assert.Equal(t, expected, f)
}
