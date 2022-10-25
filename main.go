package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"Cammyhaven", "East Willisstad", "Lake Thômas"}
	stringConcatV1 := subOptimalStringConcat(s)
	stringConcatV2 := stringConcat(s)
	stringConcatV3 := optimalStringConcat(s)
	fmt.Println(stringConcatV1)
	fmt.Println(stringConcatV2)
	fmt.Println(stringConcatV3)
}

func subOptimalStringConcat(str []string) string {
	var s string
	for _, v := range str {
		s += v
	}
	return s
}

func stringConcat(values []string) string {
	sb := strings.Builder{}
	for _, v := range values {
		sb.WriteString(v)
	}
	return sb.String()
}
func optimalStringConcat(values []string) string {
	total := 0
	// here we are not interested in the number of runes
	// but the number of bytes hence we use the len function
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}
	sb := strings.Builder{}
	// preallocate a total number bytes before iteration
	sb.Grow(total)
	for _, v := range values {
		sb.WriteString(v)
	}
	return sb.String()
}
