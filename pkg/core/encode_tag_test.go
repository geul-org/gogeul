//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeTag — direct/1-word/2-word Length boundaries and the negative body guard
package core

import "testing"

func TestEncodeTag(t *testing.T) {
	if w, _ := encodeTag(KindValue, CodecUTF8TextEven, 13); len(w) != 1 || w[0]&0xF != 13 {
		t.Fatalf("14 words must be direct: %04X", w)
	}
	if w, _ := encodeTag(KindValue, CodecUTF8TextEven, 14); len(w) != 2 || w[0]&0xF != 0xE || w[1] != 16 {
		t.Fatalf("16 words must use 1-word extension: %04X", w)
	}
	if w, _ := encodeTag(KindValue, CodecUTF8TextEven, 0xFFFD); len(w) != 2 || w[1] != 0xFFFF {
		t.Fatalf("65535 words must use 1-word extension: %04X", w)
	}
	if w, _ := encodeTag(KindValue, CodecUTF8TextEven, 0xFFFE); len(w) != 3 || w[0]&0xF != 0xF || w[1] != 1 || w[2] != 1 {
		t.Fatalf("65537 words must use 2-word extension: %04X", w)
	}
	if _, err := encodeTag(KindValue, CodecUTF8TextEven, 0xFFFFFFFF); err == nil {
		t.Fatal("record above 2^32-1 words must fail")
	}
	if _, err := encodeTag(KindValue, CodecUTF8TextEven, -1); err == nil {
		t.Fatal("negative body must fail")
	}
}
