package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeededRoll(t *testing.T) {
	dice := NewSeededDice()
	values := make(map[int]int)
	for x := 0; x < 10000; x++ {
		value := dice.Roll()
		values[value]++
	}

	assert.Empty(t, values[0])
	assert.Empty(t, values[1])
	assert.NotEmpty(t, values[2])
	assert.NotEmpty(t, values[3])
	assert.NotEmpty(t, values[4])
	assert.NotEmpty(t, values[5])
	assert.NotEmpty(t, values[6])
	assert.NotEmpty(t, values[7])
	assert.NotEmpty(t, values[8])
	assert.NotEmpty(t, values[9])
	assert.NotEmpty(t, values[10])
	assert.NotEmpty(t, values[11])
	assert.NotEmpty(t, values[12])
	assert.Empty(t, values[13])
	assert.Empty(t, values[14])
}

func TestRandomRoll(t *testing.T) {
	dice := RandomDice{}
	values := make(map[int]int)
	for x := 0; x < 10000; x++ {
		value := dice.Roll()
		values[value]++
	}

	assert.Empty(t, values[0])
	assert.Empty(t, values[1])
	assert.NotEmpty(t, values[2])
	assert.NotEmpty(t, values[3])
	assert.NotEmpty(t, values[4])
	assert.NotEmpty(t, values[5])
	assert.NotEmpty(t, values[6])
	assert.NotEmpty(t, values[7])
	assert.NotEmpty(t, values[8])
	assert.NotEmpty(t, values[9])
	assert.NotEmpty(t, values[10])
	assert.NotEmpty(t, values[11])
	assert.NotEmpty(t, values[12])
	assert.Empty(t, values[13])
	assert.Empty(t, values[14])
}
