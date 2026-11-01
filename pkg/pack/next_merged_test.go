//ff:func feature=pack type=writer control=sequence
//ff:what TestNextMerged — the least front entry comes first, counted as taken, and its run moves on; used-up runs give false; a cut run fails
package pack

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func TestNextMerged(t *testing.T) {
	var a, b hEntry
	a.key[0], b.key[0] = 1, 2
	run := func(es ...hEntry) *runReader {
		var buf []byte
		for _, e := range es {
			buf = appendHEntry(buf, e)
		}
		rr := &runReader{f: os.Stdin, r: bufio.NewReader(bytes.NewReader(buf))}
		rr.advance()
		return rr
	}
	rs := []*runReader{run(b), run(a)}
	if e, ok, err := nextMerged(rs); !ok || err != nil || e.key[0] != 1 {
		t.Fatal("first")
	}
	if rs[1].taken != 1 || rs[0].taken != 0 {
		t.Fatal("taken")
	}
	if e, ok, _ := nextMerged(rs); !ok || e.key[0] != 2 {
		t.Fatal("second")
	}
	if _, ok, _ := nextMerged(rs); ok {
		t.Fatal("used up")
	}
	cut := &runReader{f: os.Stdin, r: bufio.NewReader(bytes.NewReader(append(appendHEntry(nil, a), 1, 2)))}
	cut.advance()
	if _, _, err := nextMerged([]*runReader{cut}); err == nil {
		t.Fatal("cut run: want error")
	}
}
