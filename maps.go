package collections

func ConvertMap[K1 comparable, V1 any, K2 comparable, V2 any](
	m map[K1]V1, conv func(k K1, v V1) (K2, V2, error)) (map[K2]V2, error) {

	converted := make(map[K2]V2, len(m))

	for k, v := range m {
		convertedK, convertedV, err := conv(k, v)

		if err != nil {
			return nil, err
		}

		converted[convertedK] = convertedV
	}

	return converted, nil
}

// MapKeys returns a slice []KeyType with the keys of the map[KeyType]ValueType object.
func MapKeys[KeyType comparable, ValueType any](m map[KeyType]ValueType) []KeyType {
	l := make([]KeyType, len(m))
	i := 0

	for k := range m {
		l[i] = k
		i++
	}

	return l
}

// MapValues returns a slice []ValueType with the values of the map[KeyType]ValueType object.
func MapValues[KeyType comparable, ValueType any](m map[KeyType]ValueType) []ValueType {
	l := make([]ValueType, len(m))
	i := 0

	for _, v := range m {
		l[i] = v
		i++
	}

	return l
}

func MapHas[KeyType comparable, ValueType any](m map[KeyType]ValueType, k KeyType) bool {
	_, has := m[k]

	return has
}
