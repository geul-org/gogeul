//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestPutValue — appends an encoded Value; an invalid record fails
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestPutValue(t *testing.T) {
	if got, err := PutValue(nil, core.EncodeUintValue(7)); err != nil || !reflect.DeepEqual(got, []core.Word{0x0621, 7}) {
		t.Fatalf("got %04X %v", got, err)
	}
	if _, err := PutValue(nil, core.Record{Kind: core.KindValue, Codec: 0x80}); err == nil {
		t.Fatal("codec past 7 bits: want error")
	}
}
