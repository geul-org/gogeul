//ff:type feature=pack type=model
//ff:what streamState — what a Stream checkpoints: target generation, objects, open object offsets, inventory, written block heads, pending leaf, H runs, named records, last key, Close's progress
package pack

// streamState is the part of a Stream that a checkpoint saves (gob). Objects
// are reserved in start order; an open one has Size 0 until it closes. Data
// and Index are the open data pack and index object (Num -1: none) with the
// bytes written so far (Off, Count). Inv and Paths are the open data pack's
// inventory. Dict and DictRaw are the dictionary records use. PHeads are the
// first key and location of each written P leaf, PLeaf the entries of the
// leaf being filled. Runs are the spilled sorted H runs (RunCount names the
// next). Named are the named records written. Last is the last P key put and
// Count how many P records there are. Meta is the caller's own progress,
// saved with the same checkpoint (Checkpoint, Stream.Meta). Close is how far
// Close has come (streamClose).
type streamState struct {
	Gen      uint64
	Created  int64
	Issuers  []string
	Objects  []streamObject
	Data     streamOpen
	Index    streamOpen
	Inv      []streamInv
	Paths    []string
	Dict     streamDict
	DictRaw  []byte
	PHeads   []streamKV
	PLeaf    []streamKV
	Runs     []string
	RunCount int
	Named    []streamKV
	Last     string
	Count    int64
	Res      BuildResult
	Meta     []byte
	Close    streamClose
}
