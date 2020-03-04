package str

func strlen(str string) int {
	return len([]rune(str))
}
