//ff:func feature=inspect type=decoder control=sequence
//ff:what TestMetaValue — one known Value decodes; trailing words, a non-Value or a bad record do not
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestMetaValue(t *testing.T) {
	vw, _ := core.EncodeRecord(core.EncodeUintValue(3))
	if v, ok := metaValue(core.Record{Payload: vw}); !ok || v.Uint != 3 {
		t.Fatalf("got %+v %v", v, ok)
	}
	if _, ok := metaValue(core.Record{Payload: append(vw, 0)}); ok {
		t.Fatal("trailing word: want not ok")
	}
	if _, ok := metaValue(core.Record{Payload: []core.Word{0x1C40}}); ok {
		t.Fatal("non-Value: want not ok")
	}
	if _, ok := metaValue(core.Record{}); ok {
		t.Fatal("empty: want not ok")
	}
}
