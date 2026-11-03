//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestValuePairs — the six even/odd Value pairs start on an even C with the odd one right after (vector table)
package core

import "testing"

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
	pairs := 0
	for name, cs := range codes {
		if len(cs) == 2 && (cs[0]&1 != 0 || cs[1] != cs[0]+1) {
			t.Fatalf("%s: %02X", name, cs)
		}
		if len(cs) == 2 {
			pairs++
		}
	}
	if pairs != 6 {
		t.Fatalf("pairs %d, want 6", pairs)
	}
}
