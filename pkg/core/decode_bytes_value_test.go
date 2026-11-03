//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeBytesValue — round trip (a leading U+FEFF kept) plus wrong codec and bad pad errors
package core

import "testing"

func TestDecodeBytesValue(t *testing.T) {
	for _, s := range []string{"", "a", "ab", "한글", "\ufeff", "\ufeffAbecásia (GE)"} {
		b, err := DecodeBytesValue(EncodeBytesValue(CodecUTF8TextEven, []byte(s)), CodecUTF8TextEven)
		if err != nil || string(b) != s {
			t.Fatalf("%q: got %q err=%v", s, b, err)
		}
	}
	if _, err := DecodeBytesValue(Record{Kind: KindValue, Codec: CodecURIEven}, CodecUTF8TextEven); err == nil {
		t.Fatal("wrong codec must fail")
	}
	if _, err := DecodeBytesValue(Record{Kind: KindValue, Codec: CodecUTF8TextEven | 1, Payload: []Word{0x6162}}, CodecUTF8TextEven); err == nil {
		t.Fatal("non-zero pad must fail")
	}
	if _, err := DecodeBytesValue(Record{Kind: KindValue, Codec: CodecUTF8TextEven | 1}, CodecUTF8TextEven); err == nil {
		t.Fatal("odd without payload must fail")
	}
}
