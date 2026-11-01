//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what openRuns — open the named runs of a stream, each past the entries already taken from it (nil: none), positioned at its next entry
package pack

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

// openRuns opens the named runs in out's working folder with a runBufSize
// buffer each. taken (nil: all zero) is how many entries a merge took from
// each before it stopped; the run is reopened after them and its next entry
// read. On error the files opened so far are closed.
func openRuns(out string, names []string, taken []int64) ([]*runReader, error) {
	var rs []*runReader
	for i, n := range names {
		var skip int64
		if taken != nil {
			skip = taken[i]
		}
		f, err := os.Open(filepath.Join(out, streamDir, n))
		if err == nil {
			rr := &runReader{f: f, r: bufio.NewReaderSize(io.NewSectionReader(f, skip*hEntrySize, 1<<62), runBufSize), taken: skip}
			rs = append(rs, rr)
			err = rr.advance()
		}
		if err != nil {
			closeRuns(rs)
			return nil, err
		}
	}
	return rs, nil
}
