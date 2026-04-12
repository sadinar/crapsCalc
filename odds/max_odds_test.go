package odds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStdOddsMultipliers(t *testing.T) {
	assert.Equal(
		t,
		map[int]int{
			2:  1,
			3:  2,
			4:  3,
			5:  4,
			6:  5,
			8:  5,
			9:  4,
			10: 3,
			11: 2,
			12: 1,
		},
		GetStdMaxOdds(),
	)
}

func TestGet100xMultipliers(t *testing.T) {
	assert.Equal(
		t,
		map[int]int{
			2:  100,
			3:  100,
			4:  100,
			5:  100,
			6:  100,
			8:  100,
			9:  100,
			10: 100,
			11: 100,
			12: 100,
		},
		Get100xMaxOdds(),
	)
}

func TestGet2xMultipliers(t *testing.T) {
	assert.Equal(
		t,
		map[int]int{
			2:  2,
			3:  2,
			4:  2,
			5:  2,
			6:  2,
			8:  2,
			9:  2,
			10: 2,
			11: 2,
			12: 2,
		},
		Get2xMaxOdds(),
	)
}
