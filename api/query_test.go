package api_test

import (
	"testing"

	"github.com/han-tyumi/mcf/api"
)

func TestSearch(t *testing.T) {
	if _, err := api.Search(&api.SearchParams{}); err != nil {
		t.Fatal(err)
	}
}

func TestOne(t *testing.T) {
	if _, err := api.One(238222); err != nil {
		t.Fatal(err)
	}
}

func TestMany(t *testing.T) {
	if _, err := api.Many([]uint{238222}); err != nil {
		t.Fatal(err)
	}
}
