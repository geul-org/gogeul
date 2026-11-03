//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestFmtRef — a codec number lives under EDGE-ISSUER-2WORD
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFmtRef(t *testing.T) {
	if r := fmtRef(FormatText); r.Codec != core.CodecEdgeIssuer2Word || r.Format != FormatText {
		t.Fatalf("got %+v", r)
	}
}
