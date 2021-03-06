// Package internal provides ...
package internal

import (
	"strings"
	"unicode"
)

// 全部转大写ToUpper
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 全部转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线转大写驼峰 "_"->" "->""
func UndersocreToUpperCameCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线转小写驼峰
func UndersocreToLowerCameCase(s string) string {
	s = UndersocreToUpperCameCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰转下划线
func CameCaseToUndersocre(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))

			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
