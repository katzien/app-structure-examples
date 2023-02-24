package dao

import (
	"github.com/katzien/app-structure-examples/4-hexagonal-architecture/domain"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	t.Parallel()

	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	list, err := parseRawInput(input)
	require.NoError(t, err)

	assert.Len(t, list, 5)
	assert.Equal(t, domain.Elf{Name: 1, Load: domain.Calories(6000)}, list[0])
	assert.Equal(t, domain.Elf{Name: 2, Load: domain.Calories(4000)}, list[1])
	assert.Equal(t, domain.Elf{Name: 3, Load: domain.Calories(11000)}, list[2])
	assert.Equal(t, domain.Elf{Name: 4, Load: domain.Calories(24000)}, list[3])
	assert.Equal(t, domain.Elf{Name: 5, Load: domain.Calories(10000)}, list[4])
}
