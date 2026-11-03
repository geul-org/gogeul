//ff:func feature=core type=decoder control=sequence
//ff:what vectorRoundTrip — a record built for a vector TAG encodes to that TAG and decodes back to its Kind, Codec and Length
package core

import "testing"

// vectorRoundTrip builds a record of Kind k and codec whose total length makes
// L = l (an ISSUER codec gets an Issuer and codec number inside that length,
// l 14 and 15 are reached through the Length extension), encodes it and checks
// the TAG and the decoded Kind, Codec and size.
func vectorRoundTrip(t *testing.T, w Word, k Kind, codec uint16, l int) {
	t.Helper()
	total := l + 1
	if l == 0xE {
		total = 16
	}
	if l == 0xF {
		total = 0x10001
	}
	r := Record{Kind: k, Codec: codec}
	iw := IssuerWidth(k, codec)
	if iw > 0 && total < 1+iw+FormatWords {
		return
	}
	if iw > 0 {
		r.Issuer = make([]Word, iw)
		r.Issuer[iw-1] = 1
		r.Format = 0x00010002
	}
	head := map[int]int{0xE: 2, 0xF: 3}[l]
	r.Payload = make([]Word, total-max(head, 1)-iw-FormatWords*min(iw, 1))
	ws, err := EncodeRecord(r)
	if err != nil || ws[0] != w || len(ws) != total {
		t.Fatalf("%04X: encode %04X… (%d words) %v", w, ws[0], len(ws), err)
	}
	got, n, err := DecodeRecord(ws)
	if err != nil || n != total || got.Kind != k || got.Codec != codec || got.Format != r.Format || len(got.Issuer) != iw {
		t.Fatalf("%04X: decode %+v %d %v", w, got.Kind, n, err)
	}
}
