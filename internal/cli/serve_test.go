//ff:func feature=cli type=command control=sequence
//ff:what TestServe — prints the URL and serves the page until the listener closes
package cli

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestServe(t *testing.T) {
	file := filepath.Join(t.TempDir(), "a.geul")
	os.WriteFile(file, core.PutWords(core.WrapDocument([]core.Word{0x1C40})), 0o644)
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	done := make(chan error)
	go func() { done <- serve(ln, file, &out) }()
	resp, err := http.Get("http://" + ln.Addr().String() + "/")
	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	ln.Close()
	if err := <-done; err == nil || resp.StatusCode != 200 || !strings.Contains(string(body), "a.geul") || !strings.Contains(out.String(), ln.Addr().String()) {
		t.Fatalf("serve: %v %d %q", err, resp.StatusCode, out.String())
	}
}
