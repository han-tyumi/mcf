package api_test

import (
	"testing"

	"github.com/han-tyumi/mcf/api"
)

func TestAll(t *testing.T) {
	if _, err := api.All(&api.AllQuery{}); err != nil {
		t.Fatal(err)
	}
}

func TestOne(t *testing.T) {
	if _, err := api.One(238222); err != nil {
		t.Fatal(err)
	}
}
