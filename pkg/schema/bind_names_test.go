//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestBindNames — names on one Subject each, a repeated BIND and several names on one Subject pass; one name on two Subjects fails E12 once per rebinding; other records are ignored
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestBindNames(t *testing.T) {
	q := func(n core.Word) core.Record { return core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID, Payload: []core.Word{n}} }
	cases := []struct {
		recs []core.Record
		want string
	}{
		{[]core.Record{bindOf(1, q(1), "a"), bindOf(2, q(2), "b"), bindOf(3, q(1), "a"), bindOf(4, q(1), "c"), core.EncodeUintValue(1)}, ""},
		{[]core.Record{bindOf(1, q(1), "ab"), bindOf(2, q(2), "ab")}, "E12"},
		{[]core.Record{bindOf(1, q(1), "ab"), bindOf(2, q(2), "ab"), bindOf(3, q(3), "ab"), bindOf(4, q(1), "ab")}, "E12,E12"},
		{[]core.Record{bindOf(1, q(1), "ab"), bindOf(2, q(1), "AB")}, ""},
	}
	for i, c := range cases {
		var rs []rec
		for _, r := range c.recs {
			rs = append(rs, recOf(r))
		}
		if got := rulesOf(bindNames(rs)); got != c.want {
			t.Fatalf("case %d: %q want %q", i, got, c.want)
		}
	}
}
