package mcurse

// Mod is the JSON response of a mod retrieved from CurseForge.
type Mod struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	URL          string    `json:"websiteUrl"`
	Summary      string    `json:"summary"`
	FileID       uint      `json:"defaultFileId"`
	Downloads    float64   `json:"downloadCount"`
	LatestFiles  []ModFile `json:"latestFiles"`
	Slug         string    `json:"slug"`
	Featured     bool      `json:"isFeatured"`
	Popularity   float64   `json:"popularityScore"`
	Rank         uint      `json:"gamePopularityRank"`
	Language     string    `json:"primaryLanguage"`
	Created      string    `json:"dateCreated"`
	Updated      string    `json:"dateModified"`
	Released     string    `json:"dateReleased"`
	Available    bool      `json:"isAvailable"`
	Experimental bool      `json:"isExperimental"`
}
