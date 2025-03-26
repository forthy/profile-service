package service

import (
	PR "profile-svc/model"
	PS "profile-svc/predicates/string"

	"testing"

	E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/option"
	W "github.com/IBM/fp-go/writer"

	JA "github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
)

const (
	sampleDescription = `{"id":"XU-1352","name":"XM-Data-Loader","version":"1.24.1","domain":"EA","registerTime":"<<PRESENCE>>"}`
)

func TestRegisterWithShow(t *testing.T) {
	profileO := F.Pipe4(
		O.Of(PR.ProfileOf),
		O.Ap[func(PR.Name) func(version PR.Version) func(domain PR.Domain) PR.Profile](O.Of(PR.Id{Value: "XU-1352"})),
		O.Ap[func(version PR.Version) func(domain PR.Domain) PR.Profile](PR.NameWithPredicateOf(PS.InBetween(5)(15))("XM-Data-Loader")),
		O.Ap[func(domain PR.Domain) PR.Profile](PR.VersionWithPredicateOf(PS.InBetween(3)(6))("1.24.1")),
		O.Ap[PR.Profile](PR.DomainOf("EA")),
	)

	wo := O.Map(func(p PR.Profile) W.Writer[[]string, E.Either[error, PR.Profile]] {
		return RegisterWithShow(RegisterEventShow)(p)
	})(profileO)

	O.Fold(
		func() bool {
			return assert.Fail(t, "Failed to construct a profile")
		},
		func(we W.Writer[[]string, E.Either[error, PR.Profile]]) bool {
			pe := W.Evaluate(we)
			de := W.Execute(we)

			ja := JA.New(t)
			ja.Assertf(de[0], sampleDescription)

			return E.Fold(
				func(e error) bool {
					return assert.Fail(t, e.Error())
				},
				func(p PR.Profile) bool {
					assert.Equal(t, "XU-1352", p.Id.Value)
					assert.Equal(t, "1.24.1", p.Version.Value)
					assert.Equal(t, "EA", p.Domain.Title())
					return assert.Equal(t, "XM-Data-Loader", p.Name.Value)
				},
			)(pe)
		},
	)(wo)
}
