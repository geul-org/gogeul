//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestBindDoc — the document of a version-0 DOC-POINTER; a version, the old EXTERNAL-ISSUER form, other codecs and kinds fail
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestBindDoc(t *testing.T) {
	if d, ok := bindDoc(subjectOf(2, 0)); !ok || d != "0xffff000000000002" {
		t.Fatalf("ok: %q %v", d, ok)
	}
	bad := []core.Record{
		subjectOf(2, 1),
		core.EncodeLocalRef(0),
		{Kind: core.KindValue, Codec: core.CodecDocPointer, Payload: []core.Word{0, 0, 0xffff, 0, 0, 2, 0, 0}},
		{Kind: core.KindReference, Codec: core.CodecExternalIssuer2Word, Payload: []core.Word{0, 0, 0xffff, 0, 0, 2, 0}},
	}
	for i, r := range bad {
		if _, ok := bindDoc(r); ok {
			t.Fatalf("bad %d: want false", i)
		}
	}
}
