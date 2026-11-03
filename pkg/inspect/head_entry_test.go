//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestHeadEntry — each known Meta gets typed Data; unknown and malformed Meta and HEAD Edges stay bare
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestHeadEntry(t *testing.T) {
	lw, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("en")))
	at, _ := core.EncodeRecord(core.EncodePosixSeconds(5))
	typed := []core.Record{
		{Kind: core.KindMeta, Codec: core.MetaDefaultLanguage, Payload: lw},
		{Kind: core.KindMeta, Codec: core.MetaRecordedAt, Payload: at},
		func() core.Record { r, _ := core.EncodeDocID([]core.Word{0, 2}, []core.Word{0, 0, 0, 2}); return r }(),
		{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0x0418, 0, 1, 0, 0, 0, 2, 0, 0, 4}},
	}
	for _, r := range typed {
		if e := headEntry(r); e.Data == nil || e.Name == "" {
			t.Fatalf("0x%02X: %+v", r.Codec, e)
		}
	}
	if e := headEntry(core.Record{Kind: core.KindMeta, Codec: core.MetaSpecVersion}); e.Data != nil || e.Name != "SPEC-VERSION" {
		t.Fatalf("malformed: %+v", e)
	}
	if e := headEntry(core.Record{Kind: core.KindMeta, Codec: 0x2A, Payload: []core.Word{1}}); e.Data != nil || e.Name != "" {
		t.Fatalf("unknown: %+v", e)
	}
	src := core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Payload: append([]core.Word{0x0051, 1, 0x0051, 0}, at...)}
	if e := headEntry(src); e.Data != nil || e.Name != "SOURCE" {
		t.Fatalf("HEAD SOURCE stays a bare Edge: %+v", e)
	}
}
