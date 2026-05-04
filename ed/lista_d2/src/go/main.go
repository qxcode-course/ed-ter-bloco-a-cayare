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

func (n *Node) GetNext() *Node { return n.Next }
func (n *Node) GetPrev() *Node { return n.Prev }

type LList struct {
	head *Node
	tail *Node
	size int
}

func NewLList() *LList {
	return &LList{}
}

func (ll *LList) Front() *Node { return ll.head }
func (ll *LList) Back() *Node  { return ll.tail }
func (ll *LList) Clear(){
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}
func(ll *LList) Search(value int) *Node{
	for node := ll.head; node != nil; node = node.Next{
		if node.Value == value{
			return node
		}

		
	}
	return nil

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
func (ll *LList) Insert(pivot *Node, value int){
	if pivot == ll.head{
		ll.PushFront(value)
		return
	}
	newNode := &Node{Value: value}
	before := pivot.Prev

	newNode.Next = pivot
	newNode.Prev = before

	before.Next = newNode
	pivot.Prev = newNode

	ll.size++
}
func (ll *LList) Remove(node *Node) {
	if node == ll.head && node == ll.tail {
		ll.head = nil
		ll.tail = nil
	} else if node == ll.head {
		ll.head = node.Next
		ll.head.Prev = nil
	} else if node == ll.tail {
		ll.tail = node.Prev
		ll.tail.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	ll.size--
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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.GetNext() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.GetPrev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			 ll.Clear()
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			pivotValue, _ := strconv.Atoi(args[1])
			newValue, _ := strconv.Atoi(args[2])
			node := ll.Search(pivotValue)
			if node != nil {
				ll.Insert(node, newValue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			value, _ := strconv.Atoi(args[1])
			node := ll.Search(value)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}