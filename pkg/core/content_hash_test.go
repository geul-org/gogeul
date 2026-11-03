//ff:func feature=core type=encoder control=sequence
//ff:what TestContentHash — signatures do not change the hash, content does, and a malformed document fails
package core

import (
	"crypto/sha256"
	"testing"
)

func TestContentHash(t *testing.T) {
	name, _ := EncodeRecord(Record{Kind: KindMeta, Codec: MetaDocID, Payload: []Word{0, 0, 0, 0, 1}})
	sig, _ := EncodeRecord(Record{Kind: KindMeta, Codec: MetaSignature, Payload: []Word{7, 7, 7}})
	val, _ := EncodeRecord(EncodeUintValue(5))
	plain := WrapDocument(append(append(append([]Word{}, name...), 0x1C30), append(val, 0x1C40)...))
	signed := WrapDocument(append(append(append(append([]Word{}, name...), sig...), 0x1C30), append(val, 0x1C40)...))
	h, err := ContentHash(plain)
	if err != nil || h != sha256.Sum256(PutWords(plain[:len(plain)-1])) {
		t.Fatalf("plain: %x %v", h, err)
	}
	if hs, err := ContentHash(signed); err != nil || hs != h {
		t.Fatalf("signed: %x %v", hs, err)
	}
	other, _ := EncodeRecord(EncodeUintValue(6))
	if ho, _ := ContentHash(WrapDocument(append(append(append([]Word{}, name...), 0x1C30), append(other, 0x1C40)...))); ho == h {
		t.Fatal("different content, same hash")
	}
	if _, err := ContentHash(plain[:len(plain)-1]); err == nil {
		t.Fatal("truncated: want error")
	}
}
