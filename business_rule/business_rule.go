package business_rule

import (
	E "github.com/IBM/fp-go/either"
	J "github.com/IBM/fp-go/json"
	P "github.com/IBM/fp-go/pair"
	W "github.com/IBM/fp-go/writer"
)

type Rule struct {
	Configs map[string]string
}

// data Rule = Id Version map[string]string
// map[string]string -> Writer []string Either error Rule
func CreateRule(config map[string]string) W.Writer[[]string, E.Either[error, Rule]] {
	return func() P.Pair[E.Either[error, Rule], []string] {
		description := E.GetOrElse(
			func(e error) string {
				return e.Error()
			},
		)(E.Map[error, []byte, string](
			func(b []byte) string {
				return string(b)
			},
		)(J.Marshal(config)))

		return P.MakePair(E.Right[error](Rule{config}), []string{description})
	}
}
