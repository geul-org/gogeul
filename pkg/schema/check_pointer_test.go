//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckPointer — R3: IMPORT ID width by scheme (6/7/8 → 1/2/4 WORD), declared IMPORT ID, nonzero target; R2 on the target
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckPointer(t *testing.T) {
	ref := func(scheme uint16, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: core.RefCodec(scheme, 0), Payload: p})
	}
	p1, p2, p4 := core.SchemeExternalPointer1Word, core.SchemeExternalPointer2Word, core.SchemeExternalPointer4Word
	imports := map[uint64]bool{1: true, 0x10000: true}
	if fs := checkPointer(ref(p1, 1, 5), p1, imports); fs != nil {
		t.Fatalf("ok: %+v", fs)
	}
	if fs := checkPointer(ref(p1, 1), p1, imports); rulesOf(fs) != "R3" {
		t.Fatalf("short: %+v", fs)
	}
	if fs := checkPointer(ref(p2, 0, 1, 5), p2, imports); rulesOf(fs) != "R3" {
		t.Fatalf("wide: %+v", fs)
	}
	if fs := checkPointer(ref(p2, 1, 0, 5), p2, imports); fs != nil {
		t.Fatalf("2-word: %+v", fs)
	}
	if fs := checkPointer(ref(p4, 0, 0, 1, 0, 5), p4, imports); rulesOf(fs) != "R3" {
		t.Fatalf("4-word of a 2-word value: %+v", fs)
	}
	if fs := checkPointer(ref(p1, 2, 0, 5), p1, imports); rulesOf(fs) != "R3,R2" {
		t.Fatalf("missing + target: %+v", fs)
	}
	if fs := checkPointer(ref(p1, 1, 0), p1, imports); rulesOf(fs) != "R3" {
		t.Fatalf("document itself: %+v", fs)
	}
}
