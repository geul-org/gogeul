//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestEdgeDefinedTable — edgeDefined agrees with the plan2 name table over the whole Edge zone (layout2.md: 84 assigned defined, Reserved and empty codes not) and the §3.14 Edge rows (s314_p2.tsv)
package schema

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestEdgeDefinedTable(t *testing.T) {
	codes := edgeTableCodes(t)
	defined := 0
	for c := uint16(0); c <= 0x11F; c++ {
		want := codes[c]
		if edgeDefined(c) != want {
			t.Fatalf("Edge Codec 0x%03X (TAG 0x%03XL): table %v", c, c+0xA0, want)
		}
		if want {
			defined++
		}
	}
	if len(codes) != 88 || defined != 84 {
		t.Fatalf("%d table codes, %d assigned; want 88, 84", len(codes), defined)
	}
	s, err := os.ReadFile("../core/testdata/migration1/s314_p2.tsv")
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(strings.TrimSpace(string(s)), "\n")[1:] {
		a := strings.Split(line, "\t")
		tag, _ := strconv.ParseUint(strings.TrimSuffix(a[3], "L"), 0, 16)
		if a[0] == "Edge" && !edgeDefined(uint16(tag-0xA0)) {
			t.Fatalf("%s 0x%03XL: not defined", a[2], tag)
		}
	}
}
