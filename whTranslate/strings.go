package whtranslate

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