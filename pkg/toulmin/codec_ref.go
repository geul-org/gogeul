//ff:type feature=toulmin type=model
//ff:what CodecRef — where a TANGEUL record type lives: its direct Edge codec
package toulmin

// CodecRef identifies a TANGEUL record type on the wire: its direct TOULMIN
// Edge codec (Edge Codec = code - 0xA0; 0x060~0x07B and, since migration.md
// §3.14, 0x080~0x094). TANGEUL has no ISSUER-format records any more.
type CodecRef struct {
	Codec uint16
}
