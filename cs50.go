package main

import (
	"fmt"

	"github.com/zane-schaffer/cs50/problems"
	"github.com/zane-schaffer/cs50/utils"
)

func main() {
	problemList := []string{
		"mario",
		"cash",
		"scrabble",
		"readability",
	}
	runProblem(problemList)
}

func runProblem(problemList []string) {
	var input string
	fmt.Println("Which problem would you like to run?")
	fmt.Println("Enter `ls` for all possible problems")
loop:
	for {
		input = utils.StringInput("")
		switch input {
		case "ls":
			fmt.Println("----------------")
			fmt.Print("\n")
			for _, v := range problemList {
				fmt.Println(v)
			}
			fmt.Print("\n")
			fmt.Println("----------------")
			continue loop
		case "mario":
			problems.Mario()
			break loop
		case "cash":
			problems.Cash()
			break loop
		case "scrabble":
			problems.Scrabble()
			break loop
		case "readability":
			problems.Readability()
			break loop
		default:
			fmt.Printf("Problem %s doesn't exist\n", input)
			continue loop
		}
	}
}
