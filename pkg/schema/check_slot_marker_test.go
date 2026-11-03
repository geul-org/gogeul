//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckSlotMarker — slot markers as records are Errors; a marker C with payload is not reported (§1.1 112); pointers are not markers
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckSlotMarker(t *testing.T) {
	ref := func(codec uint16, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindReference, Codec: codec, Payload: p})
	}
	cases := []struct {
		r    rec
		sev  Severity
		part string
	}{
		{ref(0), Error, "자리 표지"},
		{ref(1), Error, "자리 표지"},
		{ref(2), Error, "자리 표지"},
	}
	for i, c := range cases {
		fs := checkReference(c.r, nil)
		if len(fs) != 1 || fs[0].Rule != "R1" || fs[0].Severity != c.sev || fs[0].Where.Part != c.part {
			t.Fatalf("case %d: %+v", i, fs)
		}
	}
	for _, r := range []rec{ref(2, 1), ref(0, 1, 2)} {
		if fs := checkReference(r, nil); fs != nil {
			t.Fatalf("marker with payload reported: %+v", fs)
		}
	}
	for _, c := range []uint16{core.CodecIssuerPointer, core.CodecDocPointer} {
		if got := rulesOf(checkReference(ref(c, 0, 0), nil)); got == "R1" {
			t.Fatalf("C %d read as a marker: %q", c, got)
		}
	}
}
