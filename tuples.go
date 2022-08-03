package collections

func First[T1, T2 any](a T1, b T2) T1 {
	return a
}

func Second[T1, T2 any](a T1, b T2) T2 {
	return b
}
