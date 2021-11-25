package whtranslate

import (
	"fmt"
	"strings"
	"unicode"
)

const maxStringLen = 50

func formatString(x string) string {
	if x == "" {
		return "<empty>"
	}
	if len(x) > maxStringLen {
		return x[:maxStringLen] + "..."
	}
	return x
}

func formatBool(x bool) string {
	if x {
		return "true"
	}
	return "false"
}

func formatUpdateMessage(fieldName string, newContent string) string {
	return fmt.Sprintf("%s updated: is now `%s`\n", fieldName, newContent)
}

func formatStatementMessage(fieldName string, content string) string {
	return fmt.Sprintf("%s: `%s`\n", fieldName, content)
}

// snakeToReadable converts a snake_case string to a readable string.
func snakeToReadable(x string) string {
	y := strings.Split(x, "_")
	z := strings.Join(y, " ")
	return string(unicode.ToUpper(rune(z[0]))) + z[1:]
}