//ff:func feature=pack type=helper control=sequence
//ff:what TestNotFound — wraps ErrNotFound with the path
package pack

import (
	"errors"
	"strings"
	"testing"
)

func TestNotFound(t *testing.T) {
	err := notFound("a/b")
	if !errors.Is(err, ErrNotFound) || !strings.Contains(err.Error(), "a/b") {
		t.Fatal(err)
	}
}
