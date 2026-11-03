//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIsValidationMeta — Meta C ≥ 0x25 (ISSUER-KEY onward) is validation class; lower Meta and other Kinds are not
package core

import "testing"

func TestIsValidationMeta(t *testing.T) {
	cases := []struct {
		k    Kind
		c    uint16
		want bool
	}{{KindMeta, 0x24, false}, {KindMeta, 0x25, true}, {KindMeta, 0x3F, true}, {KindValue, 0x25, false}, {KindEdge, 0x30, false}}
	for _, c := range cases {
		if got := IsValidationMeta(c.k, c.c); got != c.want {
			t.Fatalf("%v %X: %v", c.k, c.c, got)
		}
	}
}
