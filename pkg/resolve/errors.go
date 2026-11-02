package resolve

import "github.com/geul-org/gogeul/pkg/pack"

var (
	// ErrNotFound means a source does not have the document or version (404, 410,
	// a missing file). Resolution moves on to the next source. It is
	// pack.ErrNotFound, the one value every storage form reports.
	ErrNotFound = pack.ErrNotFound
	// ErrVerify means fetched bytes do not match the version list or the Key.
	// Resolution never falls through to another source on it. It is
	// pack.ErrVerify.
	ErrVerify = pack.ErrVerify
)
