package iteration

// repeats a character 'n' number of times and returns it
func Repeat(character string, n int) (repeated string) {
	for i := 0; i < n; i++{
		repeated += character
	}

	return
}