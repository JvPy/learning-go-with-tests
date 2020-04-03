package interation

//Repeat - Build a string of x characters
func Repeat(char string, times int) (newString string) {
	for i := 0; i < times; i++ {
		newString += char
	}

	return
}
