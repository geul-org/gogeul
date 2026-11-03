//ff:func feature=cli type=command control=iteration dimension=1
//ff:what app.versions — print one document's versions with time, hash and state, then its NEXT-LOCAL-ID
package cli

import (
	"fmt"
)

// versions prints one line per version (resolve.Dir.Log): version, publication
// time (UTC), content hash, and deprecated or deleted; then NEXT-LOCAL-ID.
// JSON is versionsJSONOf.
func (a *app) versions(addr string) error {
	k, err := a.key(addr)
	if err != nil {
		return err
	}
	d, err := a.dir()
	if err != nil {
		return err
	}
	h, err := d.Log(k.Issuer, k.Doc)
	if err != nil {
		return err
	}
	if done, err := a.json(versionsJSONOf(h)); done {
		return err
	}
	for _, e := range h.Entries {
		fmt.Fprintf(a.out, "%d\t%s\t0x%X%s\n", e.Version, versionTime(e.At), e.Hash[:], stateOf(e.VersionEntry))
	}
	fmt.Fprintf(a.out, "NEXT-LOCAL-ID %d\n", h.Next)
	return nil
}
