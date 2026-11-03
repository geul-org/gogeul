//ff:func feature=toulmin type=encoder control=selection
//ff:what ControlFormat — the GEUL ISSUER format number of a framing control
package toulmin

import "fmt"

// ControlFormat returns the format number a framing control is written under.
func ControlFormat(c ControlCode) (uint32, error) {
	switch c {
	case ControlSection:
		return FormatSection, nil
	case ControlCaseBegin:
		return FormatCaseBegin, nil
	case ControlCaseEnd:
		return FormatCaseEnd, nil
	case ControlProse:
		return FormatProse, nil
	default:
		return 0, fmt.Errorf("geul/toulmin: control 0x%02X has no record", uint16(c))
	}
}
