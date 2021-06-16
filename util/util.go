package util

func Contains(length int, match func(index int) bool) bool {
	for i := 0; i < length; i++ {
		if match(i) {
			return true
		}
	}
	return false
}
