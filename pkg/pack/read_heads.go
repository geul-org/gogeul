//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what readHeads — n block heads (H run entry format) from a head file, starting at head from
package pack

import "io"

// readHeads reads heads from…from+n-1 of f as keys and locations.
func readHeads(f io.ReaderAt, from, n int64) ([]streamKV, error) {
	b := make([]byte, n*hEntrySize)
	if _, err := f.ReadAt(b, from*hEntrySize); err != nil {
		return nil, err
	}
	out := make([]streamKV, 0, n)
	for i := int64(0); i < n; i++ {
		e := parseHEntry(b[i*hEntrySize : (i+1)*hEntrySize])
		out = append(out, streamKV{Key: string(e.key[:]), Loc: e.loc})
	}
	return out, nil
}
