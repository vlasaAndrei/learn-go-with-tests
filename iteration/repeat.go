package iteration

import "strings"

func Repeat(character string, numberOfRepeats int) string {
	var repeated strings.Builder

	for range numberOfRepeats {
		repeated.WriteString(character)
	}

	return repeated.String()
}
