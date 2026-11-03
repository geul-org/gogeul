//ff:func feature=build type=renderer control=sequence
//ff:what TestIndexJSON — the listing as the tree browser's JSON; a missing directory fails
package build

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestIndexJSON(t *testing.T) {
	root := treeFixture(t)
	var b bytes.Buffer
	if err := indexJSON(&b, treeOf(root), "/.well-known/geul/h/"); err != nil {
		t.Fatal(err)
	}
	var p indexPage
	if err := json.Unmarshal(b.Bytes(), &p); err != nil || p.Up != "/.well-known/geul/" || len(p.Entries) != 2 || p.Entries[0].URL != "/.well-known/geul/h/a" {
		t.Fatalf("%+v %v", p, err)
	}
	if err := indexJSON(&b, treeOf(root), "/.well-known/geul/none/"); err == nil {
		t.Fatal("missing: want error")
	}
}
