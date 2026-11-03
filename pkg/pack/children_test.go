//ff:func feature=pack type=resolver control=sequence
//ff:what TestChildren — direct files and directories, each directory once, at the top and below
package pack

import (
	"reflect"
	"testing"
)

func TestChildren(t *testing.T) {
	paths := []string{"0x00000000/a/v.geul", "0x00000000/a/w.geul", "0x00000000/b.geul", "h/x.geul", "x.geul"}
	if got := children(paths, ""); !reflect.DeepEqual(got, []dirEntry{{name: "0x00000000", dir: true}, {name: "h", dir: true}, {name: "x.geul"}}) {
		t.Fatal(got)
	}
	if got := children(paths, "0x00000000/"); !reflect.DeepEqual(got, []dirEntry{{name: "a", dir: true}, {name: "b.geul"}}) {
		t.Fatal(got)
	}
	if got := children(paths, "zz/"); got != nil {
		t.Fatal(got)
	}
}
