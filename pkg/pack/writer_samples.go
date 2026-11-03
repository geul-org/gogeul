//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.samples — records to train a dictionary on and others to try it on, alternating in SHA-256 order, each side capped
package pack

// samples reads up to DefaultDictSamples records: even picks train (up to
// dictTrainBytes), odd ones test (up to a quarter of that).
func (w *writer) samples() ([][]byte, [][]byte, error) {
	var train, test [][]byte
	var nTrain, nTest int
	for i, p := range sampleOrder(w.records()) {
		if i >= DefaultDictSamples || nTrain >= dictTrainBytes {
			break
		}
		b, err := w.b.Get(p.src)
		if err != nil {
			return nil, nil, err
		}
		if i%2 == 0 {
			train, nTrain = append(train, b), nTrain+len(b)
		} else if nTest < dictTrainBytes/4 {
			test, nTest = append(test, b), nTest+len(b)
		}
	}
	return train, test, nil
}
