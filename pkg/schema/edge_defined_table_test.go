//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestEdgeDefinedTable — edgeDefined agrees with the plan's Edge table (t12_edge.md: assigned rows defined, 빈칸 and deleted rows not) and the §3.14 Edge rows (s314.tsv)
package schema

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestEdgeDefinedTable(t *testing.T) {
	codes := edgeTableCodes(t)
	if len(codes) != 0x7C {
		t.Fatalf("%d table codes, want 0x7C", len(codes))
	}
	for c, want := range codes {
		if edgeDefined(c) != want {
			t.Fatalf("Edge C 0x%02X: table %v", c, want)
		}
	}
	s, err := os.ReadFile("../core/testdata/migration1/s314.tsv")
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(strings.TrimSpace(string(s)), "\n")[1:] {
		a := strings.Split(line, "\t")
		tag, _ := strconv.ParseUint(strings.TrimSuffix(a[3], "L"), 0, 16)
		if a[0] == "Edge" && !edgeDefined(uint16(tag-0x80)) {
			t.Fatalf("%s 0x%03XL: not defined", a[2], tag)
		}
	}
}
