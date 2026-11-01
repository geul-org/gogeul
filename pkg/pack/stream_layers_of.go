//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what streamLayersOf — layers as checkpoint values (streamLayer)
package pack

// streamLayersOf copies count, encoding and root of each layer.
func streamLayersOf(ls []layer) []streamLayer {
	var out []streamLayer
	for _, l := range ls {
		out = append(out, streamLayer{Count: l.count, Enc: l.enc, Root: l.root})
	}
	return out
}
