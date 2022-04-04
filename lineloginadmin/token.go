package lineloginadmin

import (
	"encoding/json"
	"fmt"

	"github.com/StandUpCode/LineLoginAdminSDK/lineloginadmin/utils"
)

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
	Email   string   `json:"email"`
}

func VerifyAccessToken(accessToken string) (*AccessTokenVerify, error) {

	uri := fmt.Sprint("https://api.line.me/oauth2/v2.1/verify?access_token=", accessToken)
	fmt.Println("VerifyAccessToken: ", uri)
	result, err := utils.Get(uri)
	verify_response := AccessTokenVerify{}

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(result, &verify_response)
	if err != nil {
		//fmt.Printf("Error: %s", err)
		return nil, err
	}
	return &verify_response, nil

}

func VerifyIDToken(IDToken, ChannelID string) (*IDTokenVerify, error) {

	uri := "https://api.line.me/oauth2/v2.1/verify"
	payload := fmt.Sprintf("id_token=%s&client_id=%s", IDToken, ChannelID)

	content_type := "application/x-www-form-urlencoded"
	result, err := utils.Post(uri, payload, content_type)

	if err != nil {
		panic(err)
	}
	verify_response := IDTokenVerify{}
	err = json.Unmarshal(result, &verify_response)
	if err != nil {
		//fmt.Printf("Error: %s", err)
		return nil, err
	}

	return &verify_response, nil
}
