//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestCodecKey — every Kind's codec spec document is the GEUL Issuer document of its widened target word
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCodecKey(t *testing.T) {
	for want, w := range map[string]core.Word{
		"0x00000000/0x000001e0": core.TargetWord(core.KindReference, core.CodecExternalPointer1Word),
		"0x00000000/0x00000050": core.TargetWord(core.KindReference, core.CodecLocal),
		"0x00000000/0x00000980": core.TargetWord(core.KindEdge, core.EdgePublication),
		"0x00000000/0x000004c0": core.TargetWord(core.KindValue, core.CodecPosixSeconds),
		"0x00000000/0x00001d10": core.TargetWord(core.KindMeta, core.MetaParent),
		"0x00000000/0x000003d0": core.TargetWord(core.KindReference, core.CodecRefIssuer2Word),
	} {
		if k := codecKey(w); k.Issuer+"/"+k.Doc != want {
			t.Fatalf("%04X: %+v", w, k)
		}
	}
}
