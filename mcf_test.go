package mcf_test

import (
	"testing"

	mcurse "github.com/han-tyumi/mcf"
)

func TestAll(t *testing.T) {
	if _, err := mcurse.All(&mcurse.AllQuery{}); err != nil {
		t.Fatal(err)
	}
}

func TestOne(t *testing.T) {
	if _, err := mcurse.One(238222); err != nil {
		t.Fatal(err)
	}
}
