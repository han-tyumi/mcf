package mcf_test

import (
	"testing"

	"github.com/han-tyumi/mcf"
)

func TestSearch(t *testing.T) {
	if _, err := mcf.Search(&mcf.SearchParams{}); err != nil {
		t.Fatal(err)
	}
}

func TestOne(t *testing.T) {
	if _, err := mcf.One(238222); err != nil {
		t.Fatal(err)
	}
}

func TestMany(t *testing.T) {
	if _, err := mcf.Many([]uint{238222}); err != nil {
		t.Fatal(err)
	}
}
