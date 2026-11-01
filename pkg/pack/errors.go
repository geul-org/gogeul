package pack

import "errors"

var (
	// ErrCorrupt means pack bytes break the format or fail an integrity check
	// (object or manifest SHA-256, zstd checksum, content hash, truncation).
	ErrCorrupt = errors.New("geul/pack: corrupt")
	// ErrStopped means StreamOptions.Pace stopped a Close at a safe point:
	// ResumeStream and Close go on from there.
	ErrStopped = errors.New("geul/pack: Close stopped")
	// ErrUnsupported means valid bytes this implementation cannot read (a
	// compressed zstd block without a zstd codec).
	ErrUnsupported = errors.New("geul/pack: unsupported")
	// ErrRange means an HTTP server did not answer a Range request with exactly
	// the bytes asked for (200 instead of 206, a wrong Content-Range, encoding).
	ErrRange = errors.New("geul/pack: range request not honored")
	// ErrRollback means latest.txt names an older generation than required.
	ErrRollback = errors.New("geul/pack: generation went back")
	// ErrConflict means latest.txt changed while a new generation was written.
	ErrConflict = errors.New("geul/pack: latest.txt changed during write")
)
