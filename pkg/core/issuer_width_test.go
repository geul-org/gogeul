//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIssuerWidth — the last three codes of the Reference, Value and Edge zones are 2/4/16 words; the code before them, other ..D~..F codes, Meta and codes past the zone are not
package core

import "testing"

func TestIssuerWidth(t *testing.T) {
	cases := []struct {
		k    Kind
		c    uint16
		want int
	}{
		{KindReference, 0x5C, 0}, {KindReference, 0x5D, 2}, {KindReference, 0x5E, 4}, {KindReference, 0x5F, 16},
		{KindReference, 0x1D, 0}, {KindReference, 0x3D, 0}, {KindReference, 0x3E, 0}, {KindReference, 0x3F, 0},
		{KindReference, 0x42, 0}, {KindReference, 0x60, 0}, {KindReference, 0x7F, 0},
		{KindValue, 0x3C, 0}, {KindValue, 0x3D, 2}, {KindValue, 0x3E, 4}, {KindValue, 0x3F, 16}, {KindValue, 0x1D, 0}, {KindValue, 0x40, 0},
		{KindEdge, 0x11C, 0}, {KindEdge, 0x11D, 2}, {KindEdge, 0x11E, 4}, {KindEdge, 0x11F, 16},
		{KindEdge, 0x3D, 0}, {KindEdge, 0x7C, 0}, {KindEdge, 0x7D, 0}, {KindEdge, 0x7F, 0}, {KindEdge, 0xBD, 0}, {KindEdge, 0xFF, 0},
		{KindEdge, 0x120, 0}, {KindEdge, 0x13D, 0}, {KindEdge, 0x13F, 0},
		{KindMeta, 0x3D, 0}, {KindMeta, 0x3F, 0},
	}
	for _, c := range cases {
		if got := IssuerWidth(c.k, c.c); got != c.want {
			t.Fatalf("IssuerWidth(%v,0x%03X)=%d want %d", c.k, c.c, got, c.want)
		}
	}
}
