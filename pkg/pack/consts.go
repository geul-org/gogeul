// Package pack stores a GEUL archive (`.well-known/geul/` paths → bytes) as
// immutable content-named objects: a manifest (`m/<sha256>.gpm`), index
// objects (`o/<sha256>.gpx`) and data packs (`o/<sha256>.gpk`), with
// `latest.txt` the only mutable file (geul-dev files/pack/README.md). The same
// objects serve a local directory (pread) and a CDN (one HTTP Range per read).
package pack

const (
	manifestMagic = "GEULMANI"
	objectMagic   = "GEULPACK"
	formatVersion = 1
	headerSize    = 64

	kindData  = 1 // data pack (.gpk)
	kindIndex = 2 // index object (.gpx)

	secObjects = 0x0001
	secIssuers = 0x0002
	secDicts   = 0x0003
	secNamed   = 0x0004
	secPLayers = 0x0005
	secHLayers = 0x0006

	encRaw  = 0
	encZstd = 1
	encGone = 0xFF // deletion marker in a leaf

	blkPLeaf  = 1
	blkPInner = 2
	blkHLeaf  = 3
	blkHInner = 4

	invP        = 1 // P record
	invH        = 2 // h/-only record
	invNamed    = 3 // named record
	invDict     = 4 // zstd dictionary
	invPDel     = 5 // P deletion
	invNamedDel = 6 // named deletion

	flagHash = 1 // the record also serves `<Issuer>/hashes/<content hash>.geul`

	spaceNone  = 0 // outside the 공표 주소 forms, or latest.geul: never stored
	spaceP     = 1
	spaceH     = 2
	spaceNamed = 3

	docsDir    = "docs"    // <Issuer>/docs/<문서 ID>/: documents (no document list)
	codecsDir  = "codecs"  // <Issuer>/codecs/<코덱 ID>/: codec documents (no codec list)
	hashesDir  = "hashes"  // <Issuer>/hashes/: documents by content hash
	issuersDir = "issuers" // issuers/: the Issuer list
	listFile   = "versions.geul"

	latestFile = "latest.txt"
	maxDepth   = 16        // directory depth guard against cycles
	maxBlock   = 128 << 10 // zstd Block_Maximum_Size
	zstdMagic  = 0xFD2FB528
	dictMagic  = 0xEC30A437 // formatted zstd dictionary (RFC 8878 5)
	hashChunk  = 8 << 20    // read size when hashing a whole object

	// DefaultPackSize is the target size of one data pack or index object (1 GiB).
	DefaultPackSize = 1 << 30
	// DefaultPLeaf is the entries per P leaf block.
	DefaultPLeaf = 16384
	// DefaultHLeaf is the entries per H leaf block.
	DefaultHLeaf = 1024
	// DefaultHFanout is the children per H internal block.
	DefaultHFanout = 1024
	// DefaultHRoot is the most children an H root (in the manifest) holds.
	DefaultHRoot = 1024
	// DefaultHLayers is how many H layers stack before they are merged.
	DefaultHLayers = 8
	// DefaultCacheBlocks is how many directory blocks a Reader keeps.
	DefaultCacheBlocks = 4096

	// DefaultZstdLevel is the zstd level of Zstd (§6: 19; klauspost/compress
	// runs it at its best speed setting).
	DefaultZstdLevel = 19
	// DefaultDictSize is the size of a newly trained dictionary (110 KiB, §9.3).
	DefaultDictSize = 110 << 10
	// DefaultDictSamples is how many records a dictionary is trained and tried on.
	DefaultDictSamples = 4096
	// DefaultWorkers is how many records a writer compresses at once when
	// Zstd.Workers is 0, or GOMAXPROCS when that is smaller (each best-level
	// encoder holds about 250 MB with a dictionary). Set Zstd.Workers (or
	// call SetWorkers) for more; StreamOptions.Build.Codec carries the Zstd.
	DefaultWorkers = 4
	// FirstDictID is the first zstd dictionary ID a writer gives (§4.6: 32768 ≤ ID < 2^31).
	FirstDictID = 1 << 15

	minZstd        = 256      // shorter input stays a raw-block frame (always single-segment, so it keeps its content size)
	dictMinSamples = 32       // fewer records train no dictionary
	dictTrainBytes = 16 << 20 // training input cap
	dictStatCount  = 256      // samples the dictionary's entropy tables are fitted on
	coverDmer      = 8        // bytes per counted substring when choosing dictionary content
	coverSegment   = 256      // bytes per chosen dictionary segment
	encodeBatch    = 64       // records per worker in one parallel encode round
)
