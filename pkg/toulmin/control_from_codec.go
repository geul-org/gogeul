//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what ControlFromCodec — the framing control a direct Edge codec declares
package toulmin

// controls lists every framing control ControlFromCodec searches.
var controls = []ControlCode{ControlSection, ControlCaseBegin, ControlCaseEnd, ControlProse}

// ControlFromCodec reports the framing control whose record uses Edge codec
// f. ok is false when f is not a control codec.
func ControlFromCodec(f uint16) (ControlCode, bool) {
	for _, c := range controls {
		if cf, _ := ControlCodec(c); cf == f {
			return c, true
		}
	}
	return 0, false
}
