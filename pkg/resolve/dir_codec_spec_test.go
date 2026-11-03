//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirCodecSpec — specs of every published codec; stray folders, folders without a version list and other GEUL documents skipped; a broken document fails
package resolve

import (
	"testing"
)

func TestDirCodecSpec(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.PublishCodec(codecDocOf(0x0B80, 4, 2))
	d.PublishCodec(codecDocOf(0x0BA0, 4))
	d.write("0x00000000/junk/versions.geul", []byte("x"))
	d.write("0x00000000/docs/0xFFFF000000000001/versions.geul", []byte("x"))
	d.write(versionPath(codecKey(0x0BC0), 1), []byte("no list"))
	s, err := d.CodecSpec()
	if err != nil || len(s) != 2 || len(s[0x0B80]) != 2 {
		t.Fatalf("got %+v %v", s, err)
	}
	d.write(versionPath(codecKey(0x0BA0), 1), codecDocOf(0x0BA0, 9))
	if _, err := d.CodecSpec(); err == nil {
		t.Fatal("tampered: want error")
	}
}
