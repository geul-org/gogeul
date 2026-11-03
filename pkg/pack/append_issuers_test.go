//ff:func feature=pack type=encoder control=sequence
//ff:what TestAppendIssuers — width in words then the ID bytes; parseIssuers reads them back
package pack

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAppendIssuers(t *testing.T) {
	ids := []string{"0x00000000", "0x00000001"}
	b := appendIssuers(nil, ids)
	if !bytes.Equal(b, []byte{2, 0, 0, 0, 2, 0, 0, 0, 0, 2, 0, 0, 0, 1}) {
		t.Fatalf("%x", b)
	}
	got, err := parseIssuers(&cursor{b: b})
	if err != nil || !reflect.DeepEqual(got, ids) {
		t.Fatal(got, err)
	}
}
