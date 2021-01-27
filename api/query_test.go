package api_test

import (
	"testing"

	"github.com/han-tyumi/mcf/api"
)

func TestAll(t *testing.T) {
	if _, err := api.Many(&api.ManyParams{}); err != nil {
		t.Fatal(err)
	}
}

func TestOne(t *testing.T) {
	if _, err := api.One(238222); err != nil {
		t.Fatal(err)
	}
}
