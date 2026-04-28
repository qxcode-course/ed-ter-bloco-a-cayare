package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s *Set) reserve(newCapacity int) {
	if newCapacity <= s.capacity {
		return
	}
	newData := make([]int, newCapacity)
	for i := 0; i < s.size; i++ {
		newData[i] = s.data[i]
	}
	s.data = newData
	s.capacity = newCapacity
}

func (s *Set) binarySearch(value int) int {
	low := 0
	high := s.size - 1
	for low <= high {
		mid := (low + high) / 2
		if s.data[mid] == value {
			return mid
		}
		if s.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func (s *Set) Insert(value int) {
	if s.binarySearch(value) != -1 {
		return
	}

	if s.size == s.capacity {
		if s.capacity == 0 {
			s.reserve(1)
		} else {
			s.reserve(s.capacity * 2)
		}
	}

	index := 0
	for index < s.size && s.data[index] < value {
		index++
	}

	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[index] = value
	s.size++
}

func (s *Set) Contains(value int) bool {
	return s.binarySearch(value) != -1
}

func (s *Set) Erase(value int) bool {
	index := s.binarySearch(value)
	if index == -1 {
		return false
	}

	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
	return true
}

func (s *Set) String() string {
	if s.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < s.size; i++ {
		sb.WriteString(strconv.Itoa(s.data[i]))
		if i < s.size-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s *Set = NewSet(0)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("$" + line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd := parts[0]

		switch cmd {
		case "init":
			cap, _ := strconv.Atoi(parts[1])
			s = NewSet(cap)
		case "insert":
			for _, part := range parts[1:] {
				val, _ := strconv.Atoi(part)
				s.Insert(val)
			}
		case "show":
			fmt.Println(s.String())
		case "contains":
			val, _ := strconv.Atoi(parts[1])
			if s.Contains(val) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "erase":
			val, _ := strconv.Atoi(parts[1])
			if !s.Erase(val) {
				fmt.Println("value not found")
			}
		case "clear":
			s.size = 0
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}