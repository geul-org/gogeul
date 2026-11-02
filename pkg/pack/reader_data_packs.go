//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.DataPacks — the object numbers of the generation's data packs, in OBJECTS order (oldest first)
package pack

// DataPacks lists the objects PackInventory reads, in replay order.
func (r *Reader) DataPacks() []uint32 {
	var out []uint32
	for i, o := range r.m.objects {
		if o.kind == kindData {
			out = append(out, uint32(i))
		}
	}
	return out
}
