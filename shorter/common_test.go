package shorter

import (
	"strings"
	"testing"
)

// Test ToShortId
func TestToShortId(t *testing.T) {
	var seq int
	var res int

	res = strings.Compare("aaaaaa", ToShortId(seq))

	if res != 0 {
		t.Fatalf("ToShartId failed for %q", seq)
	}

	// 62 ** 6 - 1
	seq = 56800235583
	res = strings.Compare("000000", ToShortId(seq))

	if res != 0 {
		t.Fatalf("ToShartId failed for %q", seq)
	}

}
