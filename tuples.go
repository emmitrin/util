package util

func First[T1, T2 any](a T1, b T2) T1 {
	return a
}

func Second[T1, T2 any](a T1, b T2) T2 {
	return b
}

type Tuple[T1, T2 any] struct {
	first  T1
	second T2
}

func (t *Tuple[T1, T2]) First() T1 {
	return t.first
}

func (t *Tuple[T1, T2]) Second() T2 {
	return t.second
}

func NewTuple[T1, T2 any](first T1, second T2) *Tuple[T1, T2] {
	return &Tuple[T1, T2]{first, second}
}
