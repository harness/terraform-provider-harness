package utils

func CoalesceStr(val string, fallback string) string {
	if val == "" {
		return fallback
	}
	return val
}

func CoalesceObj(val interface{}, fallback func() interface{}) interface{} {
	if val == nil {
		return fallback()
	}

	return val
}
