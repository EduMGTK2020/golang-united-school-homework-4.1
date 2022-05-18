package reverse_string

func ReverseString(input string) (output string) {
	var runeString = []rune(input)
	var runeLen = len(runeString)

	for i := 0; i < int(runeLen/2); i++ {
		runeString[runeLen-i-1], runeString[i] = runeString[i], runeString[runeLen-i-1]
	}

	output = string(runeString)
	return output
}
