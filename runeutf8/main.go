package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "ABC"
	str2 := "АБВ"
	fmt.Println(len(str1))
	fmt.Println(utf8.RuneCountInString(str2))
}