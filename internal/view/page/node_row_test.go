//ff:func feature=page type=builder control=sequence
//ff:what TestNodeRow — anchor, codec name, payload label and words; a lone TAG has no label
package page

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"testing"
)

func TestNodeRow(t *testing.T) {
	r := nodeRow(2, inspect.Entry{Record: core.EncodeLocalRef(7), Words: []core.Word{0x0051, 7}, Name: "LOCAL"})
	if r.Anchor != "n2" || r.Name != "LOCAL" || r.Label != "0x0007" || r.Summary != "#7" || len(r.HeadParts) != 1 || len(r.BodyParts) != 1 {
		t.Fatalf("got %+v", r)
	}
	c, _, _ := core.DecodeRecord([]core.Word{0x1018})
	if r := nodeRow(0, inspect.Entry{Record: c, Words: []core.Word{0x1018}, Name: "PUBLICATION"}); r.Label != "" {
		t.Fatalf("codec: %+v", r)
	}
}
