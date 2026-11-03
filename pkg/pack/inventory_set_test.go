//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what inventorySet — a set of several data packs written with c: shared and named records, then a generation that replaces one, adds an h/-only record and deletes a P, a named and an h/ path
package pack

import (
	"context"
	"path/filepath"
	"testing"
)

// inventorySet writes documents 1–6 (slot 1, each with its h/ copy) and
// n/a.geul, n/b.geul, then replaces document 2, adds document 7 as h/-only and
// deletes document 3, n/b.geul and document 4's h/ copy. It returns the set
// folder and the Get holding the final bytes.
func inventorySet(t *testing.T, c Codec) (string, map[string][]byte) {
	t.Helper()
	ctx := context.Background()
	docs := map[string][]byte{"issuers/0x0000000a.geul": streamTestDoc(20, 1), "issuers/0x0000000b.geul": streamTestDoc(21, 1)}
	put := []string{"issuers/0x0000000a.geul", "issuers/0x0000000b.geul"}
	for d := uint64(1); d <= 6; d++ {
		doc := streamTestDoc(d, 1)
		docs[pPath("0x00000001", d, 1)], docs[docHPath(doc)] = doc, doc
		put = append(put, pPath("0x00000001", d, 1), docHPath(doc))
	}
	out := filepath.Join(t.TempDir(), "set")
	o := small()
	o.Codec, o.PackSize = c, 500
	if _, err := Write(ctx, out, Batch{Put: put, Get: memGet(docs)}, o); err != nil {
		t.Fatal(err)
	}
	only := streamTestDoc(7, 1)
	docs[pPath("0x00000001", 2, 1)], docs[docHPath(only)] = streamTestDoc(2, 2), only
	put = []string{pPath("0x00000001", 2, 1), docHPath(only)}
	del := []string{pPath("0x00000001", 3, 1), "issuers/0x0000000b.geul", docHPath(streamTestDoc(4, 1))}
	if _, err := Write(ctx, out, Batch{Put: put, Delete: del, Get: memGet(docs)}, o); err != nil {
		t.Fatal(err)
	}
	return out, docs
}
