//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeBind — a BIND decodes, with a Reference or a codec TAG Subject; wrong codec, bad Edge Id, other Subjects, missing or trailing name and a non-text name fail
package core

import "testing"

func TestDecodeBind(t *testing.T) {
	enc := func(r Record) []Word { ws, _ := EncodeRecord(r); return ws }
	id, subj, name := enc(EncodeLocalRef(3)), enc(EncodeLocalRef(7)), enc(EncodeBytesValue(CodecUTF8TextEven, []byte("main")))
	cat := func(parts ...[]Word) []Word {
		var out []Word
		for _, p := range parts {
			out = append(out, p...)
		}
		return out
	}
	b, err := DecodeBind(Record{Kind: KindEdge, Codec: EdgeBind, Payload: cat(id, subj, name)})
	if err != nil || b.EdgeID != 3 || b.Subject.Payload[0] != 7 || string(b.Name) != "main" {
		t.Fatalf("got %+v %v", b, err)
	}
	if b, err := DecodeBind(Record{Kind: KindEdge, Codec: EdgeBind, Payload: cat(id, []Word{0x0980}, name)}); err != nil || b.Subject.Codec != EdgePublication {
		t.Fatalf("codec subject: %+v %v", b, err)
	}
	bad := [][]Word{
		nil,
		cat(enc(EncodeUintValue(1)), subj, name),
		cat(id),
		cat(id, enc(EncodeUintValue(1)), name),
		cat(id, subj),
		cat(id, subj, name, []Word{0x1C40}),
		cat(id, subj, enc(EncodeUintValue(1))),
	}
	for i, p := range bad {
		if _, err := DecodeBind(Record{Kind: KindEdge, Codec: EdgeBind, Payload: p}); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
	if _, err := DecodeBind(Record{Kind: KindEdge, Codec: EdgePublication}); err == nil {
		t.Fatal("codec: want error")
	}
}
