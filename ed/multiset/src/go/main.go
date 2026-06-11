package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	if capacity <= 0 {
		capacity = 2
	}
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	if ms.capacity == 0 {
		ms.capacity = 2
	} else {
		ms.capacity *= 2
	}
	newData := make([]int, ms.capacity)
	copy(newData, ms.data[:ms.size])
	ms.data = newData
}

func (ms *MultiSet) search(value int) (bool, int) {
	left := 0
	right := ms.size - 1
	for left <= right {
		mid := left + (right-left)/2
		if ms.data[mid] == value {
			return true, mid
		} else if ms.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false, left
}
func (ms *MultiSet) Insert(value int) {
	if ms.size >= ms.capacity {
		ms.expand()
	}
	_, index := ms.search(value)
	ms.insert(value, index)
}

func (ms *MultiSet) insert(value int, index int) error {
	if index < 0 || index > ms.size {
		return errors.New("index out of bounds")
	}

	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[index] = value
	ms.size++
	return nil
}

func (ms *MultiSet) Erase(value int) error {
	found, index := ms.search(value)
	if !found {
		return errors.New("value not found")
	}
	return ms.erase(index)
}

func (ms *MultiSet) erase(index int) error {
	if index < 0 || index >= ms.size {
		return errors.New("index out of bounds")
	}

	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	return nil
}
func (ms *MultiSet) Contains(value int) bool {
	found, _ := ms.search(value)
	return found
}
func (ms *MultiSet) Count(value int) int {
	found, index := ms.search(value)
	if !found {
		return 0
	}

	count := 0

	for i := index; i >= 0 && ms.data[i] == value; i-- {
		count++
	}

	for i := index + 1; i < ms.size && ms.data[i] == value; i++ {
		count++
	}
	return count
}

func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}
	uniqueCount := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			uniqueCount++
		}
	}
	return uniqueCount
}

func (ms *MultiSet) Clear() {
	ms.size = 0
}

func (ms *MultiSet) String() string {
	if ms.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < ms.size; i++ {
		sb.WriteString(strconv.Itoa(ms.data[i]))
		if i < ms.size-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
