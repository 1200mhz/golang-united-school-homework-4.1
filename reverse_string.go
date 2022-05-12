package reverse_string

func ReverseString(input string) (output string) {
	r := []rune(input)
	c := make([]rune, len(r))

	for k, v := range r {
		c[len(r)-k-1] = v
	}

	return string(c)
}
