package helpers

import "strings"

// TruncateString shortens s to at most maxLen characters. When the string is
// longer than maxLen it is cut and a trailing ellipsis ("...") is appended so
// that the total length does not exceed maxLen. A non-positive maxLen returns
// an empty string.
func TruncateString(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	if len(s) <= maxLen {
		return s
	}
	const ellipsis = "..."
	if maxLen <= len(ellipsis) {
		return s[:maxLen]
	}
	return s[:maxLen-len(ellipsis)] + ellipsis
}

// NormalizeIdentifier trims surrounding whitespace from an identifier and
// lower-cases it so that comparisons are case-insensitive. An identifier made
// up entirely of whitespace normalizes to an empty string.
func NormalizeIdentifier(id string) string {
	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return ""
	}
	return strings.ToLower(trimmed)
}
