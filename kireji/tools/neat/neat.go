// Package neat implements value sanitisation functions.
package neat

import "strings"

// Body returns a sanitised file body string.
func Body(body string) string {
	body = strings.TrimLeft(body, "\n\r")
	return strings.TrimRight(body, "\n\r") + "\n"
}

// Extn returns a sanitised file extension string.
func Extn(extn string) string {
	extn = strings.TrimSpace(extn)
	extn = strings.ToLower(extn)
	return "." + strings.TrimLeft(extn, ".")
}

// Name returns a sanitised file base name string.
func Name(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	return strings.ReplaceAll(name, " ", "-")
}
