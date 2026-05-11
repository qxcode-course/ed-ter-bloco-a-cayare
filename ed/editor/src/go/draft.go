package main

import (
	"bufio"
	"fmt"
	"os"
)

type DNode struct {
	Value rune
	next  *DNode
	prev  *DNode
	root  *DNode
}

type DList struct {
	root *DNode
	size int
}

func NewDList() *DList {
	root := &DNode{}
	root.next = root
	root.prev = root
	root.root = root
	return &DList{root: root, size: 0}
}

func (l *DList) Insert(it *DNode, value rune) *DNode {
	n := &DNode{Value: value, root: l.root}
	n.prev = it.prev
	n.next = it
	it.prev.next = n
	it.prev = n
	l.size++
	return n
}

func (l *DList) Erase(it *DNode) *DNode {
	if it == l.root || it == nil {
		return it
	}
	next := it.next
	it.prev.next = it.next
	it.next.prev = it.prev
	l.size--
	return next
}

func (l *DList) ToString(cursor *DNode) string {
	var res []rune
	for n := l.root.next; n != l.root; n = n.next {
		if n == cursor {
			res = append(res, '|')
		}
		res = append(res, n.Value)
	}
	if cursor == l.root {
		res = append(res, '|')
	}
	return string(res)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	input := scanner.Text()

	l := NewDList()
	cursor := l.root

	for _, char := range input {
		switch char {
		case 'R':
			l.Insert(cursor, '\n')
		case 'B':
			if cursor.prev != l.root {
				l.Erase(cursor.prev)
			}
		case 'D':
			if cursor != l.root {
				cursor = l.Erase(cursor)
			}
		case '>':
			if cursor != l.root {
				cursor = cursor.next
			}
		case '<':
			if cursor.prev != l.root {
				cursor = cursor.prev
			}
		default:
			l.Insert(cursor, char)
		}
	}

	fmt.Println(l.ToString(cursor))
}