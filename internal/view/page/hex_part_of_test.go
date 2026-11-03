//ff:func feature=page type=builder control=sequence
//ff:what TestHexPartOf — plain words name their first line; a TAG part and an embedded record decode their TAG WORD; embedded UTF8-TEXT notes its string beside the first payload word; an ENUM name notes the first word
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestHexPartOf(t *testing.T) {
	h := hexPartOf(inspect.Part{Name: "span.off", Words: []core.Word{0x0684}})
	if h.Record || len(h.Lines) != 1 || h.Lines[0].Word != "0684" || h.Lines[0].Name != "span.off" || h.Lines[0].Tag {
		t.Fatalf("plain: %+v", h)
	}
	if h := hexPartOf(inspect.Part{Name: "TAG", Words: []core.Word{0x1BD9}, Head: true}); !h.Lines[0].Tag || h.Lines[0].Note == "" {
		t.Fatalf("head TAG: %+v", h)
	}
	h = hexPartOf(inspect.Part{Name: "ID", Words: []core.Word{0x0051, 1}, Record: true})
	if !h.Record || len(h.Lines) != 2 || !h.Lines[0].Tag || h.Lines[0].Name != "TAG" || h.Lines[1].Name != "대상 ID" || h.Lines[1].Word != "0001" {
		t.Fatalf("record: %+v", h)
	}
	if h.Lines[1].Note != "" {
		t.Fatal("LOCAL is not text")
	}
	h = hexPartOf(inspect.Part{Name: "UTF8-TEXT", Words: []core.Word{0x0522, 0x6d61, 0x696e}, Record: true})
	if h.Lines[1].Name != "payload" || h.Lines[1].Note != "main" || !h.Lines[1].Text || h.Lines[2].Note != "" {
		t.Fatalf("text: %+v", h)
	}
	if h = hexPartOf(inspect.Part{Name: "UTF8-TEXT", Words: []core.Word{0x0520}, Record: true}); len(h.Lines) != 1 {
		t.Fatalf("empty text: %+v", h)
	}
	if h = hexPartOf(inspect.Part{Name: "reason", Words: []core.Word{1}, Enum: "license"}); h.Lines[0].Note != "license" {
		t.Fatalf("enum: %+v", h)
	}
}
