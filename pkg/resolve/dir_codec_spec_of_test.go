//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirCodecSpecOf — the specs of one codec document; an unpublished codec and a document of another Kind fail
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirCodecSpecOf(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.PublishCodec(codecDocOf(0x0B80, 4, 2))
	s, err := d.codecSpecOf(0x0B80)
	if err != nil || len(s[0x0B80]) != 2 {
		t.Fatal(s, err)
	}
	if _, err := d.codecSpecOf(0x0BA0); err == nil {
		t.Fatal("unpublished")
	}
	value := codecDocAs(0x0BA0, 0x07A0, 4)
	ws, _ := core.ReadWords(value)
	h, _ := core.ContentHash(ws)
	if _, err := d.publishAt(codecKey(0x0BA0), []core.Word{0, 0}, codecDocID(0x0BA0), value, h); err != nil {
		t.Fatal(err)
	}
	if _, err := d.codecSpecOf(0x0BA0); err == nil {
		t.Fatal("a Value codec document at an Edge codec ID: want error (§1.1 103)")
	}
}
