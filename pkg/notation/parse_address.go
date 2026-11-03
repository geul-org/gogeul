//ff:func feature=notation type=encoder control=selection
//ff:what parseAddress — <Issuer>/<문서>[@v] as DOC-POINTER, <Issuer>/<문서>#n as EXTERNAL-ISSUER-* of the Issuer's width
package notation

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// parseAddress resolves the address (ctx's Lookup, IDs only without one). A
// `#n` suffix points at that LOCAL of the Issuer document (latest version,
// no @version) with the EXTERNAL-ISSUER codec of the Issuer's width.
func parseAddress(tok string, ctx Context) (core.Record, error) {
	addr, local, isLocal := strings.Cut(tok, "#")
	k, err := lookupOf(ctx)(addr)
	switch {
	case err != nil:
		return core.Record{}, err
	case !isLocal:
		return core.EncodeTargetPointer(k.Issuer, k.Doc, []core.Word{core.Word(k.Version >> 16), core.Word(k.Version)})
	case k.Version != 0:
		return core.Record{}, fmt.Errorf("geul/notation: %q: an Issuer document's LOCAL follows the latest version, no @", tok)
	}
	n, err := strconv.ParseUint(local, 10, 64)
	if err != nil || n == 0 {
		return core.Record{}, fmt.Errorf("geul/notation: %q needs a LOCAL ID above 0 after #", tok)
	}
	cc, ok := map[int]uint16{2: core.CodecExternalIssuer2Word, 4: core.CodecExternalIssuer4Word, 16: core.CodecExternalIssuer16Word}[len(k.Issuer)]
	if !ok || len(k.Doc) != 4 {
		return core.Record{}, fmt.Errorf("geul/notation: %q: an EXTERNAL-ISSUER needs a 2/4/16-word Issuer and a document", tok)
	}
	payload := append(append(append([]core.Word{}, k.Issuer...), k.Doc...), core.EncodeUintValue(n).Payload...)
	return core.Record{Kind: core.KindReference, Codec: cc, Payload: payload}, nil
}
