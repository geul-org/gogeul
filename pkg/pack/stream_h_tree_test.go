//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamHTree — heads that fit the root are grouped by tree; more are grouped from the file a fanout at a time first, giving the same root as tree over them all; a short head file and a failing block write are errors
package pack

import (
	"bytes"
	"testing"
)

func TestStreamHTree(t *testing.T) {
	var file []byte
	var heads []streamKV
	for i := 0; i < 10; i++ {
		var e hEntry
		e.key[0], e.loc.Off = byte(i), uint64(i)
		file = append(file, appendHEntry(nil, e)...)
		heads = append(heads, streamKV{Key: string(e.key[:]), Loc: e.loc})
	}
	a, b := streamTestNew(t, nil), streamTestNew(t, nil)
	got, err := a.hTree(bytes.NewReader(file), 10)
	want, _ := b.tree(heads, blkHInner, 3, 3)
	if err != nil || !bytes.Equal(encodeBlock(got), encodeBlock(want)) || a.st.Index.Count != b.st.Index.Count || a.st.Index.Count == 0 {
		t.Fatalf("grouped: %v", err)
	}
	small, _ := a.hTree(bytes.NewReader(file), 3)
	if len(small.keys) != 3 {
		t.Fatal("fits the root")
	}
	if _, err := a.hTree(bytes.NewReader(file[:hEntrySize]), 3); err == nil {
		t.Fatal("short file, root: want error")
	}
	if _, err := a.hTree(bytes.NewReader(file[:hEntrySize]), 10); err == nil {
		t.Fatal("short file, grouped: want error")
	}
	if _, err := streamTestNew(t, failCodec{}).hTree(bytes.NewReader(file), 10); err == nil {
		t.Fatal("failing block: want error")
	}
}
