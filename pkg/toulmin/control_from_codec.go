//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what ControlFromFormat — the framing control a GEUL ISSUER format number declares
package toulmin

// controls lists every framing control ControlFromFormat searches.
var controls = []ControlCode{ControlSection, ControlCaseBegin, ControlCaseEnd, ControlProse}

// ControlFromFormat reports the framing control whose record uses format number
// f. ok is false when f is not a control format.
func ControlFromFormat(f uint32) (ControlCode, bool) {
	for _, c := range controls {
		if cf, _ := ControlFormat(c); cf == f {
			return c, true
		}
	}
	return 0, false
}
