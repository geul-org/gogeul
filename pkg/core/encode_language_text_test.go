//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeLanguageText — the language record then the text; odd text bytes take the odd codec (the docs' "Paris"@fr example)
package core

import (
	"slices"
	"testing"
)

func TestEncodeLanguageText(t *testing.T) {
	fr := EncodeBytesValue(CodecLanguageBCP47Even, []byte("fr"))
	ws, _ := EncodeRecord(EncodeLanguageText(fr, []byte("Paris")))
	if !slices.Equal(ws, []Word{0x0595, 0x0561, 0x6672, 0x5061, 0x7269, 0x7300}) {
		t.Fatalf("got %04X", ws)
	}
	ko := EncodeBytesValue(CodecLanguageBCP47Even, []byte("ko"))
	if ws, _ := EncodeRecord(EncodeLanguageText(ko, nil)); !slices.Equal(ws, []Word{0x0582, 0x0561, 0x6B6F}) {
		t.Fatalf("empty: %04X", ws)
	}
}
