//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeBytesValue — even/odd codec selection and the 0x00 pad byte
package core

import "testing"

func TestEncodeBytesValue(t *testing.T) {
	r := EncodeBytesValue(CodecUTF8TextEven, []byte("ab"))
	if r.Codec != 0x12 || len(r.Payload) != 1 || r.Payload[0] != 0x6162 {
		t.Fatalf("even: %+v", r)
	}
	r = EncodeBytesValue(CodecUTF8TextEven, []byte("abc"))
	if r.Codec != 0x13 || len(r.Payload) != 2 || r.Payload[1] != 0x6300 {
		t.Fatalf("odd: %+v", r)
	}
	if r = EncodeBytesValue(CodecUTF8TextEven, nil); r.Codec != 0x12 || len(r.Payload) != 0 {
		t.Fatalf("empty: %+v", r)
	}
}
