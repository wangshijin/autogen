package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func UpperFirst(s string) string {
	return strings.Title(s)
}

func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.Replace(s, string(s[0]), strings.ToLower(string(s[0])), 1)
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func FileName(path string) string {
	Separator := string(os.PathSeparator)
	path = strings.TrimSuffix(path, Separator)
	sep := `/`
	if Separator == `\` {
		sep = `\\`
	}
	reg := regexp.MustCompile(fmt.Sprintf(`^.*%s`, sep))
	prefix := reg.FindString(path)
	path = strings.TrimPrefix(path, prefix)
	return path
}
