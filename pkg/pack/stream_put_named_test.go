//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamPutNamed — a closing stream acknowledges a held path and refuses a new one; a named record joins the pack's path table and NAMED without an h/ key; a P path and a failed place are errors
package pack

import "testing"

func TestStreamPutNamed(t *testing.T) {
	s := streamTestNew(t, nil)
	r, _ := s.Encode(streamTestDoc(1, 1))
	if err := s.PutNamed("0x00000001/docs/0x00000001.geul", r); err != nil || len(s.st.Named) != 1 || s.st.Paths[0] != "0x00000001/docs/0x00000001.geul" || s.st.Inv[0].Kind != invNamed || len(s.hbuf) != 0 || s.st.Res.Shared != 0 {
		t.Fatalf("%+v %v", s.st.Inv, err)
	}
	if err := s.PutNamed("0x00000001/docs/0x5100000000000001/versions.geul", r); err == nil {
		t.Fatal("P path: want error")
	}
	s.st.Close.Phase = closeStarted
	if err := s.PutNamed("0x00000001/docs/0x00000001.geul", r); err != nil || len(s.st.Named) != 1 {
		t.Fatalf("closing, held path: %v", err)
	}
	if err := s.PutNamed("issuers/0x00000001.geul", r); err == nil {
		t.Fatal("closing, new path: want error")
	}
	s.st.Close.Phase = closeNone
	s.data.Close()
	if err := s.PutNamed("issuers/0x00000001.geul", r); err == nil {
		t.Fatal("place failure: want error")
	}
}
