package darajago

import (
	"encoding/base64"
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
	var auth Authorization
	authHeader := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(consumerKey+":"+consumerSecret)),
	}

	netPackage := newRequestPackage(nil, endpointAuth, http.MethodGet, authHeader, env)
	authResponse, err := newRequest(netPackage)
	if err != nil {
		return nil, err
	}
	if err := authResponse.Decode(&auth); err != nil {
		return nil, err
	}
	return &auth, nil
}
