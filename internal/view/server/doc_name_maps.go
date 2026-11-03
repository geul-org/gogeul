//ff:func feature=server type=helper control=sequence
//ff:what docNameMaps — an Issuer's document list parsed once per content: name → ID (own documents) and ID → name, cached by SHA-256
package server

import (
	"crypto/sha256"
	"sync"

	"github.com/geul-org/gogeul/pkg/resolve"
)

var (
	nameMu    sync.Mutex
	nameCache = map[string]nameMaps{}
)

// docNameMaps reads the newest version of the document list `<space>/docs/`
// (GEUL-ORG 문서 이름) and keeps the
// names of that Issuer's own documents. Every page and listing needs them and
// the index is one document per Issuer (560 KB for 16,589 names), so parses
// are kept by the bytes' SHA-256: a changed file is parsed again, and the
// cache is emptied past nameCacheMax entries. A missing file gives none.
func docNameMaps(t resolve.Store, space string) nameMaps {
	b, err := latestDocs(t, space)
	if err != nil {
		return nameMaps{byName: map[string]string{}, byID: map[string]string{}}
	}
	sum := sha256.Sum256(b)
	key := space + "\x00" + string(sum[:])
	nameMu.Lock()
	m, ok := nameCache[key]
	nameMu.Unlock()
	if ok {
		return m
	}
	m = nameMaps{byName: ownNames(b, space)}
	m.byID = invertNames(m.byName)
	nameMu.Lock()
	if len(nameCache) >= nameCacheMax {
		clear(nameCache)
	}
	nameCache[key] = m
	nameMu.Unlock()
	return m
}
