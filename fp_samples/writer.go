package fp_samples

import (
	"fmt"
	A "github.com/IBM/fp-go/array"
	P "github.com/IBM/fp-go/pair"
	W "github.com/IBM/fp-go/writer"
)

// Add int -> Writer []string int
func Add(v int) func(int) W.Writer[[]string, int] {
	return func(s int) W.Writer[[]string, int] {
		r := v + s
		return func() P.Pair[int, []string] {
			return P.MakePair(r, A.Of(fmt.Sprintf("%d + %d = %d", v, s, r)))
		}
	}
}

// Minus int -> Writer []string int
func Minus(v int) func(int) W.Writer[[]string, int] {
	return func(s int) W.Writer[[]string, int] {
		r := v - s
		return func() P.Pair[int, []string] {
			return P.MakePair(r, A.Of(fmt.Sprintf("%d - %d = %d", v, s, r)))
		}
	}
}
