package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	count := 0
	done := (1 << n) - 1

	var solve func(cols, diag1, diag2 int)
	solve = func(cols, diag1, diag2 int) {
		if cols == done {
			count++
			return
		}

		available := done & ^(cols | diag1 | diag2)

		for available > 0 {

			p := available & -available
			available -= p

			solve(cols|p, (diag1|p)<<1, (diag2|p)>>1)
		}
	}

	solve(0, 0, 0)
	fmt.Println(count)
}
