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
	return fmt.Sprintf("%s updated: now `%s`\n", fieldName, newContent)
}

func formatColourURL(colour string) string {
	return fmt.Sprintf("https://fakeimg.pl/256x256/%s/?text=%%20", colour)
}
