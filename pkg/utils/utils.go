package utils

func Index(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func Remove(s []string, e string) []string {
	i := Index(s, e)
	if i == -1 {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
