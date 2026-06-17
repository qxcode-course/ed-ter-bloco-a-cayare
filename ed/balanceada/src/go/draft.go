package main

import (
	"bufio"
	"fmt"
	"os"
)

func isBalanced(s string) bool {
	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(', '[':
			stack = append(stack, char)

		case ')', ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (char == ')' && top != '(') || (char == ']' && top != '[') {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		if isBalanced(input) {
			fmt.Println("balanceado")
		} else {
			fmt.Println("nao balanceado")
		}
	}

}
