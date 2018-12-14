package deque

type (
	Deque interface {
		// Adds a new item to the front of the list
		PushFirst(v interface{}) Deque

		// Adds a new item to the back of the list
		PushLast(v interface{}) Deque

		// Deletes the first item in the list and returns it
		PopFirst() interface{}

		// Returns the first list item
		PopLast() interface{}

		// Returns the first list item
		PeekFirst() interface{}

		// Returns the last list item
		PeekLast() interface{}

		// Returns the list size
		Count() int
	}

	deque []interface{}
)

func NewDeque() Deque {
	return &deque{}
}

func (d *deque) PushFirst(v interface{}) Deque {
	*d = append([]interface{}{v}, *d...)
	return d
}

func (d *deque) PushLast(v interface{}) Deque {
	*d = append(*d, v)
	return d
}

func (d *deque) PopFirst() interface{} {
	first := (*d)[0]
	*d = (*d)[1:]
	return first
}

func (d *deque) PopLast() interface{} {
	lastIndex := d.Count() - 1
	last := (*d)[lastIndex]
	*d = (*d)[:lastIndex]
	return last
}

func (d *deque) PeekFirst() interface{} {
	return (*d)[0]
}

func (d *deque) PeekLast() interface{} {
	return (*d)[d.Count() - 1]
}

func (d *deque) Count() int {
	return len(*d)
}
