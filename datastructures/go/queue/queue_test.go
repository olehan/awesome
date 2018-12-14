package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue("Adam")
	queue.Enqueue("John")
	queue.Enqueue("Anthony")

	t.Log(queue)
	t.Log(queue.Dequeue())
	t.Log(queue.Peek())
}
