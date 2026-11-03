//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestEncodeSource — a Value error is returned; SOURCE golden bytes plus the bad SPAN width error
package toulmin

import (
	"errors"
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestEncodeSource pins the SOURCE record of "ab" with a 2-word SPAN width (width
// word 0x4000) and checks an undeclarable SPAN width.
func TestEncodeSource(t *testing.T) {
	got, err := EncodeSource(TextPacket{TID: 1, Bytes: []byte("ab")}, Wire{Span: 2})
	if want := []core.Word{0x1205, 0x0051, 1, 0x4000, 0x0721, 0x6162}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("got %04X %v want %04X", got, err, want)
	}
	if _, err := EncodeSource(TextPacket{}, Wire{Span: 3}); err == nil {
		t.Fatal("span width 3: want error")
	}
	defer func(orig func([]core.Word, core.Record) ([]core.Word, error)) { putValue = orig }(putValue)
	putValue = func([]core.Word, core.Record) ([]core.Word, error) { return nil, errors.New("too long") }
	if _, err := EncodeSource(TextPacket{TID: 1, Bytes: []byte("ab")}, Wire{Span: 2}); err == nil {
		t.Fatal("value error: want error")
	}
}
