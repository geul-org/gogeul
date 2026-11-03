//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirCodecSpecs — only the asked targets; a target without a spec is left out; a broken list or a tampered document fails
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirCodecSpecs(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.PublishCodec(codecDocOf(0x0B80, 4, 2))
	d.PublishCodec(codecDocOf(0x0BA0, 4))
	s, err := d.CodecSpecs([]core.Word{0x0B80, 0x121C})
	if err != nil || len(s) != 1 || len(s[0x0B80]) != 2 {
		t.Fatalf("got %+v %v", s, err)
	}
	if s, err := d.CodecSpecs(nil); err != nil || len(s) != 0 {
		t.Fatalf("no targets: %+v %v", s, err)
	}
	d.write(versionPath(codecKey(0x0BA0), 1), codecDocOf(0x0BA0, 9))
	if _, err := d.CodecSpecs([]core.Word{0x0B80}); err != nil {
		t.Fatal("another target's fault:", err)
	}
	if _, err := d.CodecSpecs([]core.Word{0x0BA0}); err == nil {
		t.Fatal("tampered: want error")
	}
}
