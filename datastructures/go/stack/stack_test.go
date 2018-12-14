package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.Push("read message")
	stack.Push("watch some videos")
	stack.Push("feed some cats")
	stack.Push("and pop some stacks")

	t.Log(stack)
	t.Log(stack.Pop())
	t.Log(stack.Peek())
}
