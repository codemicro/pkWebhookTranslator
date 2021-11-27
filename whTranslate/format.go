/*
 *  pkWebhookTranslate, https://github.com/codemicro/pkWebhookTranslate
 *  Copyright (c) 2021 codemicro and contributors
 *
 *  SPDX-License-Identifier: MIT
 */

package whtranslate

import (
	"fmt"
	"strings"
	"time"
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

func formatTime(x time.Time) string {
	return x.Format(time.RFC1123)
}

func formatStringSlice(x []string) string {
	for i, y := range x {
		x[i] = formatString(y)
	}
	return fmt.Sprintf("[%s]", strings.Join(x, ", "))
}

func formatUpdateMessage(fieldName string, newContent string) string {
	return fmt.Sprintf("%s updated: is now `%s`\n", fieldName, newContent)
}

func formatStatementMessage(fieldName string, content string) string {
	return fmt.Sprintf("%s: `%s`\n", fieldName, content)
}

// snakeToReadable converts a snake_case string to a readable string with the first letter capitalised, and all other
// letters as they originally were.
func snakeToReadable(x string) string {
	y := strings.Split(x, "_")
	z := strings.Join(y, " ")
	return string(unicode.ToUpper(rune(z[0]))) + z[1:]
}
