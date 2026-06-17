package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	}
	return 0
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	input := scanner.Text()
	tokens := strings.Fields(input)

	var output []string
	var stack []string

	for _, token := range tokens {
		if precedence(token) == 0 {
			output = append(output, token)
		} else {
			for len(stack) > 0 && precedence(stack[len(stack)-1]) >= precedence(token) {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)

		}

	}
	for len(stack) > 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	fmt.Println(strings.Join(output, " "))
}
