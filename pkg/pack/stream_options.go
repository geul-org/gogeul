//ff:type feature=pack type=writer
//ff:what StreamOptions — a Stream's block sizes, codec, dictionary samples, Issuers, run size, merge width and Close's pace hooks
package pack

// StreamOptions tune a Stream. Build gives the pack and block sizes, the
// codec (nil is Zstd at its defaults) and the clock, as for Write; its
// Verify, Check and Fresh are not used. Issuers lists the Issuers (canonical
// hex) of the P records, in ISSUERS order. Samples are records a dictionary
// is trained on (even positions) and tried on (odd positions) before the
// first record, as Write samples a batch; fewer than 32 train none.
// RunEntries is how many H keys are held before a sorted run is spilled to
// disk (0 is DefaultRunEntries). Created is the manifest's created_at (Unix
// seconds; 0 takes Build.Now when the stream starts). MergeRuns is how many
// H runs Close merges at once (below 2: DefaultMergeRuns); more runs are
// merged into fewer first, so the merge holds MergeRuns × 1 MiB of read
// buffers however many runs the stream spilled. It does not change the bytes
// written. saveLeaves is how many H leaves Close writes between saves (below
// 1: hSaveLeaves; tests make it small).
//
// Pace, when set, is called at Close's safe points with the phase Close is in
// (close:data, close:p, close:merge, close:h, close:tree, close:index); it may
// block. An error stops Close there: what was done is saved, Close returns an
// error that is both ErrStopped and Pace's error, and ResumeStream and Close
// go on and write the same bytes. MergeWidth, when set, is asked before each
// merge how many runs it may read at once (clamped to 2…MergeRuns), so the
// merge can follow a memory budget that changes while Close runs.
type StreamOptions struct {
	Build      BuildOptions
	Issuers    []string
	Samples    [][]byte
	RunEntries int
	Created    int64
	MergeRuns  int
	Pace       func(phase string) error
	MergeWidth func() int
	saveLeaves int
}
