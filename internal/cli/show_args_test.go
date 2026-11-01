//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestShowArgs — one argument, or versions and a document; nothing else
package cli

import "testing"

func TestShowArgs(t *testing.T) {
	for _, args := range [][]string{{"a/b"}, {"versions"}, {"versions", "a/b"}} {
		if err := showArgs(nil, args); err != nil {
			t.Fatal(args, err)
		}
	}
	for _, args := range [][]string{nil, {"a/b", "c/d"}, {"versions", "a/b", "c/d"}} {
		if err := showArgs(nil, args); err == nil {
			t.Fatal(args)
		}
	}
}
