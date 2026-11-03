//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeDeltaRemove — round trip in order; another Meta, nothing, a broken record, a non-LOCAL record and a repeat fail
package core

import (
	"slices"
	"testing"
)

func TestDecodeDeltaRemove(t *testing.T) {
	if ids, err := DecodeDeltaRemove(EncodeDeltaRemove([]uint64{12, 7})); err != nil || !slices.Equal(ids, []uint64{12, 7}) {
		t.Fatalf("%v %v", ids, err)
	}
	for _, r := range []Record{
		EncodeSpecVersion(2),
		EncodeDeltaRemove(nil),
		{Kind: KindMeta, Codec: MetaDeltaRemove, Payload: []Word{0x0051}},
		{Kind: KindMeta, Codec: MetaDeltaRemove, Payload: []Word{0x0AC0, 0x6162}},
		EncodeDeltaRemove([]uint64{7, 7}),
	} {
		if _, err := DecodeDeltaRemove(r); err == nil {
			t.Fatalf("%+v: want error", r)
		}
	}
}
