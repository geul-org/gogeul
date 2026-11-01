package pack

const (
	// DefaultRunEntries is how many H keys a Stream sorts in memory per run (about 120 MB).
	DefaultRunEntries = 1 << 21
	// DefaultMergeRuns is how many H runs Close reads at once (StreamOptions.MergeRuns).
	// Each open run holds a runBufSize read buffer, so the merge holds
	// MergeRuns × runBufSize (64 MiB) whatever the number of runs.
	DefaultMergeRuns = 64
	// runBufSize is the read buffer of one H run during a merge (and the write
	// buffer of a merged run or the H head file).
	runBufSize = 1 << 20
	// hSaveLeaves is how many H leaves Close writes between two saves of its
	// position, so a stop loses at most that much work (about a million keys).
	hSaveLeaves = 1024
	// hEntrySize is the bytes of one H run entry: key 32, object 4, offset 8, length 8, encoding 1.
	hEntrySize = 53
	// streamDir is a Stream's working folder inside the set folder.
	streamDir = ".stream"
)

const (
	// streamStateFile is the checkpoint inside streamDir.
	streamStateFile = "state"
	// streamDataPart and streamIndexPart are the open data pack and index object inside streamDir.
	streamDataPart  = "data.part"
	streamIndexPart = "index.part"
	// streamHeadsFile holds the first key and location of each written H
	// leaf (H run entry format) while Close writes the H layer.
	streamHeadsFile = "hheads"
)

const (
	// closeNone: Close has not begun.
	closeNone = iota
	// closeStarted: every record is in (the H keys spilled); puts of keys
	// up to the last one are acknowledged without storing.
	closeStarted
	// closePDone: the data packs are closed and the P layer is in streamClose.P.
	closePDone
	// closeHPass: the runs are merged down to MergeRuns and the H leaves are
	// being written from the position in streamClose.
	closeHPass
	// closeHDone: the H layer is in streamClose.H.
	closeHDone
)
