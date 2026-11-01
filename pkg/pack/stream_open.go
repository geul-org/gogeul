//ff:type feature=pack type=model
//ff:what streamOpen — a Stream's open object: number (-1: none), bytes written, items added, SHA-256 state of the bytes after the header
package pack

// streamOpen is an open object: its number (-1 when none is open), the
// bytes written so far and the items added. Sum is the SHA-256 state
// (encoding.BinaryMarshaler) of bytes headerSize…Off, so a resume can tell
// the object's part file, or the object sealed from it, from any other
// (checkpoints before this field have none: reopenPart checks their
// structure instead).
type streamOpen struct {
	Num   int
	Off   uint64
	Count uint64
	Sum   []byte
}
