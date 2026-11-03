//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeLanguageText — round trip of language and text; another codec, a broken language record and bad padding fail
package core

import "testing"

func TestDecodeLanguageText(t *testing.T) {
	short := Record{Kind: KindValue, Codec: CodecLanguageShortCode, Payload: []Word{0x0783}}
	lang, text, err := DecodeLanguageText(EncodeLanguageText(short, []byte("DNA")))
	if err != nil || lang.Codec != CodecLanguageShortCode || lang.Payload[0] != 0x0783 || string(text) != "DNA" {
		t.Fatalf("%+v %q %v", lang, text, err)
	}
	for i, r := range []Record{
		EncodeBytesValue(CodecUTF8TextEven, []byte("ab")),
		{Kind: KindValue, Codec: CodecUTF8LanguageTextEven, Payload: []Word{0x056F}},
		{Kind: KindValue, Codec: CodecUTF8LanguageTextEven | 1, Payload: []Word{0x05A1, 0x0783, 0x4142}},
	} {
		if _, _, err := DecodeLanguageText(r); err == nil {
			t.Fatalf("case %d: want error", i)
		}
	}
}
