package model

import (
	PM "profile-svc/predicates/domain"
	TC "profile-svc/typeclass"

	F "github.com/IBM/fp-go/function"
	I "github.com/IBM/fp-go/iooption"
	O "github.com/IBM/fp-go/option"

	T "github.com/rushysloth/go-tsid"
)

// Domain modelling
/*
- Should a Profile type generic or a sum type?
data Id = UUID
data Name = string
data Domain = EA | MFG | ENG
data Profile = Id Name Domain
*/

type Id struct {
	Value string
}

type Name struct {
	Value string
}

type Version struct {
	Value string
}

type Domain interface {
	Title() string
}

type EA struct{}
type MFG struct{}
type ENG struct{}

func (d EA) Title() string {
	return "EA"
}

func (d MFG) Title() string {
	return "MFG"
}

func (d ENG) Title() string {
	return "ENG"
}

type Profile struct {
	Id      Id
	Name    Name
	Version Version
	Domain  Domain
}

func ProfileOf(id Id) func(Name) func(Version) func(Domain) Profile {
	return func(name Name) func(version Version) func(domain Domain) Profile {
		return func(version Version) func(domain Domain) Profile {
			return func(domain Domain) Profile {
				return Profile{id, name, version, domain}
			}
		}
	}
}

func NameWithPredicateOf(p TC.Predicate[string]) func(string) O.Option[Name] {
	return func(s string) O.Option[Name] {
		return F.Ternary(
			p,
			func(v string) O.Option[Name] {
				return O.Some(Name{Value: v})
			},
			func(_ string) O.Option[Name] {
				return O.None[Name]()
			},
		)(s)
	}
}

func VersionWithPredicateOf(p TC.Predicate[string]) func(string) O.Option[Version] {
	return func(s string) O.Option[Version] {
		return F.Ternary(
			p,
			func(v string) O.Option[Version] {
				return O.Some(Version{Value: v})
			},
			func(_ string) O.Option[Version] {
				return O.None[Version]()
			},
		)(s)
	}
}

func provideTSIDStr() O.Option[string] {
	return O.Chain(
		func(f *T.TsidFactory) O.Option[string] {
			return O.Map(
				func(id *T.Tsid) string {
					return id.ToString()
				},
			)(O.TryCatch(f.Generate))
		},
	)(O.TryCatch(T.TsidFactoryBuilder().Build))
}

func IdWithTSIDOf(tsidProvider I.IOOption[string]) O.Option[Id] {
	return O.Map(func(s string) Id {
		return Id{Value: s}
	})(tsidProvider())
}

func IdOf() O.Option[Id] {
	return IdWithTSIDOf(provideTSIDStr)
}

func IdWithPredicateOf(p TC.Predicate[string]) func(string) O.Option[Id] {
	return func(s string) O.Option[Id] {
		return F.Ternary(
			p,
			func(v string) O.Option[Id] {
				return O.Some(Id{Value: v})
			},
			func(_ string) O.Option[Id] {
				return O.None[Id]()
			},
		)(s)
	}
}

func DomainWithPredicateOf(p TC.Predicate[string]) func(string) O.Option[Domain] {
	return func(s string) O.Option[Domain] {
		return F.Ternary(
			p,
			func(v string) O.Option[Domain] {
				switch v {
				case "EA":
					return O.Some(Domain(EA{}))
				case "MFG":
					return O.Some(Domain(MFG{}))
				case "ENG":
					return O.Some(Domain(ENG{}))
				default:
					return O.None[Domain]()
				}
			},
			func(_ string) O.Option[Domain] {
				return O.None[Domain]()
			},
		)(s)
	}
}

func DomainOf(s string) O.Option[Domain] {
	return DomainWithPredicateOf(PM.DomainCompliant)(s)
}
