//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestRawEdge — an Edge's Edge Id and payload; bad words, a non-Edge and an Edge without a LOCAL Edge Id fail
package notation

import "testing"

func TestRawEdge(t *testing.T) {
	if e, err := rawEdge("0x0A43005100070001"); err != nil || e.Label != "7" || len(e.Rest) != 1 {
		t.Fatalf("%+v %v", e, err)
	}
	for _, s := range []string{"0x1", "0x00510001", "0x0A420141"} {
		if _, err := rawEdge(s); err == nil {
			t.Errorf("%s: want error", s)
		}
	}
}
