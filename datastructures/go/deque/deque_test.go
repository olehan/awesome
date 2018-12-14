package deque

import "testing"

func TestDeque(t *testing.T) {
	deque := NewDeque()
	deque.PushFirst("String").PushFirst(788)

	t.Log(deque.Count())
	t.Log(deque.PopLast())
	t.Log(deque)
}
