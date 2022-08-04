package collections

type Generator[T any] interface {
	Next() bool
	Value() T
}
