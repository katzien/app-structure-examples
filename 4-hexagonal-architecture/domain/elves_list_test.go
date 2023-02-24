package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMostCalories(t *testing.T) {
	t.Parallel()

	input := []Elf{
		{Name: 0, Load: Calories(6000)},
		{Name: 1, Load: Calories(4000)},
		{Name: 2, Load: Calories(11000)},
		{Name: 3, Load: Calories(24000)},
		{Name: 4, Load: Calories(10000)},
	}

	elf := FindMostCalories(input)
	assert.Equal(t, 3, elf.Name)
	assert.Equal(t, Calories(24000), elf.Load)
}
