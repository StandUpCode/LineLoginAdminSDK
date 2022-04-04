package lineloginadminsdk

import (
	"testing"
)

var AccessTokenTestData = []struct {
	accessToken string
}{
	{"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoiVWM3YTc0MTUyYjI5MDAzMDkwNGJiNWM1ZGExM2NiNDRmIiwiYXVkIjoiMTY1NjcxMTg0NCIsImV4cCI6MTY0OTA1NTY3OSwiaWF0IjoxNjQ5MDUyMDc5LCJhbXIiOlsibGluZXNzbyJdLCJuYW1lIjoiQVJNIiwicGljdHVyZSI6Imh0dHBzOi8vcHJvZmlsZS5saW5lLXNjZG4ubmV0LzBoQldfb0tuODBIWGhpRmdsWmhCVmlMMTVURXhVVk9Cc3dHblFDR2tFVlJFMUpJQWtwV25SYVMwQVJFeDhjY1EwbURTTlJIQlVlUzBrZiJ9.KJpi1tu4nbOagjiDJhau5rnm4c_I55eB1PcNh9FALvY"},
}

var IDTestData = []struct {
	IDToken   string
	ChannelID string
}{
	{
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoiVWM3YTc0MTUyYjI5MDAzMDkwNGJiNWM1ZGExM2NiNDRmIiwiYXVkIjoiMTY1NjcxMTg0NCIsImV4cCI6MTY0OTA2MjIyNiwiaWF0IjoxNjQ5MDU4NjI2LCJhbXIiOlsibGluZXNzbyJdLCJuYW1lIjoiQVJNIiwicGljdHVyZSI6Imh0dHBzOi8vcHJvZmlsZS5saW5lLXNjZG4ubmV0LzBoQldfb0tuODBIWGhpRmdsWmhCVmlMMTVURXhVVk9Cc3dHblFDR2tFVlJFMUpJQWtwV25SYVMwQVJFeDhjY1EwbURTTlJIQlVlUzBrZiJ9.a36c5JW0fzDx5RgYECsxLKtBoqgc7q93nZ_RrHmzeqo",
		"1656711844",
	},
}

func TestVerifyAccessToken(t *testing.T) {
	for _, data := range AccessTokenTestData {
		var err error

		if err = VerifyAccessToken(data.accessToken); err != nil {
			t.Errorf("Unable to verify accesstoken: %v", err)
		}

	}

}

func TestVerifyIDToken(t *testing.T) {
	for _, data := range IDTestData {
		var err error

		if _,err = VerifyIDToken(data.IDToken, data.ChannelID); err != nil {
			t.Errorf("Unable to verify accesstoken: %v", err)
		}

	}

}
