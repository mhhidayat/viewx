package helpers

import "strings"

func FormatFileName(filePath string) string {
	return strings.ReplaceAll(strings.TrimSuffix(filePath, ".html"), ".", "/") + ".html"
}
