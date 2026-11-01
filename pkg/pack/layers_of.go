//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what layersOf — saved layers (streamLayer) back as manifest layers
package pack

// layersOf is streamLayersOf's inverse (the decoded root left empty; the
// manifest encoder uses the stored root).
func layersOf(sl []streamLayer) []layer {
	var out []layer
	for _, l := range sl {
		out = append(out, layer{count: l.Count, enc: l.Enc, root: l.Root})
	}
	return out
}
