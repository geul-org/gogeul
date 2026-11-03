//ff:type feature=resolve type=model
//ff:what AddrKind — which of the 공표 주소 forms an Addr is: the Issuer list, a codec or document list, a codec, a document or a hash
package resolve

// AddrKind is the form of a 공표 주소 (§3.13): what Addr.Issuer, ID and Hash mean.
type AddrKind uint8

const (
	// AddrIssuers is `issuers/…`, geul.org's Issuer list.
	AddrIssuers AddrKind = iota + 1
	// AddrCodecs is `<Issuer>/codecs/…`, the Issuer's codec list.
	AddrCodecs
	// AddrCodec is `<Issuer>/codecs/<코덱 ID>/…`, one codec document.
	AddrCodec
	// AddrDocs is `<Issuer>/docs/…`, the Issuer's document list.
	AddrDocs
	// AddrDoc is `<Issuer>/docs/<문서 ID>/…`, one document.
	AddrDoc
	// AddrHash is `<Issuer>/hashes/<내용 해시>.geul`.
	AddrHash
)

// Latest is the version ID that asks for the newest version (`latest.geul`).
// Version IDs start at 1.
const Latest uint32 = 0
