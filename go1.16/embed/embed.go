// Example of using embed from go1.16
package embed

import _ "embed"

//go:embed index.html
var indexHtml string

// returns the content of an embedded file
func New() string {
	return indexHtml
}
