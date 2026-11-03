//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIssuerCodec — the ISSUER Codec of each Kind and width is ..D/..E/..F; Meta and other widths are errors
package core

import "testing"

func TestIssuerCodec(t *testing.T) {
	cases := []struct {
		k     Kind
		width int
		want  uint16
	}{
		{KindReference, 2, 0x3D}, {KindReference, 4, 0x3E}, {KindReference, 16, 0x3F},
		{KindValue, 2, 0x3D}, {KindValue, 4, 0x3E}, {KindValue, 16, 0x3F},
		{KindEdge, 2, 0x13D}, {KindEdge, 4, 0x13E}, {KindEdge, 16, 0x13F},
	}
	for _, c := range cases {
		got, err := IssuerCodec(c.k, c.width)
		if err != nil || got != c.want || IssuerWidth(c.k, got) != c.width {
			t.Fatalf("%v %d: 0x%03X %v", c.k, c.width, got, err)
		}
	}
	for _, bad := range []struct {
		k     Kind
		width int
	}{{KindValue, 1}, {KindEdge, 8}, {KindReference, 0}, {KindMeta, 2}} {
		if _, err := IssuerCodec(bad.k, bad.width); err == nil {
			t.Fatalf("%v width %d accepted", bad.k, bad.width)
		}
	}
}
