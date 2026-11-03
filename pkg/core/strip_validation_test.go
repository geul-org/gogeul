//ff:func feature=core type=helper control=sequence
//ff:what TestStripValidation — validation Meta is dropped, other Meta and BODY records are kept
package core

import (
	"reflect"
	"testing"
)

func TestStripValidation(t *testing.T) {
	name, _ := EncodeRecord(Record{Kind: KindMeta, Codec: MetaDocID, Payload: []Word{0, 0, 0, 0, 1}})
	verdict, _ := EncodeRecord(Record{Kind: KindMeta, Codec: MetaVerdict, Payload: []Word{1}})
	in := append(append(append([]Word{}, name...), verdict...), 0x1C30, 0x1C40)
	want := append(append([]Word{}, name...), 0x1C30, 0x1C40)
	if got := stripValidation(in); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %04X", got)
	}
}
