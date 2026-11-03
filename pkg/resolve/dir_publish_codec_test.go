//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirPublishCodec — a spec document lands at its codec address; odd bytes, bad hash input and documents without one DECLARE fail
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"os"
	"testing"
)

func TestDirPublishCodec(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	if e, err := d.PublishCodec(codecDocOf(0x0980, 4, 2)); err != nil || e.Version != 1 {
		t.Fatalf("publish: %+v %v", e, err)
	}
	if _, err := os.Stat(d.file("0x00000000/codecs/0x00000980/0x00000001.geul")); err != nil {
		t.Fatal(err)
	}
	if _, err := d.PublishCodec([]byte{1}); err == nil {
		t.Fatal("odd bytes: want error")
	}
	if _, err := d.PublishCodec(core.PutWords(core.WrapDocument([]core.Word{0x1C30}))); err == nil {
		t.Fatal("no END: want error")
	}
	if _, err := d.PublishCodec(core.PutWords(core.WrapDocument([]core.Word{0x1C30, 0x1C40}))); err == nil {
		t.Fatal("no DECLARE: want error")
	}
	decl, _ := core.EncodeRecord(core.EncodeDeclare(core.Declare{EdgeID: 1, Target: 0x0990}))
	if _, err := d.PublishCodec(core.PutWords(core.WrapDocument(append(append([]core.Word{0x1C30}, decl...), 0x1C40)))); err == nil {
		t.Fatal("no DOC-ID: want error")
	}
	other, _ := core.EncodeDocIDCodec([]core.Word{0, 0}, 0x09A0)
	ow, _ := core.EncodeRecord(other)
	if _, err := d.PublishCodec(core.PutWords(core.WrapDocument(append(append(append(ow, 0x1C30), decl...), 0x1C40)))); err == nil {
		t.Fatal("DOC-ID of another codec: want error")
	}
}
