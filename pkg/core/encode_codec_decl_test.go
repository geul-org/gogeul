//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeCodecDecl — the docs' LABEL Object example words
package core

import (
	"slices"
	"testing"
)

func TestEncodeCodecDecl(t *testing.T) {
	ws, _ := EncodeRecord(EncodeCodecDecl(CodecDecl{EdgeID: 9, Field: 4, Target: TargetWord(KindValue, CodecUTF8LanguageTextEven)}))
	if !slices.Equal(ws, []Word{0x09E5, 0x0051, 0x0009, 0x0051, 0x0004, 0x0580}) {
		t.Fatalf("%04X", ws)
	}
}
