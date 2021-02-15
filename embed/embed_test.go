package embed

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	data := New()
	if !strings.Contains(data, "<html>") {
		t.Errorf("embed is missing <html> %s", data)
	}
}
