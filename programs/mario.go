package programs

import (
	"fmt"

	"github.com/zane-schaffer/cs50/utils"
)

func Mario() {
	var n int
	fmt.Println("Length?")
	for {
		n = utils.IntInput()

		if n < 8 {
			break
		} else {
			fmt.Println("Length must be less than 8")
		}
	}

	fmt.Print("\n")

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for h := 1; h <= i; h++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}

}
