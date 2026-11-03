//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what streamDictOf — the dictionary a Stream trains from its samples (SHA-256 order: even train, odd test), nil when the codec takes none, the samples are too few or it saves nothing
package pack

// streamDictOf trains a dictionary (ID FirstDictID; a stream starts a set)
// on the even samples in SHA-256 order (samplesBySum, as Write samples) and
// keeps it only when it makes the odd ones smaller in total than without it. Stored, Zstd with DictSize < 0 and fewer
// than dictMinSamples samples give none. The result depends only on the
// samples and the codec settings.
func streamDictOf(o StreamOptions) *Dict {
	z, ok := o.Build.Codec.(*Zstd)
	if !ok || z.DictSize < 0 || len(o.Samples) < dictMinSamples {
		return nil
	}
	var train, test [][]byte
	for i, b := range samplesBySum(o.Samples) {
		if i%2 == 0 {
			train = append(train, b)
		} else {
			test = append(test, b)
		}
	}
	d, err := z.Train(train, FirstDictID)
	if err != nil || !dictSaves(z, d, test) {
		return nil // too little or too uniform data, or no saving: records go without one
	}
	return d
}
