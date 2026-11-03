//ff:func feature=build type=helper control=sequence
//ff:what TestInvertNames — ID → name, the first name in order for an ID with several
package build

import (
	"testing"
)

func TestInvertNames(t *testing.T) {
	l := invertNames(map[string]string{"b": "1", "a": "1", "c": "2"})
	if len(l) != 2 || l["1"] != "a" || l["2"] != "c" {
		t.Fatalf("got %v", l)
	}
}
