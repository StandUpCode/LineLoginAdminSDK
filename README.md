# Line Login Admin SDK

![alt text](LineLoginlogo.png "Line Login Logo")

## Line Login Admin SDK is a library for [Line Login](https://developers.line.biz/en/docs/line-login/).

We provide a simple and easy to use SDK for developers to integrate Line Login with Go Backend.

## 🎯 Featues:

- `Verify AccesToken`
- `Verify ID Token`

---

## Installation

```hell
    go get -u github.com/StandUpCode/LineLoginAdminSDK
```

---

## Quick Start Guide
### Verify Access Token
```go
func VerifyAccessTokenExample() {
	accessToken := "accessToken"
	verify_response, err := lineloginadmin.VerifyAccessToken(accessToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify_response: %+v\n", verify_response)
}


```
### Verify ID Token

```go

func VerifyIDTokenExample(){
	IDToken := "IDToken"
	channelID := "channelID"
	verify_response, err := lineloginadmin.VerifyIDToken(IDToken, channelID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("verify_response: %+v\n", verify_response)
}
```

---

## ⚠️ License

Copyright (c) 2022-present Pattanan and Contributors. LineLoginAdminSDK is free and open-source software licensed under the MIT License.
distributed under Creative Commons license (CC BY-SA 4.0 International).

Third-party library licenses

- [fasthttp](github.com/valyala/fasthttp)
