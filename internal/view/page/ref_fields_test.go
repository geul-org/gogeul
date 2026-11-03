//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestRefFields — Reference payload fields by C at the 5Q, pointer, CODEC-POINTER and REFERENCE-ISSUER boundaries
package page

import (
	"reflect"
	"testing"
)

func TestRefFields(t *testing.T) {
	cases := []struct {
		c    uint16
		n    int
		want []refField
	}{
		{0x1E, 2, []refField{{"IMPORT ID", 1}, {"대상 ID", 1}}},
		{0x24, 3, []refField{{"IMPORT ID", 2}, {"대상 ID", 1}}},
		{0x2C, 5, []refField{{"IMPORT ID", 4}, {"대상 ID", 1}}},
		{0x2D, 7, []refField{{"Issuer", 2}, {"문서 ID", 4}, {"대상 ID", 1}}},
		{0x36, 9, []refField{{"Issuer", 4}, {"문서 ID", 4}, {"대상 ID", 1}}},
		{0x3B, 21, []refField{{"Issuer", 16}, {"문서 ID", 4}, {"대상 ID", 1}}},
		{0x2E, 3, []refField{{"Issuer", 2}, {"문서 ID", 1}, {"대상 ID", 0}}},
		{0x14, 2, []refField{{"대상 ID", 2}}},
		{0x1D, 2, []refField{{"대상 ID", 2}}},
		{0x03, 2, []refField{{"Issuer", 2}}},
		{0x04, 8, []refField{{"Issuer", 2}, {"문서 ID", 4}, {"버전 ID", 2}}},
		{0x04, 22, []refField{{"Issuer", 16}, {"문서 ID", 4}, {"버전 ID", 2}}},
		{0x04, 5, []refField{{"Issuer", 0}, {"문서 ID", 4}, {"버전 ID", 1}}},
		{0x3C, 6, []refField{{"Issuer", 2}, {"코덱 ID", 2}, {"버전 ID", 2}}},
		{0x3C, 20, []refField{{"Issuer", 16}, {"코덱 ID", 2}, {"버전 ID", 2}}},
		{0x3D, 3, []refField{{"Payload", 3}}},
		{0x3F, 0, []refField{{"Payload", 0}}},
	}
	for i, c := range cases {
		if got := refFields(c.c, c.n); !reflect.DeepEqual(got, c.want) {
			t.Fatalf("case %d: %+v", i, got)
		}
	}
}
