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
	// latestFile is the reserved name of the newest version (공표 주소 latest.geul).
	latestFile = "latest.geul"
	// issuersDir holds geul.org's Issuer list (GEUL-ORG Issuer 이름), in place
	// of the old issuers.geul.
	issuersDir = "issuers"
	// docsDir holds an Issuer's documents, codecsDir its codec documents and
	// hashesDir its documents by content hash (in place of the old h/). There
	// is no document or codec list (§1.1 101).
	docsDir   = "docs"
	codecsDir = "codecs"
	hashesDir = "hashes"
	// listIDIssuers is the reserved Issuer list document ID
	// (core.DocIDIssuers) in canonical hex.
	listIDIssuers = "0xFFFFFFFFFFFFFFFF"
	// geulIssuer is the GEUL Issuer, whose folder holds the paths outside every
	// Issuer when an archive is kept as one set per Issuer (pack.Split).
	geulIssuer = "0x00000000"
	// setLatest, setManifests and setObjects are the entries of a pack set
	// folder. They are storage, not published paths.
	setLatest    = "latest.txt"
	setManifests = "m"
	setObjects   = "o"
)
