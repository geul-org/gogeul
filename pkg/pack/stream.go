//ff:type feature=pack type=writer
//ff:what Stream — a pack set written in one pass from records fed in key order, with bounded memory and checkpoints
package pack

import (
	"hash"
	"os"
)

// Stream writes the first generation of a pack set from P records fed in
// ascending key order (Put) and a few named records (PutNamed), for archives
// too big to stage as a directory tree (geul-dev files/pack/README.md §6,
// "대량 쓰기"). Data packs are cut at PackSize as records arrive; P leaf
// blocks are written as soon as they fill (keys arrive sorted); H keys, which
// arrive in random order, are spilled to sorted run files and merged when
// the stream closes. Memory therefore holds one P leaf, one H run, one data
// pack's inventory and the first keys of the written blocks, not the records.
// Checkpoint saves everything needed to go on (ResumeStream) after a crash;
// the bytes written do not depend on when checkpoints happen. Close saves
// its own progress the same way, so a stopped Close goes on after
// ResumeStream; its memory does not grow with the number of runs.
// dsum and isum hash the open data pack and index object as they are
// written (streamOpen.Sum).
type Stream struct {
	out   string
	o     StreamOptions
	st    streamState
	inum  map[string]uint32
	dict  *Dict
	data  *os.File
	index *os.File
	hbuf  []hEntry
	dsum  hash.Hash
	isum  hash.Hash
}
