package mcurse

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// BaseURL is the base URL for all Minecraft CurseForge queries.
const BaseURL = "https://addons-ecs.forgesvc.net/api/v2/addon/"

// SortType represents the ways mods can be sorted.
type SortType uint

// Mod sorting types.
const (
	Featured SortType = iota
	Popularity
	LastUpdate
	Name
	Author
	TotalDownloads
)

// ModsOptions specify the parameters for a mods query.
type ModsOptions struct {
	Version  string   `url:"gameVersion,omitempty"`
	Search   string   `url:"searchFilter,omitempty"`
	Page     uint     `url:"index,omitempty"`
	PageSize uint     `url:"pageSize,omitempty"`
	Sort     SortType `url:"sort,omitempty"`
}

// Mods fetches all Minecraft CurseForge mods.
func Mods(o *ModsOptions) ([]Mod, error) {
	v, err := query.Values(o)
	if err != nil {
		return nil, err
	}

	url := BaseURL + "search?gameId=432&sectionId=6"

	query := v.Encode()
	if len(query) != 0 {
		url += "&" + query
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	var mods []Mod
	d := json.NewDecoder(res.Body)
	if err := d.Decode(&mods); err != nil && err != io.EOF {
		return nil, err
	}

	return mods, nil
}