package ord

import (
	E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/option"
	OR "github.com/IBM/fp-go/ord"
	P "github.com/IBM/fp-go/predicate"

	PS "profile-svc/predicates/string"

	SM "github.com/Masterminds/semver/v3"
)

type Version struct {
	Num string
}

func versionOf(n string) O.Option[Version] {
	return F.Ternary(
		P.And(PS.ShouldBeSemVer)(PS.NotEmpty),
		func(v string) O.Option[Version] {
			return O.Some(Version{Num: v})
		},
		func(v string) O.Option[Version] {
			return O.None[Version]()
		},
	)(n)
}

var VersionOrd = OR.MakeOrd(
	func(x, y Version) int {
		xvE := E.Eitherize1(SM.StrictNewVersion)(x.Num)
		yvE := E.Eitherize1(SM.StrictNewVersion)(y.Num)

		return E.GetOrElse(func(_ error) int {
			return -1
		})(E.Chain(func(v *SM.Version) E.Either[error, int] {
			return E.Map[error](func(v2 *SM.Version) int {
				return v.Compare(v2)
			})(yvE)
		})(xvE))
	},
	func(x, y Version) bool {
		return x.Num == y.Num
	},
)
