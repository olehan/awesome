package pair

type (
	Pair interface {
		First() interface{}
		Second() interface{}
		SetFirst(v interface{}) IPair
		SetSecond(v interface{}) IPair
	}

	pair [2]interface{}
)

func NewPair(first, second interface{}) Pair {
	return &pair{first, second}
}

func (p *pair) First() interface{} {
	return p[0]
}

func (p *pair) Second() interface{} {
	return p[1]
}

func (p *pair) SetFirst(v interface{}) IPair {
	p[0] = v
	return p
}

func (p *pair) SetSecond(v interface{}) IPair {
	p[1] = v
	return p
}
