//ff:func feature=core type=decoder control=selection
//ff:what quantityForm — whether the four QUANTITY records have their forms: a numeric amount, Reference or Absent unit, numeric or Absent bounds
package core

import "errors"

// quantityForm checks the forms of a QUANTITY's amount, unit, lower and upper
// records (value/numeric/QUANTITY). The amount and each present bound are one
// of the numeric Values quantityNumber lists. The unit is a Reference of a
// real Codec or ABSENT; no other slot marker (Reference C < 3) is a unit.
func quantityForm(rs [4]Record) error {
	switch {
	case !quantityNumber(rs[0]):
		return errors.New("geul/core: QUANTITY amount is not an INTEGER, BINARY-FLOAT, FRACTION or DECIMAL Value")
	case rs[1].Kind != KindReference, IsSlotMarkerCodec(rs[1].Codec) && !IsAbsent(rs[1]):
		return errors.New("geul/core: QUANTITY unit is not a Reference or Absent")
	case !quantityNumber(rs[2]) && !IsAbsent(rs[2]), !quantityNumber(rs[3]) && !IsAbsent(rs[3]):
		return errors.New("geul/core: QUANTITY bound is not an INTEGER, BINARY-FLOAT, FRACTION or DECIMAL Value or Absent")
	default:
		return nil
	}
}
