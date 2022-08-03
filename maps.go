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
