//ff:func feature=notation type=encoder control=sequence
//ff:what TestCodecOf — Edge, Meta, Value (even member) and Reference names, any case; empty and unknown names are false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCodecOf(t *testing.T) {
	if c, ok := codecOf(core.KindEdge, "triple"); !ok || c != 0x0A {
		t.Fatal("TRIPLE")
	}
	if c, ok := codecOf(core.KindMeta, "NEXT-LOCAL-ID"); !ok || c != core.MetaNextLocalID {
		t.Fatal("NEXT-LOCAL-ID")
	}
	if c, ok := codecOf(core.KindValue, "UTF8-TEXT"); !ok || c != core.CodecUTF8TextEven {
		t.Fatal("UTF8-TEXT")
	}
	if c, ok := codecOf(core.KindReference, "DOC-POINTER"); !ok || c != core.CodecDocPointer {
		t.Fatal("DOC-POINTER")
	}
	if _, ok := codecOf(core.KindEdge, ""); ok {
		t.Fatal("empty: want false")
	}
	if _, ok := codecOf(core.KindEdge, "NOPE"); ok {
		t.Fatal("unknown: want false")
	}
}
