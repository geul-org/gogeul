//ff:func feature=core type=helper control=iteration dimension=1
//ff:what loadVector — read testdata/migration1/vector2.tsv into rows
package core

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

// loadVector reads the migration1 reference vector (old → new TAG WORD).
func loadVector(t *testing.T) []vectorRow {
	f, err := os.Open("testdata/migration1/vector2.tsv")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var rows []vectorRow
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		a := strings.Split(sc.Text(), "\t")
		old, err := strconv.ParseUint(a[0], 16, 16)
		if err != nil {
			continue
		}
		n, err := strconv.ParseInt(a[1], 16, 32)
		if err != nil {
			n = -1
		}
		rows = append(rows, vectorRow{Old: Word(old), New: int(n), Name: a[2], OldKey: a[3]})
	}
	return rows
}
