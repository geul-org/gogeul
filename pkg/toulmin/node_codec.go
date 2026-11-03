//ff:func feature=toulmin type=encoder control=selection
//ff:what NodeCodec — the record codec of a node type (rules split by role into RULE/COUNTER/EXCEPT-DECLARE)
package toulmin

import "fmt"

// NodeCodec returns the CodecRef a node of type t is written under (GEUL t002
// TOULMIN 코덱). A rule node's codec is chosen by its role; every other node type
// maps to one codec regardless of role.
func NodeCodec(t NodeType, role RoleCode) (CodecRef, error) {
	switch t {
	case NodeCase:
		return CodecRef{Codec: CodecCaseDeclare}, nil
	case NodeRule:
		return ruleCodec(role)
	case NodeTerm:
		return CodecRef{Codec: CodecSpecDeclare}, nil
	case NodeEndpoint:
		return CodecRef{Codec: CodecEntry}, nil
	case NodeTrigger:
		return CodecRef{Codec: CodecTrigger}, nil
	case NodeNamespace:
		return CodecRef{Codec: CodecNamespace}, nil
	case NodeFunc:
		return CodecRef{Codec: CodecFunc}, nil
	case NodeField:
		return CodecRef{Codec: CodecField}, nil
	case NodeAlias:
		return CodecRef{Codec: CodecAlias}, nil
	case NodeGroup:
		return CodecRef{Codec: CodecGroup}, nil
	default:
		return CodecRef{}, fmt.Errorf("geul/toulmin: unknown NodeType 0x%04X", uint16(t))
	}
}
