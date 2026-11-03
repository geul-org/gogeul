//ff:func feature=toulmin type=encoder control=selection
//ff:what propFormatRef — the GEUL ISSUER format of a relation without a direct TOULMIN code
package toulmin

import "fmt"

// propFormatRef maps the relations that have no direct TOULMIN Edge code to
// their GEUL ISSUER format numbers.
func propFormatRef(p PropCode) (CodecRef, error) {
	switch p {
	case PropRequires:
		return fmtRef(FormatRequires), nil
	case PropMeans:
		return fmtRef(FormatMeans), nil
	case PropHas:
		return fmtRef(FormatHas), nil
	case PropRuns:
		return fmtRef(FormatRuns), nil
	case PropEvaluates:
		return fmtRef(FormatEvaluates), nil
	case PropTriggers:
		return fmtRef(FormatTriggers), nil
	case PropUntil:
		return fmtRef(FormatUntil), nil
	case PropSameAs:
		return fmtRef(FormatSameAs), nil
	case PropImportedFrom:
		return fmtRef(FormatImportedFrom), nil
	case PropLabel:
		return fmtRef(FormatLabel), nil
	default:
		return CodecRef{}, fmt.Errorf("geul/toulmin: unknown PropCode 0x%04X", uint16(p))
	}
}
