//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.ManifestName — the opened manifest's SHA-256 hex (its name; pins the generation)
package pack

// ManifestName is the manifest's SHA-256 in lowercase hex (pack storage name, not a GEUL ID).
func (r *Reader) ManifestName() string {
	return r.name
}
