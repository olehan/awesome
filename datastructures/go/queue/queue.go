package queue

type (
	Queue interface {
		Enqueue(v interface{}) Queue
		Dequeue() interface{}
		Peek() interface{}
		Count() int
	}

	queue []interface{}
)

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Enqueue(v interface{}) Queue {
	*q = append(*q, v)
	return q
}

func (q *queue) Dequeue() interface{} {
	last := (*q)[0]
	*q = (*q)[1:]
	return last
}

func (q *queue) Peek() interface{} {
	return (*q)[0]
}

func (q *queue) Count() int {
	return len(*q)
}
