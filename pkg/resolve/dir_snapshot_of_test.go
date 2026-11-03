//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirSnapshotOf — a stored diff gives the snapshot back; a missing file, a too long chain, a base not in the list, a broken base, a hash mismatch and a diff that does not apply fail
package resolve

import (
	"bytes"
	"context"
	"os"
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/delta"
)

func TestDirSnapshotOf(t *testing.T) {
	ctx := context.Background()
	d := Dir{Root: t.TempDir()}
	issuer, doc := []core.Word{0, 0}, []core.Word{0xffff, 0, 0, 8}
	k := Key{Issuer: issuer, Doc: doc}.DocKey()
	v1, v2 := edgeDoc(8, "a"), edgeDoc(8, "a", "b")
	e1, _ := d.Publish(v1)
	e2, _ := d.Publish(v2)
	es := []VersionEntry{e1, e2}
	if _, err := d.snapshotOf(ctx, k, es, VersionEntry{Version: 9}, 0); err == nil {
		t.Fatal("missing file: want error")
	}
	diff, _ := delta.Diff(v1, v2, 1)
	os.WriteFile(d.file(versionPath(k, 2)), diff, 0o644)
	if b, err := d.snapshotOf(ctx, k, es, e2, 0); err != nil || !bytes.Equal(b, v2) {
		t.Fatalf("diff applied: %v", err)
	}
	if _, err := d.snapshotOf(ctx, k, es, e2, maxDeltaChain); err == nil {
		t.Fatal("chain: want error")
	}
	if _, err := d.snapshotOf(ctx, k, []VersionEntry{e2}, e2, 0); err == nil {
		t.Fatal("base not listed: want error")
	}
	if _, err := d.snapshotOf(ctx, k, []VersionEntry{{Version: 1, Hash: [32]byte{9}}, e2}, e2, 0); err == nil {
		t.Fatal("base hash: want error")
	}
	os.WriteFile(d.file(versionPath(k, 2)), edgeDoc(8, "x"), 0o644)
	if _, err := d.snapshotOf(ctx, k, es, e2, 0); err == nil {
		t.Fatal("hash mismatch: want error")
	}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	id, _ := core.EncodeDocID(issuer, doc)
	bad := slices.Concat(enc(id), enc(core.EncodeDeltaBase(1)), enc(core.EncodeDeltaRemove([]uint64{9})), []core.Word{0x1C30, 0x1C40})
	os.WriteFile(d.file(versionPath(k, 2)), core.PutWords(core.WrapDocument(bad)), 0o644)
	if _, err := d.snapshotOf(ctx, k, es, e2, 0); err == nil {
		t.Fatal("diff removing a missing Edge: want error")
	}
}
