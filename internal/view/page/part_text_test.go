//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestPartText — a raw field by ENUM name, as text when it holds whole records, or hex, a decoded Value, a LOCAL Reference as #n, a lone TAG by codec name
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestPartText(t *testing.T) {
	ref, _ := core.EncodeRecord(core.EncodeLocalRef(7))
	for want, p := range map[string]inspect.Part{
		"license":     {Words: []core.Word{1}, Enum: "license"},
		"2":           {Words: []core.Word{2}, Data: inspect.Value{Codec: core.CodecUnsignedInteger, Uint: 2}},
		"0x00000002":  {Words: []core.Word{0, 2}},
		`"v1"`:        {Words: []core.Word{0x0721, 0x7631}},
		"9":           {Record: true, Data: inspect.Value{Codec: core.CodecUnsignedInteger, Uint: 9}},
		"#7":          {Record: true, Words: ref},
		"PUBLICATION": {Record: true, Words: []core.Word{0x0B80}},
	} {
		if got := partText(p); got != want {
			t.Fatalf("got %q want %q", got, want)
		}
	}
}
