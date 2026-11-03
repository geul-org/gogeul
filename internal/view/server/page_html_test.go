//ff:func feature=build type=renderer control=sequence
//ff:what TestPageHTML — the viewer page with the tree browser and the named title; bytes that are not a document fail
package server

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPageHTML(t *testing.T) {
	root := namesFixture(t)
	data, _ := os.ReadFile(filepath.Join(root, ".well-known", "geul", "0x00000000", "docs", "0xffff000000000002", "0x00000001.geul"))
	var b bytes.Buffer
	if err := PageHTML(&b, treeOf(root), "/.well-known/geul/0x00000000/docs/0xffff000000000002/0x00000001", data, nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(b.String(), "browseBtn") || !strings.Contains(b.String(), ".well-known/geul/geul.org/docs/0xffff000000000002/0x00000001") {
		t.Fatalf("page: %s", b.String()[:200])
	}
	if err := PageHTML(&b, treeOf(root), "/x", []byte{1, 2}, nil); err == nil {
		t.Fatal("bad: want error")
	}
}
