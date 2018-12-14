package pair

import "testing"

func TestPair(t *testing.T) {
	pair := NewPair(78, 15)

	t.Log(pair.First())
	t.Log(pair.SetFirst("String").First())
	t.Log(pair)
}
