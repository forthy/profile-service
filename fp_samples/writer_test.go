package fp_samples

import (
	A "github.com/IBM/fp-go/array"
	F "github.com/IBM/fp-go/function"
	W "github.com/IBM/fp-go/writer"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	r := Add(1)(2)

	assert.Equal(t, 3, W.Evaluate(r))
	assert.Equal(t, []string{"1 + 2 = 3"}, W.Execute(r))
}

func TestMinus(t *testing.T) {
	r := Minus(3)(1)

	assert.Equal(t, 2, W.Evaluate(r))
	assert.Equal(t, []string{"3 - 1 = 2"}, W.Execute(r))
}

func TestMix(t *testing.T) {
	r := F.Pipe1(
		Add(4)(13),
		W.Chain(A.Semigroup[string](), func(v int) W.Writer[[]string, int] {
			return Minus(v)(6)
		}),
	)

	assert.Equal(t, 11, W.Evaluate(r))
	assert.Equal(t, []string{"4 + 13 = 17", "17 - 6 = 11"}, W.Execute(r))
}
