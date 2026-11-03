//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterPlanDeletes — listed paths the previous generation has; a Full batch deletes every unlisted path
package pack

import "testing"

func TestWriterPlanDeletes(t *testing.T) {
	_, out := built(t, small())
	w := testWriter(t, out, Batch{Delete: []string{"issuers/0x00000001.geul", "nope"}})
	w.sums = map[string][32]byte{}
	w.planDeletes()
	if len(w.dels) != 1 || w.dels[0] != "issuers/0x00000001.geul" {
		t.Fatal(w.dels)
	}
	w = testWriter(t, out, Batch{Full: true})
	w.sums = map[string][32]byte{"issuers/0x00000001.geul": {}}
	w.planDeletes()
	if len(w.dels) != 17 {
		t.Fatal(len(w.dels))
	}
}
