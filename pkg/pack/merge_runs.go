//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what mergeRuns — merge named H runs into one sorted run file (every entry kept), synced
package pack

import (
	"bufio"
	"os"
	"path/filepath"
)

// mergeRuns writes the entries of the runs names, in hEntryLess order and
// all of them (repeated keys are resolved by the final pass), to the run
// file dst in out's working folder, and syncs it. It holds one read buffer
// per input and one write buffer (runBufSize each).
func mergeRuns(out string, names []string, dst string) error {
	rs, err := openRuns(out, names, nil)
	if err != nil {
		return err
	}
	defer closeRuns(rs)
	f, err := os.Create(filepath.Join(out, streamDir, dst))
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriterSize(f, runBufSize)
	var b [hEntrySize]byte
	for {
		e, ok, err := nextMerged(rs)
		if err != nil {
			return err
		}
		if !ok {
			break
		}
		w.Write(appendHEntry(b[:0], e)) // errors stay in w for Flush
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return f.Sync()
}
