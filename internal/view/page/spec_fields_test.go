//ff:func feature=page type=builder control=sequence
//ff:what TestSpecFields — declared fields by name without the Edge Id; no spec falls back to raw fields
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

func TestSpecFields(t *testing.T) {
	r := core.EncodePublication(core.Publication{Codec: core.EdgeTombstone, EdgeID: 3, Issuer: []core.Word{0, 0}, Doc: []core.Word{0, 0, 0, 1}, Version: 2, Reason: 1})
	spec := []codecspec.Field{{Name: "doc", Codecs: []core.Word{core.TargetWord(core.KindReference, core.CodecDocPointer)}}, {Name: "reason", Width: 1, Enums: map[string]string{"0x0001": "license"}}}
	fs := specFields(inspectEntry(r, spec))
	if len(fs) != 2 || fs[0].Key != "doc" || fs[1].Text != "license" {
		t.Fatalf("got %+v", fs)
	}
	if fs := specFields(inspectEntry(core.Record{Issuer: []core.Word{1}}, nil)); len(fs) != 2 || fs[0].Key != "issuer" {
		t.Fatalf("raw: %+v", fs)
	}
}
