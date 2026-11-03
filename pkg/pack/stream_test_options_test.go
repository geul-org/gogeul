//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what streamTestOptions — small packs, leaves and runs so a few dozen records cross every boundary; 40 samples train a dictionary
package pack

// streamTestOptions makes a 40-document stream use several data packs, P
// leaves under a root, three H levels and several runs.
func streamTestOptions() StreamOptions {
	var samples [][]byte
	for i := uint64(1); i <= 40; i++ {
		samples = append(samples, streamTestDoc(i, 1))
	}
	return StreamOptions{Build: BuildOptions{PackSize: 3000, PLeaf: 5, HLeaf: 4, HFanout: 3, HRoot: 3},
		Issuers: []string{"0x00000001"}, Samples: samples, RunEntries: 7, Created: 1_800_000_000}
}
