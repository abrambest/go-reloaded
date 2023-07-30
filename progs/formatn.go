package progs

import (
	"strings"
)

func FormatN(s string) string {
	s = strings.ReplaceAll(s, "\\n", " \\n ")
	s = strings.ReplaceAll(s, "\\N", " \\n ")

	strArr := strings.Split(s, "\\n")
	for i := range strArr {
		strArr[i] = strings.Trim(strArr[i], " ")
	}

	s = strings.Join(strArr, "\n")

	return s
}
