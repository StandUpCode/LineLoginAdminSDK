package lineloginadminsdk

import (
	"fmt"

	lineloginadmin "github.com/StandUpCode/LineLoginAdminSDK/lineloginadmin"
)

func VerifyAccessTokenExample() {
	accessToken := "accessToken"
	verify_response, err := lineloginadmin.VerifyAccessToken(accessToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify_response: %+v\n", verify_response)
}

func VerifyIDTokenExample(){
	IDToken := "IDToken"
	channelID := "channelID"
	verify_response, err := lineloginadmin.VerifyIDToken(IDToken, channelID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify_response: %+v\n", verify_response)
}