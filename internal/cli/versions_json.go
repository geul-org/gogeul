//ff:type feature=cli type=model
//ff:what versionsJSON — geul show versions as JSON: the versions and NEXT-LOCAL-ID
package cli

// versionsJSON is a document's history (resolve.History) for --format json.
type versionsJSON struct {
	Versions    []versionJSON `json:"versions"`
	NextLocalID uint64        `json:"next_local_id"`
}
