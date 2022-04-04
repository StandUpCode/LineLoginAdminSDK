package lineloginadminsdk

import (
	"fmt"

	"github.com/StandUpCode/LineLoginAdminSDK/utils"
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

func VerifyIDToken(idToken string) error {
	
}