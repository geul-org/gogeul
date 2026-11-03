//ff:func feature=toulmin type=encoder control=sequence
//ff:what fmtRef — the CodecRef of a GEUL ISSUER format number
package toulmin

// fmtRef returns the CodecRef of format number f under the GEUL Issuer.
func fmtRef(f uint32) CodecRef {
	return CodecRef{Codec: codecIssuer2Word, Format: f}
}
