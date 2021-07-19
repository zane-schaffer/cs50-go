package programs

import (
	"fmt"

	"github.com/zane-schaffer/cs50/utils"
)

func Cash() {
	fmt.Println("Change: ")
	input := utils.FloatInput()
	var count int
	for input > 0 {
		switch {
		case input >= 0.25:
			input -= 0.25
			count++
		case input >= 0.10:
			input -= 0.10
			count++
		case input >= 0.05:
			input -= 0.05
			count++
		case input >= 0.01:
			input -= 0.01
			count++
		}
	}
	fmt.Println(count)
}
