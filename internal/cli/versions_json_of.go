//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what versionsJSONOf — a document's history as geul show versions JSON: hashes in hex, times in RFC 3339
package cli

import (
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// versionsJSONOf turns h into versionsJSON, writing each time as the text
// form does (versionTime) except that an unknown time is left out.
func versionsJSONOf(h resolve.History) versionsJSON {
	out := versionsJSON{Versions: []versionJSON{}, NextLocalID: h.Next}
	for _, e := range h.Entries {
		at := strings.TrimPrefix(versionTime(e.At), "-")
		out.Versions = append(out.Versions, versionJSON{Version: e.Version, At: at, Hash: fmt.Sprintf("0x%X", e.Hash[:]), State: strings.TrimPrefix(stateOf(e.VersionEntry), "\t")})
	}
	return out
}
