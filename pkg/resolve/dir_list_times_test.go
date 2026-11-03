//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirListTimes — a chain of lists newest first; the limit, odd bytes, a broken document and a missing parent end it
package resolve

import (
	"context"
	"testing"
)

func TestDirListTimes(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.Publish(edgeDoc(7, "a"))
	d.Publish(edgeDoc(7, "b"))
	k := sampleKeyOf(7)
	b, _ := d.Fetch(context.Background(), listPath(k))
	if ts := d.listTimes(geulIssuer, b, 5); len(ts) != 2 || ts[0].n != 2 || ts[1].n != 1 {
		t.Fatalf("%+v", ts)
	}
	if ts := d.listTimes(geulIssuer, b, 1); len(ts) != 1 {
		t.Fatalf("limit: %+v", ts)
	}
	for _, x := range [][]byte{{1}, {0, 5}} {
		if ts := d.listTimes(geulIssuer, x, 5); len(ts) != 0 {
			t.Fatalf("bad bytes: %+v", ts)
		}
	}
}
