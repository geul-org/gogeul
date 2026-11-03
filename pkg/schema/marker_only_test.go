//ff:func feature=schema type=rule control=sequence
//ff:what TestMarkerOnly — markers with another target pass; 0000 alone or markers alone are E5; non-CODEC records are skipped
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestMarkerOnly(t *testing.T) {
	codec := func(id, f uint64, tg core.Word) core.Record {
		return core.EncodeCodecDecl(core.CodecDecl{EdgeID: id, Field: f, Target: tg})
	}
	ok := []rec{recOf(core.EncodeLocalRef(1)), recOf(codec(4, 2, 0x05C0)), recOf(codec(5, 2, core.Absent)), recOf(codec(6, 3, core.UnknownValue)), recOf(codec(7, 3, 0x000F))}
	if fs := markerOnly(ok); len(fs) != 0 {
		t.Fatalf("ok: %+v", fs)
	}
	if got := rulesOf(markerOnly([]rec{recOf(codec(4, 2, core.Absent)), recOf(codec(5, 3, 0x05C0))})); got != "E5" {
		t.Fatalf("0000 alone: %q", got)
	}
	if got := rulesOf(markerOnly([]rec{recOf(codec(4, 2, core.UnknownValue)), recOf(codec(5, 2, core.NoValue)), recOf(codec(6, 3, core.Absent))})); got != "E5,E5" {
		t.Fatalf("markers alone: %q", got)
	}
}
