package utils

func MapToStringSlice(m map[string]string, separator string) []string {
	var keys []string
	for k, v := range m {
		keys = append(keys, k+separator+v)
	}
	return keys
}
