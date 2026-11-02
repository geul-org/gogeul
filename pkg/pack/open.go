//ff:func feature=pack type=resolver control=sequence
//ff:what Open — read latest.txt (unless pinned), fetch and hash-check the manifest, parse it, decode the layer roots
package pack

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

// Open opens one generation of the pack set in s. The manifest's SHA-256 must
// be its name and, when read through latest.txt, its generation the one
// latest.txt names and at least o.MinGeneration. Only a missing latest.txt (or
// pinned manifest) is ErrNotFound; a latest.txt whose manifest is missing is
// ErrCorrupt, so "no set here" and "a broken set" stay apart.
func Open(ctx context.Context, s Store, o Options) (*Reader, error) {
	name, gen := o.Manifest, uint64(0)
	if name == "" {
		b, err := s.ReadAll(ctx, latestFile)
		if err != nil {
			return nil, err
		}
		if gen, name, err = parseLatest(b); err != nil {
			return nil, err
		}
	}
	b, err := s.ReadAll(ctx, "m/"+name+".gpm")
	if errors.Is(err, ErrNotFound) && gen != 0 {
		return nil, corrupt("latest.txt names a missing manifest")
	}
	if err != nil {
		return nil, err
	}
	if sum := sha256.Sum256(b); hex.EncodeToString(sum[:]) != name {
		return nil, corrupt("manifest SHA-256 is not its name")
	}
	m, err := parseManifest(b)
	if err != nil {
		return nil, err
	}
	if gen != 0 && m.gen != gen {
		return nil, corrupt("latest.txt generation differs from the manifest")
	}
	if m.gen < o.MinGeneration {
		return nil, fmt.Errorf("%w: generation %d < %d", ErrRollback, m.gen, o.MinGeneration)
	}
	r := newReader(s, o, m, name)
	if err := r.decodeLayers(ctx, m.pLayers, true); err != nil {
		return nil, err
	}
	if err := r.decodeLayers(ctx, m.hLayers, false); err != nil {
		return nil, err
	}
	return r, nil
}
