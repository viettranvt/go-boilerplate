package string_util

import "strings"

func NormalizedData(data string, isToUpper bool) string {
	data = strings.TrimSpace(strings.Join(strings.Fields(data), " "))

	if isToUpper {
		return strings.ToUpper(data)
	}

	return data
}
