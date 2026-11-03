//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestCodecTargetOf — a codec spec document ID gives its target; ordinary documents, 0, a Length and bad hex do not
package resolve

import "testing"

func TestCodecTargetOf(t *testing.T) {
	if w, ok := codecTargetOf("0x00000b80"); !ok || w != 0x0B80 {
		t.Fatalf("got %04X %v", w, ok)
	}
	for _, s := range []string{"0xffff0001", "0x00000000", "0x00000010", "0x00000020", "0x00000b81", "0x00002b80", "0x0000000000000b80", "0x000b80", "0X00000B80"} {
		if _, ok := codecTargetOf(s); ok {
			t.Fatalf("%s: want false", s)
		}
	}
}
