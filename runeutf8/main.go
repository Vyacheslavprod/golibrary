package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "ABC"
	str2 := "АБВГДЕ"
	fmt.Println(len(str1))
	fmt.Println(utf8.RuneCountInString(str2))

	str3 := []rune(str2)

	for i, el := range str1 {
		fmt.Println(i, string(el))
	}
	fmt.Println()
	for i, el := range str3 {
		fmt.Println(i, string(el))
	}
}