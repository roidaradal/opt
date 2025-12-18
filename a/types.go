// Package a contains types used in the discrete optimization problems
package a

type StrInt struct {
	Str string
	Int int
}

// Create new StrInt
func NewStrInt(str string, x int) StrInt {
	return StrInt{str, x}
}

// Destructure StrInt
func (s StrInt) Tuple() (string, int) {
	return s.Str, s.Int
}
