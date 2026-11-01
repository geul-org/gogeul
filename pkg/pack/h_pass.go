//ff:type feature=pack type=model
//ff:what hPass — the H pass of Close in memory: the runs being merged, the head file and its writer, the pending leaf, keys, heads written, leaves since the last save
package pack

import (
	"bufio"
	"os"
)

// hPass is the state of the H pass (Stream.hPass): rs the runs (at most
// MergeRuns), heads the head file written through w, leaf the H leaf being
// filled, n the distinct keys so far, count the heads written and since the
// leaves written since the last save. Its memory is bounded by MergeRuns,
// HLeaf and the buffers, not by the number of keys.
type hPass struct {
	rs    []*runReader
	heads *os.File
	w     *bufio.Writer
	leaf  []streamKV
	n     uint64
	count int64
	since int
}
