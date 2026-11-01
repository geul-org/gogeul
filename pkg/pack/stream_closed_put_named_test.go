//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamClosedPutNamed — a held path is acknowledged, a new one is an error
package pack

import "testing"

func TestStreamClosedPutNamed(t *testing.T) {
	s := streamTestNew(t, nil)
	s.st.Named = []streamKV{{Key: "a.geul"}, {Key: "b.geul"}}
	if err := s.closedPutNamed("b.geul"); err != nil {
		t.Fatal(err)
	}
	if err := s.closedPutNamed("c.geul"); err == nil {
		t.Fatal("new path: want error")
	}
}
