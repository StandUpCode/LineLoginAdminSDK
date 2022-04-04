package lineloginadmin

import (
	"encoding/json"
	"fmt"

	"github.com/StandUpCode/LineLoginAdminSDK/lineloginadmin/models"
	"github.com/StandUpCode/LineLoginAdminSDK/lineloginadmin/utils"
)

func VerifyAccessToken(accessToken string) error {

	uri := fmt.Sprint("https://api.line.me/oauth2/v2.1/verify?access_token=", accessToken)
	fmt.Println("VerifyAccessToken: ", uri)
	err := utils.Get(uri)

	if err != nil {
		panic(err)
	}
	return nil

}

func VerifyIDToken(IDToken , ChannelID string) (*models.IDTokenVerify, error){

	uri := "https://api.line.me/oauth2/v2.1/verify"
	payload := fmt.Sprintf("id_token=%s&client_id=%s", IDToken, ChannelID)
	
	content_type := "application/x-www-form-urlencoded"
	result, err := utils.Post(uri, payload, content_type)
	
	if err != nil {
		panic(err)
	}
	verify_response := models.IDTokenVerify{}
	err = json.Unmarshal(result, &verify_response)
	if err != nil {
		//fmt.Printf("Error: %s", err)
		return nil, err
	}
	
	return &verify_response, nil
}

