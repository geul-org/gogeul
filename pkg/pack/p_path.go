//ff:func feature=pack type=helper control=sequence
//ff:what pPath — the address of a P key: `<I>/docs/<D>/versions.geul` for slot 0, the version file otherwise
package pack

import "fmt"

// pPath is the path of (issuer, doc, slot) relative to `.well-known/geul/`.
func pPath(issuer string, d uint64, s uint32) string {
	if s == 0 {
		return fmt.Sprintf("%s/%s/0x%016X/%s", issuer, docsDir, d, listFile)
	}
	return fmt.Sprintf("%s/%s/0x%016X/0x%08X.geul", issuer, docsDir, d, s)
}
