package typeclass

// show :: t -> string
type Show[T any] = func(T) string
