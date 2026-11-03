//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatRecord — ABSENT UNKNOWN NONE for the slot markers (a reserved one as raw words), short forms for Reference and Value, raw words otherwise
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatRecord(t *testing.T) {
	if formatRecord(absent()) != "ABSENT" || formatRecord(core.EncodeLocalRef(3)) != "#3" || formatRecord(core.EncodeUintValue(2)) != "UNSIGNED-INTEGER(2)" {
		t.Fatal("short forms")
	}
	if got := formatRecord(core.Record{Kind: core.KindReference, Codec: core.CodecGeul, Payload: []core.Word{5}}); got != "0x00A10005" {
		t.Fatal(got)
	}
	if got := formatRecord(core.Record{Kind: core.KindValue, Codec: core.CodecBinaryFloat, Payload: []core.Word{0}}); len(got) < 3 || got[:2] != "0x" {
		t.Fatal(got)
	}
	unknown, _, _ := core.DecodeRecord([]core.Word{core.UnknownValue})
	none, _, _ := core.DecodeRecord([]core.Word{core.NoValue})
	reserved, _, _ := core.DecodeRecord([]core.Word{0x0011, 7})
	if formatRecord(unknown) != "UNKNOWN" || formatRecord(none) != "NONE" || formatRecord(reserved) != "0x00110007" {
		t.Fatal("slot markers")
	}
	if got := formatRecord(core.EncodeSpecVersion(2)); got[:2] != "0x" {
		t.Fatal(got)
	}
}
