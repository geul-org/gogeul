//ff:func feature=resolve type=helper control=selection
//ff:what Addr.folder — the folder of an address relative to `.well-known/geul/`: issuers, <I>/codecs/<C>, <I>/docs/<D> or <I>/hashes
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// folder is the folder a's file lives in, relative to `.well-known/geul/`.
func (a Addr) folder() string {
	is := core.FormatID(a.Issuer)
	switch a.Kind {
	case AddrIssuers:
		return issuersDir
	case AddrCodec:
		return is + "/" + codecsDir + "/" + core.FormatID(a.ID)
	case AddrDoc:
		return is + "/" + docsDir + "/" + core.FormatID(a.ID)
	default:
		return is + "/" + hashesDir
	}
}
