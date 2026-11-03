//ff:func feature=pack type=writer control=sequence
//ff:what TestUnpackOne — writes one file with parents; fetch and directory failures are returned
package pack

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestUnpackOne(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	base := t.TempDir()
	if err := unpackOne(ctx, r, base, "0x00000000/codecs/0x000003c0/0x00000001.geul"); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(base, "0x00000000", "codecs", "0x000003c0", "0x00000001.geul")); err != nil {
		t.Fatal(err)
	}
	if err := unpackOne(ctx, r, base, "none"); err == nil {
		t.Fatal("fetch")
	}
	os.WriteFile(filepath.Join(base, "f"), nil, 0o644)
	if err := unpackOne(ctx, r, filepath.Join(base, "f"), "issuers/0x00000001.geul"); err == nil {
		t.Fatal("mkdir")
	}
}
