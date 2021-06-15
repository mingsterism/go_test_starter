package custom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiply(t *testing.T) {
	result := Multiply(10, 20)
	assert.Equal(t, 200, result)
}
