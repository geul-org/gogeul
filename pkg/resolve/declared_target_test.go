//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestDeclaredTarget — the one DECLARE's target; no DECLARE, two DECLAREs, odd bytes and a broken document give false
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDeclaredTarget(t *testing.T) {
	if w, ok := declaredTarget(codecDocOf(0x05A0, 4)); !ok || w != 0x05A0 {
		t.Fatalf("one: %04X %v", w, ok)
	}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	in := append(enc(core.EncodeSpecVersion(core.SpecVersion)), core.TargetWord(core.KindMeta, core.MetaBody))
	in = append(in, enc(core.EncodeDeclare(core.Declare{EdgeID: 1, Target: 0x0980}))...)
	in = append(in, enc(core.EncodeDeclare(core.Declare{EdgeID: 2, Target: 0x0990}))...)
	two := core.PutWords(core.WrapDocument(append(in, core.TargetWord(core.KindMeta, core.MetaEnd))))
	for i, b := range [][]byte{edgeDoc(7, "a"), two, {0x1c}, {0, 0}} {
		if w, ok := declaredTarget(b); ok {
			t.Fatalf("bad %d: %04X", i, w)
		}
	}
}
