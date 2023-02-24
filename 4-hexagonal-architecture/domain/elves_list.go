package domain

// FindMostCalories takes a list of all elves with their calories
// and returns the elf which carries the most calories.
func FindMostCalories(list []Elf) Elf {
	var elf Elf
	mostCalories := Calories(0)

	for _, e := range list {
		if e.Load > mostCalories {
			mostCalories = e.Load
			elf = e
		}
	}

	return elf
}
