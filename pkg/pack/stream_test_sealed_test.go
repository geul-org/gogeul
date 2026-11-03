//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what streamTestSealed — put test records one by one until the first data pack is cut; return the stream, the sealed pack's path and the open-pack checkpoints (open object, inventory, paths) taken before each record
package pack

import (
	"path/filepath"
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// streamTestSealed fills the first data pack of a new test stream (a named
// record first, so the pack has a path) and returns, once the next record
// has cut it, the sealed object and the snapshots taken while it was open.
func streamTestSealed(t *testing.T) (*Stream, string, []streamSnap) {
	t.Helper()
	s := streamTestNew(t, nil)
	r, _ := s.Encode([]byte("q1 0x5100000000000001\n"))
	s.PutNamed("0x00000001/codecs/0x00000420/0x00000001.geul", r)
	var snaps []streamSnap
	for i := uint64(2); s.st.Data.Num == 0; i++ {
		snaps = append(snaps, streamSnap{open: s.st.Data, inv: slices.Clone(s.st.Inv), paths: slices.Clone(s.st.Paths)})
		r, _ := s.Encode(streamTestDoc(i/2, uint32(i%2)))
		if err := s.Put([]core.Word{0, 1}, 0x5100000000000000|i/2, uint32(i%2), r); err != nil {
			t.Fatal(err)
		}
	}
	o := s.st.Objects[0]
	return s, filepath.Join(s.out, filepath.FromSlash(objectName(object{kind: o.Kind, hash: o.Hash}))), snaps
}
