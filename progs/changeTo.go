package progs

import (
	"fmt"
	"strconv"
	"strings"
)

func ChangeTo(s string) string {

	strArr := strings.Split(s, "\n")

	for i := 0; i < len(strArr); i++ {
		strFields := strings.Fields(string(strArr[i]))
		for j := 0; j < len(strFields); j++ {
			if strFields[j] == "(cap)" {
				strFields[j-1] = strings.Title(strFields[j-1])
				strFields = append(strFields[:j], strFields[j+1:]...)
				strArr[i] = strings.Join(strFields, " ")

			}
			if strFields[j] == "(up)" {

				strFields[j-1] = strings.ToUpper(strFields[j-1])
				strFields = append(strFields[:j], strFields[j+1:]...)
				strArr[i] = strings.Join(strFields, " ")

			}
			if strFields[j] == "(low)" {

				strFields[j-1] = strings.ToLower(strFields[j-1])
				strFields = append(strFields[:j], strFields[j+1:]...)
				strArr[i] = strings.Join(strFields, " ")

			} else if strFields[j] == "(hex)" {

				hex, err := strconv.ParseInt(strFields[j-1], 16, 32)
				if err != nil {
					fmt.Println(err)
				}
				strFields[j-1] = strconv.Itoa(int(hex))

				strFields = append(strFields[:j], strFields[j+1:]...)
				strArr[i] = strings.Join(strFields, " ")

			} else if strFields[j] == "(bin)" {

				hex, err := strconv.ParseInt(strFields[j-1], 2, 32)
				if err != nil {
					fmt.Println(err)
				}
				strFields[j-1] = strconv.Itoa(int(hex))

				strFields = append(strFields[:j], strFields[j+1:]...)
				strArr[i] = strings.Join(strFields, " ")

			}

		}

	}
	s = strings.Join(strArr, "\n")

	return s
}
