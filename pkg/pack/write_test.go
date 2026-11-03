//ff:func feature=pack type=writer control=sequence
//ff:what TestWrite — a batch becomes generation 1, then a second batch adds generation 2 on top
package pack

import (
	"context"
	"path/filepath"
	"testing"
)

func TestWrite(t *testing.T) {
	ctx := context.Background()
	out := filepath.Join(t.TempDir(), "p")
	doc := fixtureFile(t, "issuers/0x00000001.geul")
	res, err := Write(ctx, out, Batch{Put: []string{"issuers/0x00000001.geul"}, Get: memGet(map[string][]byte{"issuers/0x00000001.geul": doc})}, BuildOptions{Verify: true})
	if err != nil || res.Generation != 1 || !res.Changed || res.Records != 1 {
		t.Fatal(res, err)
	}
	res, err = Write(ctx, out, Batch{Put: []string{"issuers/0x00000002.geul"}, Get: memGet(map[string][]byte{"issuers/0x00000002.geul": {1}})}, BuildOptions{Verify: true})
	if err != nil || res.Generation != 2 {
		t.Fatal(res, err)
	}
	if b, err := opened(t, out).Fetch(ctx, "issuers/0x00000001.geul"); err != nil || string(b) != string(doc) {
		t.Fatal(err)
	}
}
