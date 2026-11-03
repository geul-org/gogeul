//ff:func feature=core type=helper control=sequence
//ff:what TestRecordName — a known Value ISSUER format by its name, an unknown one by the ISSUER codec name, others by CodecName
package core

import "testing"

func TestRecordName(t *testing.T) {
	if got := RecordName(Record{Kind: KindValue, Codec: CodecValueIssuer2Word, Issuer: []Word{0, 2}, Format: 0x0700}); got != "COMPLEX64" {
		t.Fatalf("complex64: %q", got)
	}
	if got := RecordName(Record{Kind: KindValue, Codec: CodecValueIssuer2Word, Issuer: []Word{0, 9}, Format: 0x0700}); got != "VALUE-ISSUER-2WORD" {
		t.Fatalf("unknown format: %q", got)
	}
	if got := RecordName(Record{Kind: KindValue, Codec: CodecFraction}); got != "FRACTION" {
		t.Fatalf("fraction: %q", got)
	}
}
