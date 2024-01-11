package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Max_float(t *testing.T) {
	assert.Equal(t, 2.23, Max(2.1, 2.23))
}

func Test_Min_int(t *testing.T) {
	assert.Equal(t, -23, Min(-2, -23))
}

func Test_Max_rune(t *testing.T) {
	var a rune = 1
	var b rune = -1
	assert.Equal(t, a, Max(a, b))
}

func Test_Max_byte(t *testing.T) {
	var a byte = 1
	var b byte = 0
	assert.Equal(t, a, Max(a, b))
}

func Test_Min_uintptr(t *testing.T) {
	var a uintptr = 33333
	var b uintptr = 1
	assert.Equal(t, b, Min(a, b))
}
