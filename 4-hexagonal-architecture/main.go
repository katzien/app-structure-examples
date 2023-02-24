package main // change this to package main to be able to run main()

import (
	"fmt"
	"github.com/katzien/app-structure-examples/4-hexagonal-architecture/dao"
	"github.com/katzien/app-structure-examples/4-hexagonal-architecture/domain"
)

func main() {
	list, err := dao.ParsePuzzleInputFile()
	if err != nil {
		panic(err)
	}

	elf := domain.FindMostCalories(list)

	fmt.Printf("Elf %d carries %s.\n", elf.Name, elf.Load.String())
}
