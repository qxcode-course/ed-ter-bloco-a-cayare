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
	Next  *Node
	Prev  *Node
}

type LList struct {
	head *Node
	tail *Node
	size int
}

func NewLList() *LList {
	return &LList{}
}

func (ll *LList) PushBack(value int) {
	newNode := &Node{Value: value}
	if ll.size == 0 {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.Prev = ll.tail
		ll.tail.Next = newNode
		ll.tail = newNode
	}
	ll.size++
}

func (ll *LList) PushFront(value int) {
	newNode := &Node{Value: value}
	if ll.size == 0 {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.Next = ll.head
		ll.head.Prev = newNode
		ll.head = newNode
	}
	ll.size++
}

func (ll *LList) String() string {
	var elements []string
	for node := ll.head; node != nil; node = node.Next {
		elements = append(elements, strconv.Itoa(node.Value))
	}
	return "[" + strings.Join(elements, ", ") + "]"
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
			// fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			// ll.Clear()
		case "walk":
			// Requer Next() e Prev()
		case "replace":
			// Requer Search()
		case "insert":
			// Requer Search() e Insert()
		case "remove":
			// Requer Search() e Remove()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}