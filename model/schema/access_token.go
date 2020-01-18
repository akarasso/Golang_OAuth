package schema

type AccessTokenSchema struct {
	UserId      string   `json:"user_id"`
	ClientId    string   `json:"client_id"`
	AccessToken string   `json:"access_token"`
	Scopes      []string `json:"scopes"`
	TimeToLive  int      `json:"ttl"`
}
