//ff:func feature=core type=helper control=iteration dimension=1
//ff:what loadLayout — read testdata/migration1/layout2.md: the section table (section → Kind, title) and the name table (TAG ID → name)
package core

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

// loadLayout reads the two tables of layout2.md. A section row is
// "| 3 `0011` | `0x060L~0x07FL` | Value | …", a name row "| `0x060L` | V:BOOL |".
func loadLayout(t *testing.T) layoutTable {
	b, err := os.ReadFile("testdata/migration1/layout2.md")
	if err != nil {
		t.Fatal(err)
	}
	kinds := map[string]Kind{"Reference": KindReference, "Value": KindValue, "Edge": KindEdge, "Meta": KindMeta}
	l := layoutTable{names: map[uint16]string{}}
	sections := 0
	for _, line := range strings.Split(string(b), "\n") {
		a := strings.Split(line, "|")
		if len(a) < 4 {
			continue
		}
		c0, c1 := strings.TrimSpace(a[1]), strings.TrimSpace(a[2])
		if code, err := strconv.ParseUint(strings.Trim(c0, "`L"), 0, 16); err == nil && strings.HasPrefix(c0, "`0x") {
			l.names[uint16(code)] = c1
			continue
		}
		s, err := strconv.Atoi(strings.Fields(c0 + " x")[0])
		title := strings.TrimSpace(a[3])
		if k, ok := kinds[strings.Fields(title + " x")[0]]; err == nil && ok && s >= 0 && s < 16 {
			l.kind[s], l.title[s] = k, title
			sections++
		}
	}
	if sections != 16 || len(l.names) != 231 {
		t.Fatalf("layout2.md: %d sections, %d names", sections, len(l.names))
	}
	return l
}
