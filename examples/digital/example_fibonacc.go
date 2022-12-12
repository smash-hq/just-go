package digital

import "container/list"

var (
	RecursionTimes            = 0
	StackTimes                = 0
	stack          *list.List = list.New()
	RecursionCh               = make(chan bool, 1)
	StackCh                   = make(chan bool, 1)
)

func FibonaccWithRecursion(n int) int {
	if n == 1 || n == 2 {
		return n - 1
	}
	RecursionTimes++
	return FibonaccWithRecursion(n-1) + FibonaccWithRecursion(n-2)
}

func FibonaccWithStack(n int) int {
	if n == 1 || n == 2 {
		return n - 1
	}
	stack.PushBack(0)
	stack.PushBack(1)
	for i := 2; i < n; i++ {
		a := stack.Back()
		stack.Remove(a)
		b := stack.Back()
		stack.Remove(b)
		stack.PushBack(a.Value.(int))
		stack.PushBack(a.Value.(int) + b.Value.(int))
		StackTimes += 5
	}
	a := stack.Back()
	res := stack.Remove(a)
	return res.(int)
}
