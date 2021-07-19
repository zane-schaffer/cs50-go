package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func StringInput(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func FloatInput() float64 {
	var f float64
	fmt.Scanln(&f)
	return f
}

func IntInput() int {
	var i int
	fmt.Scanln(&i)
	return i
}
