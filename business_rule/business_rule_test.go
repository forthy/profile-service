package business_rule

import (
	"github.com/stretchr/testify/assert"
	"testing"

	E "github.com/IBM/fp-go/either"
	W "github.com/IBM/fp-go/writer"
)

const (
	configJson = `{"X":"hello","Y":"world"}`
)

func TestCreateRule(t *testing.T) {
	config := map[string]string{
		"X": "hello",
		"Y": "world",
	}

	wcr := CreateRule(config)
	descriptions := W.Execute(wcr)
	result := W.Evaluate(wcr)

	assert.Equal(t, descriptions, []string{configJson})
	assert.Equal(t, E.Right[error](Rule{config}), result)
}
