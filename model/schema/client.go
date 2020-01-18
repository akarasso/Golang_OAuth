package schema

type ClientSchema struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	UserId       string   `json:"user_id"`
	RedirectUri  string   `json:"redirect_uri"`
	Scopes       []string `json:"scopes"`
	GrantTypes   []string `json:"grant_types"`
}
