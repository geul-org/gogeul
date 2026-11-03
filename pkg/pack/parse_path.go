//ff:func feature=pack type=helper control=selection
//ff:what parsePath — a 공표 주소 path into its key space: P (docs versions and version lists), H (hashes), named (the Issuer list and codecs), none otherwise
package pack

import "strings"

// parsePath maps a path relative to `.well-known/geul/` (§3.13 공표 주소) to
// its key space: `<I>/docs/<D>/<V>.geul` and `<I>/docs/<D>/versions.geul` to
// P, `<I>/hashes/<H>.geul` to H, the Issuer list (`issuers/`) and the codecs
// (`<I>/codecs/<C>/`) to the named space. latest.geul is a reserved name
// served by reading the version list, never stored, and every other spelling
// is outside the forms — `<I>/docs/<V>.geul` and `<I>/codecs/<V>.geul` too,
// as there is no document or codec list (§1.1 101): all are spaceNone.
// The Issuer may be zero-padded and the hex digits of any case (§1.1 122);
// ref.issuer and ref.path are canonical (A-F upper case, §1.1 120).
func parsePath(p string) pathRef {
	seg := canonSegs(strings.Split(p, "/"))
	if len(seg) == 2 && seg[0] == issuersDir && storedTail(seg[1]) {
		return pathRef{space: spaceNamed, path: strings.Join(seg, "/")}
	}
	if len(seg) < 3 {
		return pathRef{}
	}
	is, ok := normalIssuer(seg[0])
	if !ok {
		return pathRef{}
	}
	seg[0] = is
	ref := pathRef{issuer: is, path: strings.Join(seg, "/")}
	switch {
	case seg[1] == hashesDir && len(seg) == 3:
		return parseHPath(ref, seg[2])
	case seg[1] == docsDir && len(seg) == 4:
		return parsePPath(ref, seg[2], seg[3])
	case seg[1] == codecsDir && len(seg) == 4 && isHex(seg[2], 8) && storedTail(seg[3]):
		ref.space = spaceNamed
		return ref
	default:
		return pathRef{}
	}
}
