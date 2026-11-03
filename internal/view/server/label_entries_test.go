//ff:func feature=server type=helper control=sequence
//ff:what TestLabelEntries — a labelled Issuer directory shows its name and keeps its ID; files and unlabelled directories stay
package server

import (
	"reflect"
	"testing"
)

func TestLabelEntries(t *testing.T) {
	es := []indexEntry{{Name: "0x00000000/", URL: "/a/0x00000000/"}, {Name: "0x0001/", URL: "/a/0x0001/"}, {Name: "0x00000000", URL: "/a/0x0000"}}
	labelEntries(es, map[string]string{"0x00000000": "geul.org"})
	want := []indexEntry{{Name: "geul.org/", URL: "/a/0x00000000/", ID: "0x00000000"}, {Name: "0x0001/", URL: "/a/0x0001/"}, {Name: "0x00000000", URL: "/a/0x0000"}}
	if !reflect.DeepEqual(es, want) {
		t.Fatalf("got %+v", es)
	}
}
