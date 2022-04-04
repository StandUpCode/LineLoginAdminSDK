package lineloginadminsdk

type AccessTokenVerify struct {
	ClientID  string `json:"client_id"`
	ExpiresIn int    `json:"expires_in"`
	Scope     string `json:"scope"`
}
type IDTokenVerify struct {
	ISS     string   `json:"iss"`
	Aud     string   `json:"aud"`
	Exp     int      `json:"exp"`
	Iat     int      `json:"iat"`
	Sub     string   `json:"sub"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
	Amr     []string `json:"amr"`
	email string `json:"email"`
}
