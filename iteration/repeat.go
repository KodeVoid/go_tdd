package iteration

// Repeat takes a string character and an integer count n,
// and returns a new string consisting of the character repeated n times.
func Repeat(character string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
