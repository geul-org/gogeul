//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeBind — a BIND round-trips through DecodeBind; a zero Edge Id or a broken Subject fails
package core

import "testing"

func TestEncodeBind(t *testing.T) {
	subj, _ := EncodeIssuerPointer([]Word{0, 0})
	r, err := EncodeBind(Bind{EdgeID: 3, Subject: subj, Name: []byte("geul-org")})
	if err != nil {
		t.Fatal(err)
	}
	b, err := DecodeBind(r)
	if err != nil || b.EdgeID != 3 || string(b.Name) != "geul-org" {
		t.Fatalf("%+v %v", b, err)
	}
	if _, err := EncodeBind(Bind{EdgeID: 0, Subject: subj}); err == nil {
		t.Fatal("Edge Id 0: want error")
	}
	if _, err := EncodeBind(Bind{EdgeID: 1, Subject: Record{Kind: KindEdge, Codec: 0x140}}); err == nil {
		t.Fatal("codec past the Edge zone: want error")
	}
}
