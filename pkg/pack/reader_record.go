//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.Record — a record's original bytes by location: one range read, then unframe by encoding
package pack

import "context"

// Record reads and restores the record at l.
func (r *Reader) Record(ctx context.Context, l Loc) ([]byte, error) {
	b, err := r.raw(ctx, kindData, l)
	if err != nil {
		return nil, err
	}
	return r.unstore(ctx, l.Enc, b, false)
}
