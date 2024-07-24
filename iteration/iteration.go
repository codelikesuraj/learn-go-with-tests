package iteration

import "strings"

func Repeat(character string, count int) string {
	// var repeated string

	// for range count {
	// 	repeated += character
	// }
	
	// return repeated
	return strings.Repeat(character, count)
}
