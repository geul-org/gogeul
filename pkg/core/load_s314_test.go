//ff:func feature=core type=helper control=iteration dimension=1
//ff:what loadS314 — read testdata/migration1/s314.tsv into rows (TAG `0x05EL` read with L=0)
package core

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

// loadS314 reads the §3.14 table. A row that does not parse fails the test.
func loadS314(t *testing.T) []s314Row {
	b, err := os.ReadFile("testdata/migration1/s314.tsv")
	if err != nil {
		t.Fatal(err)
	}
	var rows []s314Row
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n")[1:] {
		a := strings.Split(line, "\t")
		f, e1 := strconv.ParseUint(a[1], 0, 32)
		tag, e2 := strconv.ParseUint(strings.TrimSuffix(a[3], "L")+"0", 0, 16)
		id, e3 := strconv.ParseUint(a[4], 0, 32)
		if len(a) != 5 || e1 != nil || e2 != nil || e3 != nil {
			t.Fatalf("bad row %q", line)
		}
		rows = append(rows, s314Row{Kind: a[0], Format: uint32(f), Name: a[2], Tag: Word(tag), CodecID: uint32(id)})
	}
	return rows
}
