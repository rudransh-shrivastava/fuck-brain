package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, true
}

func (s *Stack) Peek() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}

func main() {
	if len(os.Args) < 2 {
		println("Usage: fuckbrain <filename>")
		return
	}
	file_name := os.Args[1]
	if strings.Split(file_name, ".")[1] != "bf" {
		println("Invalid file type. Please provide a .bf file")
		return
	}

	file, err := os.ReadFile(file_name)
	if err != nil {
		println("Error opening file:", err)
		return
	}
	arrTape := make([]byte, 30000)
	ptr := 0
	stack := Stack{}
	for i := 0; i < len(file); i++ {
		char := rune(file[i])
		switch char {
		case '>':
			if ptr == len(arrTape)-1 {
				log.Fatal("Index out of bounds of memory tape")
				return
			}
			ptr++
		case '<':
			if ptr == 0 {
				log.Fatal("Index out of bounds of memory tape")
				return
			}
			ptr--
		case '+':
			if arrTape[ptr] == 255 {
				arrTape[ptr] = 0
			}
			arrTape[ptr]++
		case '-':
			if arrTape[ptr] == 0 {
				arrTape[ptr] = 255
			}
			arrTape[ptr]--
		case '.':
			fmt.Printf("%c", arrTape[ptr])
		case ',':
			fmt.Scanf("%c", &arrTape[ptr])
		case '[':
			stack.Push(i)
		case ']':
			if len(stack) == 0 {
				log.Fatal("Unmatched ']' at index ", i)
			}
			start, _ := stack.Pop()
			if arrTape[ptr] != 0 {
				i = start - 1
			}
		}
	}
}
