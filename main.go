package main

import (
	"fmt"
)

func main() {
	str := "ИЦРЧЭЫЩШШЩРЬЩЩМДРШУРМЮПРЭЪЩЬЭРЪРШШЩТЛЧРШКЭЗЩМЖВШЩРМЮЧЛСШЩРБЩЩМДРШУР"

	alphabetRune := getAlphabet()
	strRune := []rune(str)
	enCrypt(strRune, alphabetRune, 1)
}

func printRune(r []rune) {
	var str string
	for _, v := range r {
		str += string(v)
	}
	fmt.Println(str)
}

func getAlphabet() []rune {
	var r []rune
	for i := 'А'; i <= 'Я'; i++ {
		r = append(r, i)
	}
	return r
}

func enCrypt(strRune, alphabetRune []rune, shift int) {
	for i := 0; i < 32; i++ {
		var str string
		for _, v := range strRune {
			v = v + rune(i)
			if v > 'Я' {
				v -= 32
			}
			str += string(v)
		}
		fmt.Printf("%v : %s\n", i, str)
	}
}
