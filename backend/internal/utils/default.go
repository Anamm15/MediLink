package utils

func GetBoolOrDefault(value *bool, defaultValue bool) bool {
	if value != nil {
		return *value
	}

	return defaultValue
}
