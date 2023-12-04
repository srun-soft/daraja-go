package darajago

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type authResponse struct {
	AccessToken string `json:"access_token"` // The access token to be used in subsequent API calls
	ExpiresIn   string `json:"expires_in"`   // The number of seconds before the access token expires
}

type Authorization struct {
	authResponse
}

func newAuthorization(consumerKey, consumerSecret string, env Environment) (*Authorization, error) {
	authHeader := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(consumerKey+":"+consumerSecret)),
	}

	netPackage := newRequestPackage(nil, endpointAuth, http.MethodGet, authHeader, env)
	authResponse, err := newRequest(netPackage)
	if err != nil {
		return nil, err
	}

	// 处理成功的情况，通过类型断言获取具体的 Body
	res, ok := authResponse.Body.(Authorization)
	if !ok {
		// 类型断言失败，处理错误
		fmt.Println("Error: Unable to assert type")
	} else {
		// 成功获取 Authorization
		fmt.Println("Authorization:", res)
	}
	return &res, nil
}
