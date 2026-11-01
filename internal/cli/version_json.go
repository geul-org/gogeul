//ff:type feature=cli type=model
//ff:what versionJSON — one version of geul show versions as JSON: number, time, content hash in hex, state
package cli

// versionJSON is one version: At is RFC 3339 UTC ("" when unknown), Hash is
// 0x and 64 hex digits, State is "deprecated", "deleted" or "".
type versionJSON struct {
	Version uint32 `json:"version"`
	At      string `json:"at,omitempty"`
	Hash    string `json:"hash"`
	State   string `json:"state,omitempty"`
}
