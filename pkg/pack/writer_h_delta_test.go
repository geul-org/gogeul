//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterHDelta — changed h/ paths with their records, deleted ones as markers, sorted; other spaces ignored
package pack

import (
	"strings"
	"testing"
)

func TestWriterHDelta(t *testing.T) {
	a, b := strings.Repeat("a", 32), strings.Repeat("b", 32)
	w := &writer{puts: []*put{{ref: pathRef{space: spaceH, hash: b}, loc: Loc{Off: 9}}, {path: "x"}}, dels: []string{hPath("0x00000000", a), "y"}}
	d := w.hDelta()
	if len(d) != 2 || d[0].key != a || d[0].loc.Enc != encGone || d[1].loc.Off != 9 {
		t.Fatal(d)
	}
}
