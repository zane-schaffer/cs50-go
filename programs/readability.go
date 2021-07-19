package programs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Readability() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	letters, words, sentences := 0.00, 1.00, 0.00

	for _, v := range text {
		if unicode.IsLetter(v) {
			letters++
		} else if unicode.IsSpace(v) {
			words++
		} else if v == '.' || v == '?' || v == '!' {
			sentences++
		}
	}

	fmt.Println("Letter(s):", letters)
	fmt.Println("Word(s):", words)
	fmt.Println("Sentence(s):", sentences)

	L := (letters / words) * 100.00
	S := (sentences / words) * 100.00
	index := int(0.0588*L - 0.296*S - 15.8)

	switch {
	case index < 1:
		fmt.Println("Before Grade 1")
	case index >= 1 && index < 16:
		fmt.Println("Grade", index)
	case index >= 16:
		fmt.Println("Grade 16+")
	}
}
