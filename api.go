package mcf

import (
	"bytes"
	"encoding/json"
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
		return nil, apiErr("search", q, err)
	}

	url := BaseURL + "/search?gameId=432&sectionId=6"

	query := v.Encode()
	if len(query) != 0 {
		url += "&" + query
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, apiErr("search", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, apiErr("search", url, res.Status)
	}

	var mods []Mod
	d := json.NewDecoder(res.Body)
	if err := d.Decode(&mods); err != nil && err != io.EOF {
		return nil, apiErr("search", url, err)
	}

	return mods, nil
}

// All returns all Minecraft CurseForge mods.
func All() ([]Mod, error) {
	return Search(&SearchParams{})
}

// One fetches a single Minecraft CurseForge mod by ID.
func One(id uint) (*Mod, error) {
	url := fmt.Sprintf("%s/%d", BaseURL, id)
	res, err := http.Get(url)
	if err != nil {
		return nil, apiErr("one", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, apiErr("one", url, res.Status)
	}

	mod := &Mod{}
	d := json.NewDecoder(res.Body)
	if err := d.Decode(mod); err != nil && err != io.EOF {
		return nil, apiErr("one", url, err)
	}

	return mod, nil
}

// Many fetches multiple Minecraft CurseForge mods by ID.
func Many(ids []uint) ([]Mod, error) {
	body, err := json.Marshal(ids)
	if err != nil {
		return nil, apiErr("many", ids, err)
	}

	res, err := http.Post(BaseURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, apiErr("many", ids, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, apiErr("many", ids, res.Status)
	}

	var mods []Mod
	d := json.NewDecoder(res.Body)
	if err := d.Decode(&mods); err != nil && err != io.EOF {
		return nil, apiErr("many", ids, err)
	}

	return mods, nil
}

// Files fetches all the files for a Minecraft CurseForge mod by ID.
func Files(id uint) ([]ModFile, error) {
	url := fmt.Sprintf("%s/%d/files", BaseURL, id)
	res, err := http.Get(url)
	if err != nil {
		return nil, apiErr("files", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, apiErr("files", url, res.Status)
	}

	var files []ModFile
	d := json.NewDecoder(res.Body)
	if err := d.Decode(&files); err != nil && err != io.EOF {
		return nil, apiErr("files", url, err)
	}

	return files, nil
}

func apiErr(fn string, id, err interface{}) error {
	return fmt.Errorf("%s: %v: %v", fn, id, err)
}
