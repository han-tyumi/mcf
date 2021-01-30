package mcf

// ModFile is the JSON response of a mod file retrieved from CurseForge.
type ModFile struct {
	ID            uint          `json:"id"`
	DisplayName   string        `json:"displayName"`
	Name          string        `json:"fileName"`
	Uploaded      string        `json:"fileDate"`
	Size          uint          `json:"fileLength"`
	ReleaseType   uint          `json:"releaseType"`
	Status        uint          `json:"fileStatus"`
	URL           string        `json:"downloadUrl"`
	Alternate     bool          `json:"isAlternate"`
	AlternateID   uint          `json:"alternateFileId"`
	Dependencies  []interface{} `json:"dependencies"`
	Available     bool          `json:"isAvailable"`
	Modules       []interface{} `json:"modules"`
	Fingerprint   uint          `json:"packageFingerprint"`
	Versions      []string      `json:"gameVersion"`
	Metadata      interface{}   `json:"installMetadata"`
	PackID        interface{}   `json:"serverPackFileId"`
	InstallScript bool          `json:"hasInstallScript"`
}
