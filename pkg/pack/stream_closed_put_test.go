//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamClosedPut — keys up to the last are acknowledged; a later key, or any key with no record, is an error
package pack

import "testing"

func TestStreamClosedPut(t *testing.T) {
	s := streamTestNew(t, nil)
	if err := s.closedPut(pKey(0, 1, 0)); err == nil {
		t.Fatal("no record: want error")
	}
	streamTestPut(t, s, 1, 2)
	if err := s.closedPut(pKey(0, 0x5100000000000002, 1)); err != nil {
		t.Fatalf("last key: %v", err)
	}
	if err := s.closedPut(pKey(0, 0x5100000000000003, 0)); err == nil {
		t.Fatal("later key: want error")
	}
}
