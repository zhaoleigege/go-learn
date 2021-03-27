package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isAllEnglishWord(str string) bool {
	if str == "" {
		return true
	}
	if strings.TrimSpace(str) == "" {
		return true
	}
	str = strings.Replace(str, "\u200B", " ", -1)

	reg, err := regexp.Compile("^[A-Za-z\\s]+$")
	if err != nil {
		fmt.Println("regex编译错")
		return false
	}
	result := false
	test := reg.FindAllString(str, -1)
	for i := 0; i < len(test); i++ {
		result = true
	}

	fmt.Println(reg.FindStringSubmatch(str))

	return result
}

func main() {
	str1 := "UBONRAT ​THONGHOR "
	str2 := "UBONRAT THONGHOR "

	fmt.Println(isAllEnglishWord(str1))
	fmt.Println(isAllEnglishWord(str2))

	// fmt.Println(strings.Compare(str1, str2))

	// fmt.Println([]byte(str1))
	// fmt.Println([]byte(str2))

	// for i := 0; i < len(str1); i++ {
	// 	fmt.Printf("%c = %c : %t\n", str1[i], str2[i], str1[i] == str2[i])
	// }
}
