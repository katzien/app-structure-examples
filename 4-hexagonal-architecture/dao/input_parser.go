package dao

import (
	"github.com/katzien/app-structure-examples/4-hexagonal-architecture/domain"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParsePuzzleInputFile() ([]domain.Elf, error) {
	bytes, err := os.ReadFile("./puzzle_input")
	if err != nil {
		return nil, err
	}

	return parseRawInput(string(bytes))
}

func parseRawInput(r string) ([]domain.Elf, error) {
	lines := strings.Split(r, "\n")
	elfIndex := 1
	calorieCount := 0

	list := []domain.Elf{}
	for _, l := range lines {
		if len(l) == 0 {
			e := domain.Elf{
				Name: elfIndex,
				Load: domain.Calories(calorieCount),
			}
			list = append(list, e)

			calorieCount = 0
			elfIndex++
			continue
		}

		c, err := strconv.Atoi(l)
		if err != nil {
			log.Printf("can't parse this calories line: %v", l)
			return nil, err
		}
		calorieCount += c
	}

	list = append(list, domain.Elf{
		Name: elfIndex,
		Load: domain.Calories(calorieCount),
	})

	return list, nil
}
