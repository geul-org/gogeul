//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what nextMerged — the smallest front entry over the runs (hEntryLess), counted as taken and its run advanced; false when all are used up
package pack

// nextMerged takes the least front entry of rs and moves that run on. A
// merge reads at most MergeRuns runs, so a scan is enough.
func nextMerged(rs []*runReader) (hEntry, bool, error) {
	var best *runReader
	for _, rr := range rs {
		if rr.ok && (best == nil || hEntryLess(rr.cur, best.cur)) {
			best = rr
		}
	}
	if best == nil {
		return hEntry{}, false, nil
	}
	e := best.cur
	best.taken++
	return e, true, best.advance()
}
