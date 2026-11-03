//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestFetchAddr — every address gives the same bytes from files, pack and zstd pack (latest.geul the newest live version, a padded Issuer the same Issuer); missing addresses are ErrNotFound in every form
package resolve

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
)

func TestFetchAddr(t *testing.T) {
	ctx := context.Background()
	tree := formsTree(t)
	v2, _ := tree.Find("0x00000000/docs/0xffff000000000007/0x00000002.geul")
	names, _ := tree.Fetch(ctx, namesPath(geulIssuer))
	got := map[string][]byte{
		"0x00000000/docs/0xffff000000000007/latest.geul":                            v2,
		"0x0000000000000000/docs/0xffff000000000007/latest.geul":                    v2,
		"0x" + strings.Repeat("0", 64) + "/docs/0xffff000000000007/0x00000002.geul": v2,
		"0x00000000/docs/latest.geul":                                               names,
		"0x00000000/docs/0x00000001.geul":                                           names,
	}
	addrs := []string{"issuers/latest.geul", "0x00000000/docs/0xffff000000000007/versions.geul", "0x00000000/docs/0xffff000000000007/0x00000001.geul",
		"0x00000002/docs/0x0000000000000005/latest.geul", "0x00000000/docs/0xffff000000000008/versions.geul"}
	for _, a := range addrs {
		b, err := tree.Find(a)
		if err != nil {
			t.Fatal(a, err)
		}
		got[a] = b
	}
	missing := []string{"0x00000000/docs/0xffff000000000008/latest.geul", "0x00000000/docs/0xffff000000000008/0x00000001.geul", "issuers/versions.geul",
		"0x00000000/codecs/latest.geul", "0x00000009/docs/latest.geul", "0x00000000/docs/0xffff000000000009/versions.geul"}
	for _, f := range formsOf(t, tree) {
		for a, want := range got {
			addr, err := ParseAddr(a)
			if err != nil {
				t.Fatal(a, err)
			}
			if b, err := FetchAddr(ctx, f.dir, addr); err != nil || !bytes.Equal(b, want) {
				t.Fatalf("%s: %s: %d bytes %v", f.name, a, len(b), err)
			}
		}
		for _, a := range missing {
			addr, _ := ParseAddr(a)
			if _, err := FetchAddr(ctx, f.dir, addr); !errors.Is(err, ErrNotFound) {
				t.Fatalf("%s: %s: %v", f.name, a, err)
			}
		}
	}
}
