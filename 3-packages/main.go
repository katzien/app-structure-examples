package packages // change this to package main to be able to run main()

import (
	"fmt"
	"github.com/katzien/app-structure-examples/3-packages/calculator"
	"os"
)

func main() {
	list, err := os.ReadFile("./puzzle_input")
	if err != nil {
		panic(err)
	}

	input := parseInput(string(list))

	e, c := calculator.FindMostCalories(input)

	fmt.Printf("Elf %d carries %d calories\n", e, c)
}
