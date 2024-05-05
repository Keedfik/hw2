package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	var lowText string
	for _, l := range text {
		if unicode.IsSpace(l) {
			l++
		} else {
			lowText += string(unicode.ToLower(l))
		}
	}
	fmt.Printf(lowText + "\n")

	lettersIn := make(map[rune]int)
	for _, l := range lowText {
		lettersIn[l]++
	}

	var percent float64
	for symb, count := range lettersIn {
		percent = 100 / float64(len(lowText)) * float64(count) / 100
		fmt.Printf("%c - %d %0.2f\n", symb, count, percent)
	}

}
