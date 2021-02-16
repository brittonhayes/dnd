package utils

import "strings"

func Sanitize(name string) string {
	name = strings.TrimPrefix(name, "/")
	return strings.ReplaceAll(strings.ToLower(name), " ", "-")
}
