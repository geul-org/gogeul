//ff:func feature=pack type=helper control=sequence
//ff:what sampleManifest — a manifest with every section filled, for encode/parse round trips
package pack

// sampleManifest returns a small but complete manifest.
func sampleManifest() *manifest {
	root := zstdFrame(encodeBlock(block{kind: blkPLeaf, keys: []string{pKey(0, 8, 0)}, locs: []Loc{{Off: 64, Len: 3}}}))
	return &manifest{
		gen: 2, created: 1_800_000_000, prev: [32]byte{1},
		objects: []object{{kind: kindData, size: 100, hash: [32]byte{2}}, {kind: kindIndex, size: 64, hash: [32]byte{3}}},
		issuers: []string{"0x00000000", "0x00000001"},
		dicts:   []dictEntry{{id: 40000, loc: Loc{Off: 64, Len: 10}, hash: [32]byte{4}}},
		named:   []namedEntry{{"a.geul", Loc{Off: 67, Len: 1}}, {"b.geul", Loc{Off: 68, Len: 2, Enc: 1}}},
		pLayers: []layer{{count: 1, enc: encZstd, root: root}},
		hLayers: []layer{},
	}
}
