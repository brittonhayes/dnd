package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func sanitize(name string) string {
	name = strings.TrimPrefix(name, "/")
	return strings.ReplaceAll(strings.ToLower(name), " ", "-")
}

func URL(base string, endpoint string, name string) string {
	return fmt.Sprintf("%s%s/%s", base, endpoint, sanitize(name))
}

func URLWithQuery(base string, endpoint string, name string, values url.Values) string {
	return fmt.Sprintf("%s%s/%s?%s", base, endpoint, sanitize(name), values.Encode())
}
