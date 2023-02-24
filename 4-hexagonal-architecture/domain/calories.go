package domain

import "fmt"

type Calories int

func (c Calories) String() string {
	return fmt.Sprintf("%d kcal", c)
}
