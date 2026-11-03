//ff:func feature=schema type=decoder control=sequence
//ff:what importSplit — the IMPORT ID words of an IMPORT-DOC after its DOC-POINTER; false when the DOC-POINTER does not decode
package schema

import "github.com/geul-org/gogeul/pkg/core"

// importSplit decodes the DOC-POINTER (or CODEC-POINTER) that leads an IMPORT-DOC payload and
// returns the words after it, the IMPORT ID.
func importSplit(r rec) ([]core.Word, bool) {
	ref, n, err := core.DecodeRecord(r.R.Payload)
	if err != nil {
		return nil, false
	}
	if _, _, _, err := core.DecodeTargetPointer(ref); err != nil {
		return nil, false
	}
	return r.R.Payload[n:], true
}
