//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestFetchAddr — every address gives the same bytes from files, pack and zstd pack (latest.geul the newest live version, a padded Issuer the same Issuer); missing addresses are ErrNotFound in every form
package resolve

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
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
	addrs := []string{"issuers/latest.geul", "issuers/versions.geul", "0x00000000/docs/versions.geul", "0x00000000/docs/0xffff000000000007/versions.geul", "0x00000000/docs/0xffff000000000007/0x00000001.geul",
		"0x00000002/docs/0x0000000000000005/latest.geul", "0x00000000/docs/0xffff000000000008/versions.geul"}
	for _, a := range addrs {
		b, err := tree.Find(a)
		if err != nil {
			t.Fatal(a, err)
		}
		got[a] = b
	}
	missing := []string{"0x00000000/docs/0xffff000000000008/latest.geul", "0x00000000/docs/0xffff000000000008/0x00000001.geul", "0x00000000/docs/0xfffffffffffffffe/latest.geul", "0x00000000/docs/0xffffffffffffffff/versions.geul",
		"0x00000000/codecs/latest.geul", "0x00000009/docs/latest.geul", "0x00000000/docs/0xffff000000000009/versions.geul"}
	for _, f := range formsOf(t, tree) {
		checkAddrs(t, f, got, missing)
	}
	d := Dir{Root: t.TempDir()}
	d.Publish(edgeDoc(7, "a"))
	d.Deprecate([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7}, 1)
	if _, err := FetchAddr(ctx, d, Addr{Kind: AddrDoc, Issuer: []core.Word{0, 0}, ID: []core.Word{0xffff, 0, 0, 7}}); !errors.Is(err, ErrNotFound) {
		t.Fatal("only a deprecated version: want ErrNotFound,", err)
	}
	d.write(listPath(sampleKeyOf(7)), []byte{1})
	if _, err := FetchAddr(ctx, d, Addr{Kind: AddrDoc, Issuer: []core.Word{0, 0}, ID: []core.Word{0xffff, 0, 0, 7}}); err == nil {
		t.Fatal("broken version list: want error")
	}
}
