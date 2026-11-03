//ff:func feature=inspect type=encoder control=sequence
//ff:what tombstoneSpec — test fixture: TOMBSTONE declared as a DOC-POINTER record and reason 1 (code 1 = license)
package inspect

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// tombstoneSpec declares TOMBSTONE for layout tests.
func tombstoneSpec() codecspec.Spec {
	return codecspec.Spec{core.TargetWord(core.KindEdge, core.EdgeTombstone): {
		{Name: "doc", Codecs: []core.Word{core.TargetWord(core.KindReference, core.CodecDocPointer)}}, {Name: "reason", Width: 1, Enums: map[string]string{"0x0001": "license"}},
	}}
}
