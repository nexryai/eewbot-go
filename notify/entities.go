package notify

type DiscordImg struct {
	URL string `json:"url"`
	H   int    `json:"height"`
	W   int    `json:"width"`
}

type DiscordAuthor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon_url"`
}

type DiscordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type DiscordEmbed struct {
	Title  string         `json:"title"`
	Desc   string         `json:"description"`
	URL    string         `json:"url"`
	Color  int            `json:"color"`
	Image  DiscordImg     `json:"image"`
	Author DiscordAuthor  `json:"author"`
	Fields []DiscordField `json:"fields"`
}

type DiscordHook struct {
	Username  string         `json:"username"`
	AvatarUrl string         `json:"avatar_url"`
	Content   string         `json:"content"`
	Embeds    []DiscordEmbed `json:"embeds"`
}

// Misskey
type MisskeyNote struct {
	Instance   string
	Token      string   `json:"i"`
	Text       string   `json:"text"`
	Visibility string   `json:"visibility"`
	FileIds    []string `json:"fileIds"`
}

type MisskeyDriveUploadForm struct {
	InstanceHost string
	Token        string
	Data         []byte
}

type MisskeyDriveFile struct {
	FileID string `json:"id"`
}
