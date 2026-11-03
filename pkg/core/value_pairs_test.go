//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestValuePairs — the seven even/odd Value pairs (six of the vector table, OPAQUE-BYTES of §3.14) start on an even C with the odd one right after (vector table)
package core

import (
	"strings"
	"testing"
)

func TestValuePairs(t *testing.T) {
	codes := map[string][]uint16{}
	for _, r := range loadVector(t) {
		k, name, _, l := vectorExpect(r.Name, r.OldKey)
		if r.New < 0 || k != KindValue || l != 0 {
			continue
		}
		_, c := SplitTarget(Word(r.New))
		codes[name] = append(codes[name], c)
	}
	for _, r := range loadS314(t) {
		if k, c := SplitTarget(r.Tag); k == KindValue {
			name := strings.TrimSuffix(strings.TrimSuffix(r.Name, "-EVEN"), "-ODD")
			codes[name] = append(codes[name], c)
		}
	}
	pairs := 0
	for name, cs := range codes {
		if len(cs) == 2 && (cs[0]&1 != 0 || cs[1] != cs[0]+1) {
			t.Fatalf("%s: %02X", name, cs)
		}
		if len(cs) == 2 {
			pairs++
		}
	}
	if pairs != 7 {
		t.Fatalf("pairs %d, want 7", pairs)
	}
}
