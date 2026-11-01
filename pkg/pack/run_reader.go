//ff:type feature=pack type=model
//ff:what runReader — one H run file read entry by entry during the merge, counting the entries it gave
package pack

import (
	"bufio"
	"os"
)

// runReader reads one sorted run. cur is the entry at the front; ok is false
// once the run is used up. taken counts the entries the merge took from it
// (nextMerged), where a resumed merge reopens it (openRuns).
type runReader struct {
	f     *os.File
	r     *bufio.Reader
	cur   hEntry
	ok    bool
	taken int64
}
