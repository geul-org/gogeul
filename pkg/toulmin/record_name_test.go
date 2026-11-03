//ff:func feature=toulmin type=helper control=sequence
//ff:what TestRecordName — direct codecs and format numbers name; unknown refs are empty
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordName(t *testing.T) {
	if RecordName(CodecRef{Codec: CodecDoOnce}) != "DO-ONCE" || RecordName(CodecRef{Codec: core.CodecEdgeIssuer2Word, Format: FormatLabel}) != "LABEL" {
		t.Fatal("name mismatch")
	}
	if RecordName(CodecRef{Codec: 0x66}) != "" || RecordName(CodecRef{Codec: core.CodecEdgeIssuer2Word, Format: 0xCFFF}) != "" {
		t.Fatal("unknown refs must be empty")
	}
}
