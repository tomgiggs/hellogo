package utils

func StringInSlice(arr []string, t string) bool {
	for _, s := range arr {
		if t == s {
			return true
		}
	}
	return false
}
