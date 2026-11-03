//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriterPlanPut — invalid paths, a Get error, unreadable previous records, unchanged (nil) and changed puts
package pack

import (
	"os"
	"strings"
	"testing"
)

func TestWriterPlanPut(t *testing.T) {
	_, out := built(t, small())
	doc := fixtureFile(t, "issuers/0x00000001.geul")
	get := memGet(map[string][]byte{"issuers/0x00000001.geul": doc, "0x00000000/codecs/0x00000420/0x00000001.geul": {1, 2}})
	w := testWriter(t, out, Batch{Get: get})
	w.sums = map[string][32]byte{}
	for _, p := range []string{"", strings.Repeat("a", 0x10000), "none"} {
		if _, err := w.planPut(p); err == nil {
			t.Fatal(len(p))
		}
	}
	if _, err := w.planPut("issuers/0x00000002.geul"); err == nil {
		t.Fatal("Get error: want error")
	}
	if pu, err := w.planPut("issuers/0x00000001.geul"); pu != nil || err != nil {
		t.Fatal(pu, err)
	}
	pu, err := w.planPut("0x00000000/codecs/0x00000420/0x00000001.geul")
	if err != nil || pu == nil || pu.chOK || pu.ref.space != spaceNamed {
		t.Fatal(pu, err)
	}
	os.Remove(objPath(t, w.prev, out, kindData))
	if _, err := w.planPut("issuers/0x00000001.geul"); err == nil {
		t.Fatal("unreadable previous record")
	}
}
