package config

import (
	"os"
	"strconv"

	F "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/option"
	P "github.com/IBM/fp-go/predicate"

	PI "profile-svc/predicates/integer"
)

type Port struct {
	Value int
}

func ReadReservedPort(key string) O.Option[Port] {
	// predicate -> string -> Option Port
	// portStr := viper.Get(key).(string)
	portStr := os.Getenv(key)

	return F.Pipe2(
		O.FromNillable(&portStr),
		O.Chain(func(p *string) O.Option[int] {
			portInt, err := strconv.Atoi(*p)

			if err != nil {
				return O.None[int]()
			}
			return O.Some(portInt)
		}),
		O.Chain(func(p int) O.Option[Port] {
			return O.Map(func(i int) Port {
				return Port{Value: i}
			})(O.FromPredicate(P.And(PI.InBetween(0)(1023))(PI.NotZero))(p))
		}),
	)
}

func ReadUnreservedPort(key string) O.Option[Port] {
	// predicate -> string -> Option Port
	// portStr := viper.Get(key).(string)
	portStr := os.Getenv(key)

	return F.Pipe2(
		O.FromNillable(&portStr),
		O.Chain(func(p *string) O.Option[int] {
			portInt, err := strconv.Atoi(*p)

			if err != nil {
				return O.None[int]()
			}
			return O.Some(portInt)
		}),
		O.Chain(func(p int) O.Option[Port] {
			return O.Map(func(i int) Port {
				return Port{Value: i}
			})(O.FromPredicate(P.And(PI.InBetween(1024)(65535))(PI.NotZero))(p))
		}),
	)
}
