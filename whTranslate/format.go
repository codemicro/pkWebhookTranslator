package whtranslate

import "fmt"

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

func formatUpdateMessage(fieldName string, newContent string) string {
	return fmt.Sprintf("%s updated: is now `%s`\n", fieldName, newContent)
}

func formatStatementMessage(fieldName string, content string) string {
	return fmt.Sprintf("%s: `%s`\n", fieldName, content)
}