package EvaluateReversePolishNotation

import (
	"log"
	"strconv"
)

func getOperator(token string) func(a, b int) int {
	switch token {
	case "+":
		return func(a, b int) int {
			return a + b
		}
	case "-":
		return func(a, b int) int {
			return a - b
		}
	case "*":
		return func(a, b int) int {
			return a * b
		}
	case "/":
		return func(a, b int) int {
			return a / b
		}
	}
	return func(a, b int) int {
		log.Fatal("wrong token", token)
		return 0
	}
}

type Stack struct {
	nums []int
}

func (s *Stack) Insert(n int) {
	s.nums = append(s.nums, n)
}

func (s *Stack) Pop() int {
	n := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return n
}

func evalRPN(tokens []string) int {
	stack := Stack{[]int{}}
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			operator := getOperator(token)
			n1 := stack.Pop()
			n2 := stack.Pop()
			num = operator(n2, n1)
		}
		stack.Insert(num)
	}
	return stack.Pop()
}
