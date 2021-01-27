package mcf

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// BaseURL is the base URL for all Minecraft CurseForge queries.
const BaseURL = "https://addons-ecs.forgesvc.net/api/v2/addon"

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

// SearchParams specify the parameters for searching for mods.
type SearchParams struct {
	Version  string   `url:"gameVersion,omitempty"`
	Search   string   `url:"searchFilter,omitempty"`
	Page     uint     `url:"index,omitempty"`
	PageSize uint     `url:"pageSize,omitempty"`
	Sort     SortType `url:"sort,omitempty"`
}

// Search returns all Minecraft CurseForge mods based on query parameters.
func Search(q *SearchParams) ([]Mod, error) {
	v, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	url := BaseURL + "/search?gameId=432&sectionId=6"

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

// All returns all Minecraft CurseForge mods.
func All() ([]Mod, error) {
	return Search(&SearchParams{})
}

// One fetches a single Minecraft CurseForge mod by ID.
func One(id uint) (*Mod, error) {
	res, err := http.Get(fmt.Sprintf("%s/%d", BaseURL, id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	mod := &Mod{}
	d := json.NewDecoder(res.Body)
	if err := d.Decode(mod); err != nil && err != io.EOF {
		return nil, err
	}

	return mod, nil
}

// Many fetches multiple Minecraft CurseForge mods by ID.
func Many(ids []uint) ([]Mod, error) {
	body, err := json.Marshal(ids)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(BaseURL, "application/json", bytes.NewBuffer(body))
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
