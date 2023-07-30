package progs

import (
	"fmt"
	"strconv"
	"strings"
)

func NumChangeTo(s string) string {

	strN := strings.Split(s, "\n")

	strWithN := ""

	for n := 0; n < len(strN); n++ {
		if n != len(strN)-1 {
			strWithN += strN[n] + "\\n"
		} else {
			strWithN += strN[n]
		}

	}

	strFields := strings.Fields(strWithN)

	for i := 0; i < len(strFields); i++ {
		if strFields[i] == "(cap," {
			num, err := strconv.Atoi(string(strFields[i+1][0]))
			if err != nil {
				fmt.Println(err)
			}
			for j := i - num; j < i; j++ {
				strFields[j] = strings.Title(strFields[j])

			}
			strFields = append(strFields[:i], strFields[i+2:]...)
		}
		if strFields[i] == "(low," {
			num, err := strconv.Atoi(string(strFields[i+1][0]))
			if err != nil {
				fmt.Println(err)
			}
			for j := i - num; j < i; j++ {
				strFields[j] = strings.ToLower(strFields[j])

			}
			strFields = append(strFields[:i], strFields[i+2:]...)
		}
		if strFields[i] == "(up," {
			num, err := strconv.Atoi(string(strFields[i+1][0]))
			if err != nil {
				fmt.Println(err)
			}
			for j := i - num; j < i; j++ {
				strFields[j] = strings.ToUpper(strFields[j])

			}
			strFields = append(strFields[:i], strFields[i+2:]...)
		}
	}
	s = strings.Join(strFields, " ")

	return s
}
