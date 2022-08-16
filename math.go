package util

type SignedNumerals interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Abs[T SignedNumerals](v T) T {
	if v < 0 {
		return -v
	}

	return v
}
