package lineloginadminsdk


type AccessTokeVerify struct {
	ClientID string `json:"client_id"`
	ExpiresIn int `json:"expires_in"`
	Scope string `json:"scope"`
}