package common

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
