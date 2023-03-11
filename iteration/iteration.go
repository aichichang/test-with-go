package iteration

func Repeat(char string, num int) string {
	var repeat string

	for i := 0; i < num; i++ {
		repeat += char
	}

	return repeat
}
