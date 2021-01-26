package mcf_test

import (
	"testing"

	mcurse "github.com/han-tyumi/mcf"
)

func TestAll(t *testing.T) {
	if _, err := mcurse.All(&mcurse.ModsOptions{}); err != nil {
		t.Fatal(err)
	}
}
