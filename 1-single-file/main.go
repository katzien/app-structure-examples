package single_file // change this to package main to be able to run main()

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	list, err := os.ReadFile("./puzzle_input")
	if err != nil {
		panic(err)
	}

	input := parseInput(string(list))

	elf, calories := findMostCalories(input)

	fmt.Printf("Elf %d carries %d calories\n", elf, calories)
}

func parseInput(raw string) map[int]int {
	input := map[int]int{}

	lines := strings.Split(raw, "\n")
	elfIndex := 0
	calorieCount := 0

	for _, l := range lines {
		if len(l) == 0 {
			input[elfIndex] = calorieCount
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

	input[elfIndex] = calorieCount

	return input
}

func findMostCalories(input map[int]int) (int, int) {
	elf := 0
	mostCalories := 0

	for k, v := range input {
		if v > mostCalories {
			mostCalories = v
			elf = k
		}
	}

	return elf + 1, mostCalories
}
