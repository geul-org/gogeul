//ff:type feature=pack type=model
//ff:what streamLayer — a finished P or H layer (count, root encoding, stored root) with exported fields for a checkpoint
package pack

// streamLayer mirrors layer (count, enc, root) so Close can save a finished
// layer (gob); the decoded root is not needed to write the manifest.
type streamLayer struct {
	Count uint64
	Enc   byte
	Root  []byte
}
