//Объединение строк
package prose

import "strings"

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result
	}
}

/*
fmt.Println(strings.Join([]string{"05", "14", "2018"}, "/"))
fmt.Println(strings.Join([]string{"state", "of", "the", "art"}, "-"))

fmt.Println(JoinWithCommas([]string{"apple", "orange", "pear", "banana"}))
*/