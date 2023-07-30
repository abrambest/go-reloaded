package progs

import (
	"strings"
)

func Atoan(s string) string {
	arrSplit := strings.Split(s, "\n")

	for i := 0; i < len(arrSplit); i++ {
		str := strings.Fields(arrSplit[i])
		for j := 0; j < len(str); j++ {
			if str[j] == "A" && string(str[j+1][0]) == "a" ||
				str[j] == "A" && string(str[j+1][0]) == "e" ||
				str[j] == "A" && string(str[j+1][0]) == "i" ||
				str[j] == "A" && string(str[j+1][0]) == "o" ||
				str[j] == "A" && string(str[j+1][0]) == "u" ||
				str[j] == "A" && string(str[j+1][0]) == "h" {

				str[j] = "An"

			}
			if str[j] == "a" && string(str[j+1][0]) == "a" ||
				str[j] == "a" && string(str[j+1][0]) == "e" ||
				str[j] == "a" && string(str[j+1][0]) == "i" ||
				str[j] == "a" && string(str[j+1][0]) == "o" ||
				str[j] == "a" && string(str[j+1][0]) == "u" ||
				str[j] == "a" && string(str[j+1][0]) == "h" {

				str[j] = "an"

			}
			arrSplit[i] = strings.Join(str, " ")

		}

	}
	s = strings.Join(arrSplit, "\n")
	return s

}
