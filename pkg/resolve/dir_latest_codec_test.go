//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirLatestCodec — the latest version; missing, empty-list and tampered documents fail
package resolve

import (
	"bytes"
	"errors"
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestDirLatestCodec(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.PublishCodec(codecDocOf(0x0980, 4))
	d.PublishCodec(codecDocOf(0x0980, 4, 2))
	if b, err := d.latestCodec(0x0980); err != nil || !bytes.Equal(b, codecDocOf(0x0980, 4, 2)) {
		t.Fatalf("latest: %v", err)
	}
	if _, err := d.latestCodec(0x0990); !errors.Is(err, ErrNotFound) {
		t.Fatalf("missing: %v", err)
	}
	k := codecKey(0x09A0)
	d.write(listPath(k), listDocOf(0, VersionEntry{Version: 1, Deleted: true}))
	if _, err := d.latestCodec(0x09A0); err == nil {
		t.Fatal("bad list: want error")
	}
	d.write(listPath(k), encodeVersionList([]core.Record{core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Issuer: []core.Word{0, 0}, Doc: codecDocID(0x09A0), Version: 1})}, nil, 0, 0))
	if _, err := d.latestCodec(0x09A0); !errors.Is(err, ErrNotFound) {
		t.Fatalf("no file: %v", err)
	}
	d.write(versionPath(k, 1), codecDocOf(0x09A0))
	if _, err := d.latestCodec(0x09A0); !errors.Is(err, ErrVerify) {
		t.Fatalf("tampered: %v", err)
	}
	d.write(listPath(k), encodeVersionList([]core.Record{core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Issuer: []core.Word{0, 0}, Doc: codecDocID(0x09A0), Version: 1}), core.EncodePublication(core.Publication{Codec: core.EdgeTombstone, EdgeID: 2, Issuer: []core.Word{0, 0}, Doc: codecDocID(0x09A0), Version: 1})}, nil, 0, 0))
	if _, err := d.latestCodec(0x09A0); !errors.Is(err, ErrNotFound) {
		t.Fatalf("deleted: %v", err)
	}
}
