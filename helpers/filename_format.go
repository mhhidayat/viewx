package helpers

import "strings"

func FormatFileName(filePath string) string {
	if strings.HasSuffix(filePath, ".html") {
		return filePath
	}
	return filePath + ".html"
}
