package mcurse

// ModFile is the JSON response of a mod file retrieved from CurseForge.
type ModFile struct {
	ID          uint     `json:"id"`
	Versions    []string `json:"gameVersion"`
	Name        string   `json:"file_name"`
	Size        uint     `json:"fileLength"`
	ReleaseType uint     `json:"releaseType"`
	Slug        string   `json:"mod_key"`
	URL         string   `json:"downloadUrl"`
	Downloads   uint     `json:"download_count"`
	Uploaded    string   `json:"fileDate"`
	Available   bool     `json:"isAvailable"`
	Alternate   bool     `json:"isAlternate"`
	AlternateID uint     `json:"alternateFileId"`
}
