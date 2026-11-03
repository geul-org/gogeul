//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterPartition — P keys by Issuer number, H keys raw, named sorted
package pack

import (
	"strings"
	"testing"
)

func TestWriterPartition(t *testing.T) {
	w := &writer{inum: map[string]uint32{"0x00000000": 0}}
	h := "0x00000000/hashes/0x" + strings.Repeat("ab", 32) + ".geul"
	ps, hs, ns := w.partition(map[string]Loc{"0x00000000/docs/0x0000000000000008/versions.geul": {}, h: {Off: 1}, "b": {}, "a": {}})
	if len(ps) != 1 || ps[0].key != pKey(0, 8, 0) || len(hs) != 1 || hs[0].loc.Off != 1 || ns[0].path != "a" || ns[1].path != "b" {
		t.Fatal(ps, hs, ns)
	}
}
