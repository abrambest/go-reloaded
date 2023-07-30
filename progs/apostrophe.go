package progs

import (
	"strings"
)

func Apostrophe(s string) string {
	s = strings.ReplaceAll(s, "'", " ' ")

	strArr := strings.Fields(s)
	var tmpArr []string

	ap := false

	for i := 0; i < len(strArr); i++ {
		if strArr[i] != "'" {
			tmpArr = append(tmpArr, strArr[i])
		}
		if strArr[i] == "'" {
			if !ap {
				ap = true
				strArr[i+1] = strArr[i] + strArr[i+1]
			} else {
				tmpArr[len(tmpArr)-1] = tmpArr[len(tmpArr)-1] + strArr[i]
				ap = false
			}
		}

		s = strings.Join(tmpArr, " ")

	}
	return s
}
