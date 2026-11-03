//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestReadLength — LLLL in the low 4 bits: direct, 1-word and 2-word extensions, canonical flags and truncation
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestReadLength(t *testing.T) {
	type res struct {
		total, head   int
		canonical, ok bool
	}
	cases := []struct {
		ws   []core.Word
		want res
	}{
		{[]core.Word{0x0623}, res{4, 1, true, true}},
		{[]core.Word{0x060E, 20}, res{20, 2, true, true}},
		{[]core.Word{0x060E, 3}, res{3, 2, false, true}},
		{[]core.Word{0x060E, 1}, res{1, 2, false, false}},
		{[]core.Word{0x060F, 1, 0}, res{0x10000, 3, true, true}},
		{[]core.Word{0x060F, 0, 20}, res{20, 3, false, true}},
		{[]core.Word{0x060E}, res{}},
		{[]core.Word{0x060F, 1}, res{}},
	}
	for i, c := range cases {
		total, head, canonical, ok := readLength(c.ws, 0)
		if (res{total, head, canonical, ok}) != c.want {
			t.Fatalf("case %d: %v %v %v %v", i, total, head, canonical, ok)
		}
	}
}
