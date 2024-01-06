package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Contains_int(t *testing.T) {
	list := []int{1, 2, 3, 4}
	element := 1
	assert.Equal(t, true, Contains(list, element))
}

func Test_Contains_rune(t *testing.T) {
	list := []rune{1, 2, 3, 4}
	var element rune = 1
	assert.Equal(t, true, Contains(list, element))
}
