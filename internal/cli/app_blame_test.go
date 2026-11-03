//ff:func feature=cli type=command control=sequence
//ff:what TestAppBlame — up to a version; JSON; a version with an Edge lacking its LOCAL Edge Id fails
package cli

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestAppBlame(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	if _, err := h.run("", "blame", "x"); err == nil {
		t.Fatal("bad address: want error")
	}
	if out := h.must("blame", addr+"@1"); strings.Contains(out, "ALIAS") {
		t.Fatal(out)
	}
	if out := h.must("blame", addr, "--format", "json"); !strings.Contains(out, `"changed": 2`) {
		t.Fatal(out)
	}
	id, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0, 0, 0, 9})
	noEdgeID := core.Record{Kind: core.KindEdge, Codec: 0x04, Payload: []core.Word{0x0421, 1}}
	odd, _ := core.EncodeDocument([]core.Record{id}, []core.Record{noEdgeID})
	(resolve.Dir{Root: h.data}).Publish(odd)
	if _, err := h.run("", "blame", "0x00000000/0x0000000000000009"); err == nil {
		t.Fatal("an Edge without a LOCAL Edge Id: want error")
	}
}
