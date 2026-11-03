//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeSpecVersion — v1 encodes as 0x1890 0x0001
package core

import (
	"reflect"
	"testing"
)

func TestEncodeSpecVersion(t *testing.T) {
	if ws, err := EncodeRecord(EncodeSpecVersion(SpecVersion)); err != nil || !reflect.DeepEqual(ws, []Word{0x1C51, 0x0001}) {
		t.Fatalf("got %04X %v", ws, err)
	}
}
