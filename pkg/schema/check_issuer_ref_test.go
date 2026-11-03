//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckIssuerRef — R4: EXTERNAL-ISSUER-2/4WORD payload length, canonical Issuer and a nonzero target; R2 on the target
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckIssuerRef(t *testing.T) {
	ref := func(p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: core.CodecExternalIssuer2Word, Payload: p})
	}
	s2 := core.SchemeExternalIssuer2Word
	if fs := checkIssuerRef(ref(0, 2, 0, 0, 0, 2), s2); rulesOf(fs) != "R4" {
		t.Fatalf("short: %+v", fs)
	}
	if fs := checkIssuerRef(ref(0, 2, 0, 0, 0, 2, 7), s2); fs != nil {
		t.Fatalf("ok: %+v", fs)
	}
	if fs := checkIssuerRef(ref(0, 2, 0, 0, 0, 2, 0), s2); rulesOf(fs) != "R4" {
		t.Fatalf("document itself: %+v", fs)
	}
	if fs := checkIssuerRef(ref(0, 2, 0, 0, 0, 0, 7), s2); fs != nil {
		t.Fatalf("document 0 (term index), target 7: %+v", fs)
	}
	if fs := checkIssuerRef(ref(0, 2, 0, 0, 0, 2, 0, 7), s2); rulesOf(fs) != "R2" {
		t.Fatalf("leading zero: %+v", fs)
	}
	ref4 := func(p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: core.CodecExternalIssuer4Word, Payload: p})
	}
	s4 := core.SchemeExternalIssuer4Word
	if fs := checkIssuerRef(ref4(0, 1, 0, 0x20, 0, 0, 0, 2, 7), s4); fs != nil {
		t.Fatalf("4-word: %+v", fs)
	}
	if fs := checkIssuerRef(ref4(0, 0, 0, 0x20, 0, 0, 0, 2, 7), s4); rulesOf(fs) != "R4" {
		t.Fatalf("4-word of a 2-word value: %+v", fs)
	}
}
