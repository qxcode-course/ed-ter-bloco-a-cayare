package main

import (
	"fmt"
	"strings"
)


func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	target := it.next
	if target == l.root {
		target = target.next
	}
	return target
}


func ToStr(l *DList[int], sword *DNode[int]) string {
	var elements []string
	for n := l.root.next; n != l.root; n = n.next {
		if n == sword {
			elements = append(elements, fmt.Sprintf("%v>", n.Value))
		} else {
			elements = append(elements, fmt.Sprintf("%v", n.Value))
		}
	}
	return "[ " + strings.Join(elements, " ") + " ]"
}


func main() {
	var qtd, chosen int
	if _, err := fmt.Scan(&qtd, &chosen); err != nil {
		return
	}

	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}

	
	sword := l.Front()
	for i := 0; i < chosen-1; i++ {
		sword = Next(l, sword)
	}

	
	for l.Size() > 1 {
		fmt.Println(ToStr(l, sword))
		
		
		victim := Next(l, sword)
		l.Erase(victim)
		
		
		sword = Next(l, sword)
	}
	
	
	fmt.Println(ToStr(l, sword))
}