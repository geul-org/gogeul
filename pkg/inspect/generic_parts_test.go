//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestGenericParts — an Edge's leading LOCAL id, embedded records after it, a Meta's Value, a raw Meta, a bare payload and an empty payload
package inspect

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestGenericParts(t *testing.T) {
	cases := []struct {
		e     Entry
		names []string
	}{
		{Entry{Record: core.Record{Kind: core.KindEdge, Payload: []core.Word{0x0051, 1, 9}}}, []string{"ID", "payload"}},
		{Entry{Record: core.Record{Kind: core.KindMeta, Payload: []core.Word{0x0421, 1}}, Data: Value{}}, []string{"value"}},
		{Entry{Record: core.Record{Kind: core.KindEdge, Payload: []core.Word{0x0421, 1}}}, []string{"UNSIGNED-INTEGER"}},
		{Entry{Record: core.Record{Kind: core.KindEdge, Payload: []core.Word{0x0051, 1, 0x0051, 7, 0x0522, 0x6d61, 0x696e}}}, []string{"ID", "LOCAL", "UTF8-TEXT"}},
		{Entry{Record: core.Record{Kind: core.KindMeta, Payload: []core.Word{0x0002}}, Data: Value{}}, []string{"payload"}},
		{Entry{Record: core.Record{Kind: core.KindMeta}}, nil},
	}
	for i, c := range cases {
		var names []string
		for _, p := range genericParts(c.e) {
			names = append(names, p.Name)
		}
		if !slices.Equal(names, c.names) {
			t.Fatalf("case %d: %v", i, names)
		}
	}
}
