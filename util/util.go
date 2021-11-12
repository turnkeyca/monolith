package util

const REGEX_UUID = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"

type GenericError string
type ValidationError string

func Contains(length int, match func(index int) bool) bool {
	for i := 0; i < length; i++ {
		if match(i) {
			return true
		}
	}
	return false
}
