package webhook

type Webhook struct {
	CallbackUrl string     `json:"callback_url"`
	PushData    PushData   `json:"push_data"`
	Repository  Repository `json:"repository"`
}

type PushData struct {
	PushedAt uint64   `json:"pushed_at"`
	Images   []string `json:"images"`
	Tag      string   `json:"tag"`
	Pusher   string   `json:"pusher"`
}

type Repository struct {
	Status          string `json:"status"`
	Description     string `json:"description"`
	IsTrusted       bool   `json:"is_trusted"`
	FullDescription string `json:"full_description"`
	RepoUrl         string `json:"repo_url"`
	Owner           string `json:"owner"`
	IsOfficial      bool   `json:"is_official"`
	IsPrivate       bool   `json:"is_private"`
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	StarCount       uint64 `json:"star_count"`
	CommentCount    uint64 `json:"comment_count"`
	DateCreated     uint64 `json:"date_created"`
	RepoName        string `json:"repo_name"`
}
