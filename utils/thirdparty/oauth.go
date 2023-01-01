package thirdparty

import (
	"context"
	"encoding/json"
	"fmt"
	"ikuzports/features/user"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
)

func InitOauth() *oauth2.Config {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "https://rubahmerah.site/auth/callback",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/calendar.events", "https://www.googleapis.com/auth/calendar"},
		Endpoint: google.Endpoint,
	}

	return googleOauthConfig
}

func GetUserInfo(oauth *oauth2.Config, state string, code string, oauthStateString string) (user.GoogleCore, error) {

	var userGoogleCore user.GoogleCore

	if state != oauthStateString {
		log.Println("invalid oauth state")
		return user.GoogleCore{}, fmt.Errorf("invalid oauth state")
	}

	token, err := oauth.Exchange(context.Background(), code)
	if err != nil {
		log.Println("code exchange failed")
		return user.GoogleCore{}, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Println("failed getting user info")
		return user.GoogleCore{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("failed reading response body")
		return user.GoogleCore{}, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	errjson := json.Unmarshal(contents, &userGoogleCore)
	if errjson != nil {
		log.Println("cant unmarshal json")
		return user.GoogleCore{}, fmt.Errorf("cant unmarshal json")
	}

	return userGoogleCore, nil
}
