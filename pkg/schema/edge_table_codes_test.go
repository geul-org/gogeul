//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what edgeTableCodes — the Edge codes of the plan2 name table (layout2.md), each with whether it is assigned (Reserved is not)
package schema

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

// edgeTableCodes reads the name table of layout2.md (migration.md §3.15) and
// keeps the Edge rows (TAG 0x0A0L~0x1BFL) as Edge Codec (code - 0xA0) → true
// for an assigned code, false for Reserved.
func edgeTableCodes(t *testing.T) map[uint16]bool {
	b, err := os.ReadFile("../core/testdata/migration1/layout2.md")
	if err != nil {
		t.Fatal(err)
	}
	out := map[uint16]bool{}
	for _, line := range strings.Split(string(b), "\n") {
		a := strings.Split(line, "|")
		if len(a) < 3 || !strings.HasPrefix(strings.TrimSpace(a[1]), "`0x") {
			continue
		}
		code, err := strconv.ParseUint(strings.Trim(strings.TrimSpace(a[1]), "`L"), 0, 16)
		if err == nil && code >= 0x0A0 && code < 0x1C0 {
			out[uint16(code-0xA0)] = strings.TrimSpace(a[2]) != "Reserved"
		}
	}
	return out
}
