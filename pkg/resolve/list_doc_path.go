//ff:func feature=resolve type=helper control=sequence
//ff:what listDocPath — the stored address of a list document (issuers/, <Issuer>/docs/, <Issuer>/codecs/): its version 1
package resolve

// listDocPath is where a list document is kept: `<folder>/0x00000001.geul`,
// folder being `issuers`, `<Issuer>/docs` or `<Issuer>/codecs`. List documents
// carry no version history in this release (migration.md §8.4: the Issuer and
// document lists are not versioned yet), so the only version is 1, rewritten
// in place, and `latest.geul` of a list is that version.
func listDocPath(folder string) string {
	return folder + "/" + FormatVersion(1) + geulExt
}
