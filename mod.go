package mcf

// ModLatestFile is the JSON response for a mod's latest files.
type ModLatestFile struct {
	ModFile
	Changelog             interface{} `json:"changelog"`
	Compatible            bool        `json:"isCompatibleWithClient"`
	PackageType           uint        `json:"categorySectionPackageType"`
	Access                uint        `json:"restrictProjectFileAccess"`
	ProjectStatus         uint        `json:"projectStatus"`
	CacheID               uint        `json:"renderCacheId"`
	LegacyMappingID       interface{} `json:"fileLegacyMappingId"`
	ProjectID             uint        `json:"projectId"`
	ParentProjectFileID   interface{} `json:"parentProjectFileId"`
	ParentLegacyMappingID interface{} `json:"parentFileLegacyMappingId"`
	TypeID                interface{} `json:"fileTypeId"`
	ExposeAsAlternative   interface{} `json:"exposeAsAlternative"`
	FingerprintID         uint        `json:"packageFingerprintId"`
	VersionReleased       string      `json:"gameVersionDateReleased"`
	VersionMappingID      uint        `json:"gameVersionMappingId"`
	VersionID             uint        `json:"gameVersionId"`
	GameID                uint        `json:"gameId"`
	Pack                  bool        `json:"isServerPack"`
}

// Mod is the JSON response of a mod retrieved from CurseForge.
type Mod struct {
	ID           uint            `json:"id"`
	Name         string          `json:"name"`
	Authors      []interface{}   `json:"authors"`
	Attachments  []interface{}   `json:"attachments"`
	URL          string          `json:"websiteUrl"`
	GameID       uint            `json:"gameId"`
	Summary      string          `json:"summary"`
	FileID       uint            `json:"defaultFileId"`
	Downloads    float64         `json:"downloadCount"`
	LatestFiles  []ModLatestFile `json:"latestFiles"`
	Slug         string          `json:"slug"`
	Featured     bool            `json:"isFeatured"`
	Popularity   float64         `json:"popularityScore"`
	Rank         uint            `json:"gamePopularityRank"`
	Language     string          `json:"primaryLanguage"`
	Created      string          `json:"dateCreated"`
	Updated      string          `json:"dateModified"`
	Released     string          `json:"dateReleased"`
	Available    bool            `json:"isAvailable"`
	Experimental bool            `json:"isExperimental"`
}
