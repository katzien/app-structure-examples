package multi_file // change this to package main to be able to run main()

func findMostCalories(input caloriesByElfList) (elf, calories) {
	elfNo := elf(0)
	mostCalories := calories(0)

	for k, v := range input {
		if v > mostCalories {
			mostCalories = v
			elfNo = k
		}
	}

	return elfNo + 1, mostCalories
}
