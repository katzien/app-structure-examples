package multi_file // change this to package main to be able to run main()

import (
	"strconv"
	"strings"
)

func parseInput(raw string) caloriesByElfList {
	input := caloriesByElfList{}

	lines := strings.Split(raw, "\n")
	elfIndex := 0
	calorieCount := 0

	for _, l := range lines {
		if len(l) == 0 {
			input[elf(elfIndex)] = calories(calorieCount)
			calorieCount = 0
			elfIndex++
			continue
		}

		c, err := strconv.Atoi(l)
		if err != nil {
			panic("can't parse the calories!")
		}
		calorieCount += c
	}

	input[elf(elfIndex)] = calories(calorieCount)

	return input
}
