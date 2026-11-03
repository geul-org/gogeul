//ff:func feature=resolve type=rule control=sequence
//ff:what TestDirSameCodec — the same target passes, another target at the codec ID is refused (§1.1 103); deleted versions, documents, DECLARE-less input pass; a missing version file fails
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirSameCodec(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	k, id := codecKey(0x0980), codecDocID(0x0980)
	d.write(versionPath(k, 1), codecDocOf(0x0980, 4))
	d.write(versionPath(k, 2), codecDocOf(0x05A0, 4))
	live := versionLog{Entries: []VersionEntry{{Version: 1}}}
	if err := d.sameCodec(k, id, live, codecDocOf(0x0980, 4, 2)); err != nil {
		t.Fatalf("same codec: %v", err)
	}
	if err := d.sameCodec(k, id, live, codecDocOf(0x05A0, 4)); err == nil {
		t.Fatal("another Kind at the same codec ID: want refusal")
	}
	if err := d.sameCodec(k, id, live, codecDocOf(0x0990, 4)); err == nil {
		t.Fatal("another codec at the same codec ID: want refusal")
	}
	deleted := versionLog{Entries: []VersionEntry{{Version: 2, Deleted: true}, {Version: 1}}}
	if err := d.sameCodec(k, id, deleted, codecDocOf(0x0980, 4)); err != nil {
		t.Fatalf("deleted version 2 is not compared: %v", err)
	}
	if err := d.sameCodec(k, []core.Word{0xffff, 0, 0, 7}, live, codecDocOf(0x05A0, 4)); err != nil {
		t.Fatalf("a document ID is not a codec: %v", err)
	}
	if err := d.sameCodec(k, id, live, edgeDoc(7, "a")); err != nil {
		t.Fatalf("no DECLARE: %v", err)
	}
	missing := versionLog{Entries: []VersionEntry{{Version: 3}}}
	if err := d.sameCodec(k, id, missing, codecDocOf(0x0980, 4)); err == nil {
		t.Fatal("missing version file: want error")
	}
}
