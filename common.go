package golib

func Contains[T comparable](a []T, b T) bool {
	for _, element := range a {
		if element == b {
			return true
		}
	}
	return false
}
