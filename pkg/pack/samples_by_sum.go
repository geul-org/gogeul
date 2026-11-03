//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what samplesBySum — dictionary samples in SHA-256 order of their bytes, so the even/odd train/test split does not follow the caller's order
package pack

import (
	"bytes"
	"crypto/sha256"
	"slices"
)

// samplesBySum returns bs reordered by the SHA-256 of each sample, as Write
// orders its samples (sampleOrder). A Stream's samples come in Put order,
// which alternates a version list and a version (`versions.geul`, `<V>.geul`,
// …), so taking the even positions in that order trained the dictionary on
// version lists alone. Equal sums keep their input order. The result depends
// only on the samples.
func samplesBySum(bs [][]byte) [][]byte {
	sums := make([][32]byte, len(bs))
	idx := make([]int, len(bs))
	for i, b := range bs {
		sums[i], idx[i] = sha256.Sum256(b), i
	}
	slices.SortStableFunc(idx, func(x, y int) int { return bytes.Compare(sums[x][:], sums[y][:]) })
	out := make([][]byte, len(bs))
	for i, j := range idx {
		out[i] = bs[j]
	}
	return out
}
