package model

import (
	"testing"

	"github.com/stretchr/testify/assert"

	D "profile-svc/predicates/domain"

	O "github.com/IBM/fp-go/option"
)

func TestDomainWithPredicateOf(t *testing.T) {
	d1Opt := DomainWithPredicateOf(D.DomainCompliant)("SD")

	assert.True(t, O.IsNone(d1Opt))

	O.Fold(
		func() bool {
			return assert.Fail(t, "EA should be a correct domain")
		},
		func(d Domain) bool {
			return assert.Equal(t, "EA", d.Title())
		},
	)(DomainWithPredicateOf(D.DomainCompliant)("EA"))
}

func TestIdOf(t *testing.T) {
	idO := IdOf()
	
	assert.True(t, O.IsSome(idO))
}
