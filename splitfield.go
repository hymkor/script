package main

import (
	"strings"
	"unicode/utf8"
)

func splitField(s string) (fields []string) {
	for len(s) > 0 {
		s = strings.TrimLeft(s, " \t")
		var field strings.Builder
		quote := false
		for len(s) > 0 {
			c, siz := utf8.DecodeRuneInString(s)
			if !quote && (c == ' ' || c == '\t') {
				break
			}
			if c == '"' {
				quote = !quote
			} else {
				field.WriteRune(c)
			}
			s = s[siz:]
		}
		fields = append(fields, field.String())
	}
	return
}
