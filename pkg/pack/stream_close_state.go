//ff:type feature=pack type=model
//ff:what streamClose — how far Close has come, saved with the checkpoint: phase, finished layers, merged-away runs, the H pass position (entries taken per run, pending leaf, keys, written heads)
package pack

// streamClose is Close's progress (streamState.Close), so a stopped Close
// goes on where it last saved (ResumeStream, then Close again). Phase is one
// of closeNone…closeHDone. P and H are the finished layers. Spent are runs
// merged into later ones: kept on disk (not stray) until the set is done.
// During the H pass Taken is how many entries each of Runs gave, HLeaf the
// leaf being filled, HKeys the distinct keys so far and HHeads how many leaf
// heads streamHeadsFile holds.
type streamClose struct {
	Phase  int
	P      []streamLayer
	H      []streamLayer
	Spent  []string
	Taken  []int64
	HLeaf  []streamKV
	HKeys  uint64
	HHeads int64
}
