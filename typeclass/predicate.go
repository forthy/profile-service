package typeclass

type Predicate[T any] = func(v T) bool
