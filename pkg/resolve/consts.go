package resolve

const (
	// DefaultProxy is the GEULPROXY used when the variable is empty (GEUL-ORG 출처 목록).
	DefaultProxy = "https://geul.org,registry"
	// geulOrg is the base of the public geul.org source, skipped for private documents.
	geulOrg = "https://geul.org"
	// wellKnown is the path every source serves documents under (RFC 8615).
	wellKnown = ".well-known/geul"
	// geulExt is the suffix of a binary GEUL document (GEUL-ORG 형식 지정).
	geulExt = ".geul"
	// listFile is a document's version list document (edge/PUBLICATION).
	listFile = "versions.geul"
	// namesFile is an Issuer's document name index (GEUL-ORG 문서 이름).
	namesFile = "documents.geul"
	// issuersFile is geul.org's Issuer name index (GEUL-ORG Issuer 이름).
	issuersFile = "issuers.geul"
	// geulIssuer is the GEUL Issuer, whose folder holds the paths outside every
	// Issuer when an archive is kept as one set per Issuer (pack.Split).
	geulIssuer = "0x0000"
	// setLatest, setManifests and setObjects are the entries of a pack set
	// folder. They are storage, not published paths.
	setLatest    = "latest.txt"
	setManifests = "m"
	setObjects   = "o"
)
