//ff:func feature=resolve type=helper control=selection
//ff:what listKey — the DocKey of a list document: the reserved list ID under its Issuer (the Issuer list under geul.org)
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// listKey is the DocKey of the list a names (AddrIssuers, AddrDocs or
// AddrCodecs): its reserved document ID (core.DocIDIssuers, DocIDDocs,
// DocIDCodecs) under its Issuer, the Issuer list always under geul.org
// 0x00000000. Its versions and versions.geul live in the list's folder
// (docFolder).
func listKey(a Addr) DocKey {
	switch a.Kind {
	case AddrIssuers:
		return DocKey{Issuer: geulIssuer, Doc: listIDIssuers}
	case AddrDocs:
		return DocKey{Issuer: core.FormatID(a.Issuer), Doc: listIDDocs}
	default:
		return DocKey{Issuer: core.FormatID(a.Issuer), Doc: listIDCodecs}
	}
}
