package main

import (
	"fmt"

	"github.com/zane-schaffer/cs50/programs"
	"github.com/zane-schaffer/cs50/utils"
)

func main() {
	programList := []string{
		"mario",
		"cash",
		"scrabble",
		"readability",
	}
	runProgram(programList)
}

func runProgram(programList []string) {
	var input string
	fmt.Println("Which program would you like to run?")
	fmt.Println("Enter `ls` for all possible programs")
loop:
	for {
		input = utils.StringInput("")
		switch input {
		case "ls":
			fmt.Println("----------------")
			fmt.Print("\n")
			for _, v := range programList {
				fmt.Println(v)
			}
			fmt.Print("\n")
			fmt.Println("----------------")
			continue loop
		case "mario":
			programs.Mario()
			break loop
		case "cash":
			programs.Cash()
			break loop
		case "scrabble":
			programs.Scrabble()
			break loop
		case "readability":
			programs.Readability()
			break loop
		default:
			fmt.Printf("Program %s doesn't exist\n", input)
			continue loop
		}
	}
}
