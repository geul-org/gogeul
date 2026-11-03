//ff:func feature=toulmin type=helper control=sequence
//ff:what TestRecordName — direct codecs name, the §3.14 codes included; unknown codecs and EDGE-ISSUER are empty
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecordName(t *testing.T) {
	if RecordName(CodecRef{Codec: CodecDoOnce}) != "DO-ONCE" || RecordName(CodecRef{Codec: CodecLabel}) != "LABEL" ||
		RecordName(CodecRef{Codec: CodecSupersedeChain}) != "SUPERSEDE-CHAIN" || RecordName(CodecRef{Codec: CodecVerdict}) != "VERDICT" || RecordName(CodecRef{Codec: CodecEvidence}) != "EVIDENCE" {
		t.Fatal("name mismatch")
	}
	if RecordName(CodecRef{Codec: 0x66}) != "" || RecordName(CodecRef{Codec: core.CodecEdgeIssuer2Word}) != "" || RecordName(CodecRef{Codec: 0x95}) != "" {
		t.Fatal("unknown refs must be empty")
	}
}
