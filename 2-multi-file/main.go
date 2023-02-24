package multi_file // change this to package main to be able to run main()

import (
	"fmt"
	"os"
)

// Here we've added some custom types to help the functions in different files "communicate" their intent better.
// You could of course use the plain types (int, map) instead, but now that the functions live in different files
// it might be easier to remember what all the return values are by using custom types.
type (
	elf               int
	calories          int
	caloriesByElfList map[elf]calories
)

func main() {
	list, err := os.ReadFile("./puzzle_input")
	if err != nil {
		panic(err)
	}

	input := parseInput(string(list))

	e, c := findMostCalories(input)

	fmt.Printf("Elf %d carries %d calories\n", e, c)
}
