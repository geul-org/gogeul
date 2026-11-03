//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeText — a Value error is returned; TEXT golden bytes with a LOCAL Edge Id
package toulmin

import (
	"errors"
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestEncodeText pins the TEXT record of "abc" (EDGE-ISSUER-2WORD, GEUL Issuer 0x00000000,
// TEXT direct code 0x121L (§3.14), LOCAL Edge Id 2, UTF8-TEXT-ODD Value).
func TestEncodeText(t *testing.T) {
	got, err := EncodeText(TextPacket{TID: 2, Bytes: []byte("abc")}, testWire)
	if want := []core.Word{0x1215, 0x0051, 2, 0x0732, 0x6162, 0x6300}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("got %04X %v want %04X", got, err, want)
	}
	defer func(orig func([]core.Word, core.Record) ([]core.Word, error)) { putValue = orig }(putValue)
	putValue = func([]core.Word, core.Record) ([]core.Word, error) { return nil, errors.New("too long") }
	if _, err := EncodeText(TextPacket{TID: 2, Bytes: []byte("abc")}, testWire); err == nil {
		t.Fatal("value error: want error")
	}
}
