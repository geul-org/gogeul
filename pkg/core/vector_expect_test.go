//ff:func feature=core type=helper control=selection
//ff:what vectorExpect — the Kind, spec name and L the vector table states for a new TAG name
package core

import (
	"strconv"
	"strings"
)

// vectorExpect reads a vector new_name ("E:TRIPLE/L9", "WIKIDATA-QID×Generic/L1",
// "V:ISSUER-125/L0", "ABSENT") into the Kind, the spec name, the Quantifier
// name ("" when none) and the L it states. ISSUER rows take their width from
// the old 7-bit code in oldKey (125/126/127 → 2/4/16 WORD).
func vectorExpect(name, oldKey string) (Kind, string, string, int) {
	body, l, _ := strings.Cut(name, "/L")
	length, _ := strconv.Atoi(l)
	prefix, rest, ok := strings.Cut(body, ":")
	width := map[string]string{"125": "2WORD", "126": "4WORD", "127": "16WORD"}[strings.Split(oldKey, "/")[1]]
	switch {
	case !ok:
		scheme, q, _ := strings.Cut(body, "×")
		return KindReference, scheme, q, length
	case prefix == "V" && strings.HasPrefix(rest, "ISSUER-"):
		return KindValue, "VALUE-ISSUER-" + width, "", length
	case prefix == "V":
		return KindValue, strings.TrimSuffix(strings.TrimSuffix(rest, "-EVEN"), "-ODD"), "", length
	case prefix == "E" && strings.HasPrefix(rest, "ISSUER-"):
		return KindEdge, "EDGE-ISSUER-" + width, "", length
	case prefix == "E":
		return KindEdge, rest, "", length
	default:
		return KindMeta, rest, "", length
	}
}
