package schema

type AuthorizationCodeSchema struct {
	UserId            string   `json:"user_id"`
	ClientId          string   `json:"client_id"`
	AuthorizationCode string   `json:"authorization_code"`
	RedirectUri       string   `json:"redirect_uri"`
	Scopes            []string `json:"scopes"`
	TimeToLive        int      `json:"ttl"`
}
