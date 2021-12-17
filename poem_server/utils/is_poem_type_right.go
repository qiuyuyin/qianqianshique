package utils

func PoemTypeRight(poemType string) bool {
	strings := []string{"shi", "ci", "qu"}
	for _, v := range strings {
		if poemType == v {
			return true
		}
	}
	return false
}
