package config

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/markbates/goth"
	googleProvider "github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Verified bool   `json:"verified_email"`
	Picture  string `json:"picture"`
}

func InitProvider() {
	goth.UseProviders(googleProvider.New(Config("GOOGLE_AUTH_CLIENT_ID"), Config("GOOGLE_AUTH_SECRET"), Config("GOOGLE_AUTH_CALLBACK")))
}

func ConfigGoogle() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     Config("GOOGLE_AUTH_CLIENT_ID"),
		ClientSecret: Config("GOOGLE_AUTH_SECRET"),
		RedirectURL:  Config("GOOGLE_AUTH_CALLBACK"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}

// Get User Info of user
func GetUserInfo(token string) GoogleResponse {
	reqURL, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		panic(err)
	}
	ptoken := fmt.Sprintf("Bearer %s", token)
	res := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Authorization": {ptoken},
		},
	}
	req, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var data GoogleResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	return data
}
