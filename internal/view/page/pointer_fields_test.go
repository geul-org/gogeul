//ff:func feature=page type=helper control=sequence
//ff:what TestPointerFields — DOC-POINTER and CODEC-POINTER payloads split into Issuer, ID and version; a short payload keeps what is there
package page

import (
	"reflect"
	"testing"
)

func TestPointerFields(t *testing.T) {
	if got := pointerFields(8, "문서 ID", 4); !reflect.DeepEqual(got, []refField{{"Issuer", 2}, {"문서 ID", 4}, {"버전 ID", 2}}) {
		t.Fatal(got)
	}
	if got := pointerFields(20, "코덱 ID", 2); !reflect.DeepEqual(got, []refField{{"Issuer", 16}, {"코덱 ID", 2}, {"버전 ID", 2}}) {
		t.Fatal(got)
	}
	if got := pointerFields(3, "문서 ID", 4); !reflect.DeepEqual(got, []refField{{"Issuer", 0}, {"문서 ID", 3}, {"버전 ID", 0}}) {
		t.Fatal(got)
	}
}
