package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}
	root.next = root
	root.prev = root
	return &LList{root: root, size: 0}
}

func (l *LList) String() string {
	var parts []string
	for curr := l.root.next; curr != l.root; curr = curr.next {
		parts = append(parts, fmt.Sprintf("%d", curr.Value))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) PushFront(value int) {
	newNode := &Node{Value: value}
	first := l.root.next
	
	newNode.prev = l.root
	newNode.next = first
	
	l.root.next = newNode
	first.prev = newNode
	l.size++
}

func (l *LList) PushBack(value int) {
	newNode := &Node{Value: value}
	last := l.root.prev
	
	newNode.prev = last
	newNode.next = l.root
	
	last.next = newNode
	l.root.prev = newNode
	l.size++
}

func (l *LList) PopFront() {
	if l.size == 0 {
		return
	}
	first := l.root.next
	second := first.next
	
	l.root.next = second
	second.prev = l.root
	l.size--
}

func (l *LList) PopBack() {
	if l.size == 0 {
		return
	}
	last := l.root.prev
	penultimate := last.prev
	
	l.root.prev = penultimate
	penultimate.next = l.root
	l.size--
}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			// Corrigido: Lendo da esquerda para a direita normalmente
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}