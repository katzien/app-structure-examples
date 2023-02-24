package calculator

// Declaring a few exported types here makes it easier for the rest of the code to interact with our package.
// These types add meaning to the generic types which can make the code more understandable.
// You could of course use the plain int or map types instead!

type (
	Elf               int
	Calories          int
	CaloriesByElfList map[Elf]Calories
)

// FindMostCalories takes the list of calories carried by each elf,
// and returns the elf which carries the most calories and the amount of calories they carry.
func FindMostCalories(input CaloriesByElfList) (Elf, Calories) {
	elfNo := Elf(0)
	mostCalories := Calories(0)

	for k, v := range input {
		if v > mostCalories {
			mostCalories = v
			elfNo = k
		}
	}

	return elfNo + 1, mostCalories
}
