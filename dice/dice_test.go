package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoll(t *testing.T) {
	values := make(map[int]int)
	for x := 0; x < 1000; x++ {
		value := Dice{}.Roll()
		values[value]++
	}

	assert.Empty(t, values[0])
	assert.Empty(t, values[1])
	assert.Empty(t, values[13])
	assert.Empty(t, values[14])
}
