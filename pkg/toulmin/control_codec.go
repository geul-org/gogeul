//ff:func feature=toulmin type=encoder control=selection
//ff:what ControlCodec — the direct Edge codec of a framing control
package toulmin

import "fmt"

// ControlCodec returns the Edge codec a framing control is written under.
func ControlCodec(c ControlCode) (uint16, error) {
	switch c {
	case ControlSection:
		return CodecSection, nil
	case ControlCaseBegin:
		return CodecCaseBegin, nil
	case ControlCaseEnd:
		return CodecCaseEnd, nil
	case ControlProse:
		return CodecProse, nil
	default:
		return 0, fmt.Errorf("geul/toulmin: control 0x%02X has no record", uint16(c))
	}
}
