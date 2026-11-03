//ff:func feature=toulmin type=encoder control=selection
//ff:what propMovedRef — the codec of a relation formerly written as a GEUL ISSUER format (0x12BL~0x134L, §3.14)
package toulmin

import "fmt"

// propMovedRef maps the relations outside the TOULMIN 0x100L~0x11BL codes to
// their direct codes 0x12BL~0x134L (formerly GEUL ISSUER formats, migration.md
// §3.14).
func propMovedRef(p PropCode) (CodecRef, error) {
	switch p {
	case PropRequires:
		return CodecRef{Codec: CodecRequires}, nil
	case PropMeans:
		return CodecRef{Codec: CodecMeans}, nil
	case PropHas:
		return CodecRef{Codec: CodecHas}, nil
	case PropRuns:
		return CodecRef{Codec: CodecRuns}, nil
	case PropEvaluates:
		return CodecRef{Codec: CodecEvaluates}, nil
	case PropTriggers:
		return CodecRef{Codec: CodecTriggers}, nil
	case PropUntil:
		return CodecRef{Codec: CodecUntil}, nil
	case PropSameAs:
		return CodecRef{Codec: CodecSameAs}, nil
	case PropImportedFrom:
		return CodecRef{Codec: CodecImportedFrom}, nil
	case PropLabel:
		return CodecRef{Codec: CodecLabel}, nil
	default:
		return CodecRef{}, fmt.Errorf("geul/toulmin: unknown PropCode 0x%04X", uint16(p))
	}
}
