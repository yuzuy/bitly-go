package bitly

type BitlinkBody struct {
	BitlinkUpdate
	HasReferences
}

type BitlinkUpdate struct {
	ID             string          `json:"id"`
	Archived       bool            `json:"archived"`
	ClientID       string          `json:"client_id"`
	CreatedAt      string          `json:"created_at"`
	CustomBitlinks []string        `json:"custom_bitlinks"`
	Deeplinks      []*DeeplinkRule `json:"deeplinks"`
	Link           string          `json:"link"`
	LongURL        string          `json:"long_url"`
	Tags           []string        `json:"tags"`
	Title          string          `json:"title"`
}

type DeeplinkRule struct {
	AppGUID     string `json:"app_guid"`
	AppURIPath  string `json:"app_uri_path"`
	Bitlink     string `json:"bitlink"`
	BrandGUID   string `json:"brand_guid"`
	Created     string `json:"created_at"`
	GUID        string `json:"guid"`
	InstallType string `json:"install_type"`
	InstallURL  string `json:"install_url"`
	Modified    string `json:"modified"`
	OS          string `json:"os"`
}

type HasReferences struct {
	References map[string]string `json:"references"`
}
