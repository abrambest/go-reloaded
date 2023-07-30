package progs

import (
	"strings"
)

func Punctuation(s string) string {

	s = strings.ReplaceAll(s, ".", ". ")
	s = strings.ReplaceAll(s, ",", ", ")
	s = strings.ReplaceAll(s, "!", "! ")
	s = strings.ReplaceAll(s, "?", "? ")
	s = strings.ReplaceAll(s, ":", ": ")
	s = strings.ReplaceAll(s, ";", "; ")

	arrSplit := strings.Split(s, "\n")

	for k := 0; k < len(arrSplit); k++ {

		s := strings.Join(strings.Fields(arrSplit[k]), " ")

		for i := 0; i < len(s); i++ {

			if s[i] == ',' && s[i-1] == ' ' || s[i] == '.' && s[i-1] == ' ' || s[i] == '!' && s[i-1] == ' ' || s[i] == '?' && s[i-1] == ' ' || s[i] == ';' && s[i-1] == ' ' || s[i] == ':' && s[i-1] == ' ' {
				s = s[:i-1] + string(s[i]) + s[i+1:]
			}

		}
		arrSplit[k] = s

	}

	s = strings.Join(arrSplit, "\n")
	s = strings.ReplaceAll(s, "( ", "(")
	s = strings.ReplaceAll(s, " )", ") ")
	s = strings.ReplaceAll(s, ")", ") ")

	return s

}
