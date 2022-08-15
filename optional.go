package collections

type Optional[T any] struct {
	value    T
	hasValue bool
}

func (o *Optional[T]) HasValue() bool {
	return o.hasValue
}

func (o *Optional[T]) Value() T {
	return o.value
}

func (o *Optional[T]) Unwrap() (T, bool) {
	return o.value, o.hasValue
}

func MakeOptional[T any](v T) Optional[T] {
	return Optional[T]{
		value:    v,
		hasValue: true,
	}
}
