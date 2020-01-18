package schema

type RefreshTokenSchema struct {
	UserId       string   `json:"user_id"`
	ClientId     string   `json:"client_id"`
	RefreshToken string   `json:"refresh_token"`
	Scopes       []string `json:"scopes"`
}
