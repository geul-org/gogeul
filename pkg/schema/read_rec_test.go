//ff:func feature=schema type=rule control=sequence
//ff:what TestReadRec — framing: prefix F4, Length F2/F3, ISSUER extension (Issuer + 2-word codec number) F5 and its canonical Issuer width
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestReadRec(t *testing.T) {
	if r, fs, ok := readRec([]core.Word{0x0621, 5}, 0, 0); !ok || fs != nil || r.Total != 2 || r.R.Payload[0] != 5 {
		t.Fatalf("plain: %+v %v %v", r, fs, ok)
	}
	if _, fs, _ := readRec([]core.Word{0x2421, 5}, 0, 0); rulesOf(fs) != "F4" || fs[0].Severity != Error {
		t.Fatalf("001: %+v", fs)
	}
	if _, fs, _ := readRec([]core.Word{0x4421, 5}, 0, 0); rulesOf(fs) != "F4" || fs[0].Severity != Error {
		t.Fatalf("010: %+v", fs)
	}
	if _, fs, ok := readRec([]core.Word{0x0622, 5}, 0, 0); ok || rulesOf(fs) != "F2" {
		t.Fatalf("short: %+v %v", fs, ok)
	}
	if _, fs, ok := readRec([]core.Word{0x062E, 3, 5}, 0, 0); !ok || rulesOf(fs) != "F3" {
		t.Fatalf("non-canonical: %+v %v", fs, ok)
	}
	if r, fs, ok := readRec([]core.Word{0x09D3, 0, 1, 0}, 0, 0); !ok || rulesOf(fs) != "F5" || len(r.R.Payload) != 3 {
		t.Fatalf("issuer truncated: %+v %v", fs, ok)
	}
	if r, fs, ok := readRec([]core.Word{0x09D6, 0, 2, 0, 0x0702, 7, 8}, 0, 0); !ok || fs != nil || r.R.Issuer[1] != 2 || r.R.Format != 0x0702 || len(r.R.Payload) != 2 {
		t.Fatalf("issuer: %+v %v", r, fs)
	}
	if r, fs, ok := readRec([]core.Word{0x05D6, 0, 2, 0x0001, 0x0002, 7, 8}, 0, 0); !ok || fs != nil || r.R.Kind != core.KindReference || r.R.Format != 0x00010002 {
		t.Fatalf("REFERENCE-ISSUER: %+v %v", r, fs)
	}
	if _, fs, ok := readRec([]core.Word{0x1BE7, 0, 1, 0, 0x20, 0, 0xC000, 7}, 0, 0); !ok || fs != nil {
		t.Fatalf("4-word issuer: %+v %v", fs, ok)
	}
	if _, fs, ok := readRec([]core.Word{0x1BE7, 0, 0, 0, 0x20, 0, 0xC000, 7}, 0, 0); !ok || rulesOf(fs) != "F5" {
		t.Fatalf("4-word issuer of a 2-word value: %+v %v", fs, ok)
	}
}
