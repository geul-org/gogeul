//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what edgeTableCodes — the old Edge C 0x00~0x7B of t12_edge.md, each with whether the plan assigns it a new code
package schema

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

// edgeTableCodes reads the plan's Edge table and expands its ranges: a row
// whose new code column is a number is assigned, 빈칸 and deleted rows are not.
func edgeTableCodes(t *testing.T) map[uint16]bool {
	b, err := os.ReadFile("../core/testdata/migration1/t12_edge.md")
	if err != nil {
		t.Fatal(err)
	}
	out := map[uint16]bool{}
	for _, line := range strings.Split(string(b), "\n") {
		a := strings.Split(line, "|")
		if len(a) < 6 || !strings.HasPrefix(strings.TrimSpace(a[1]), "`0x") {
			continue
		}
		lo, hi, _ := strings.Cut(strings.Trim(strings.TrimSpace(a[1]), "`"), "`~`")
		from, _ := strconv.ParseUint(lo, 0, 16)
		to := from
		if hi != "" {
			to, _ = strconv.ParseUint(hi, 0, 16)
		}
		for c := from; c <= to && c < 0x7C; c++ {
			out[uint16(c)] = strings.HasPrefix(strings.TrimSpace(a[4]), "`0x")
		}
	}
	return out
}
