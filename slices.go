package util

func Filter[T any](s []T, f func(T) bool) []T {
	var filtered []T

	for _, v := range s {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func Convert[T1, T2 any](s []T1, conv func(T1) (T2, error)) ([]T2, error) {
	converted := make([]T2, len(s))

	for i, v := range s {
		nv, err := conv(v)
		if err != nil {
			return nil, err
		}

		converted[i] = nv
	}

	return converted, nil
}

func Compare[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}
