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

// func getClient(config *oauth2.Config) *http.Client {
// 	// The file token.json stores the user's access and refresh tokens, and is
// 	// created automatically when the authorization flow completes for the first
// 	// time.
// 	tokFile := "token.json"
// 	tok, err := tokenFromFile(tokFile)
// 	if err != nil {
// 		tok = getTokenFromWeb(config)
// 		saveToken(tokFile, tok)
// 	}
// 	return config.Client(context.Background(), tok)
// }

// // Request a token from the web, then returns the retrieved token.
// func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	fmt.Printf("Go to the following link in your browser then type the "+
// 		"authorization code: \n%v\n", authURL)

// 	var authCode string
// 	if _, err := fmt.Scan(&authCode); err != nil {
// 		log.Fatalf("Unable to read authorization code: %v", err)
// 	}

// 	tok, err := config.Exchange(context.TODO(), authCode)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve token from web: %v", err)
// 	}
// 	return tok
// }

// // Retrieves a token from a local file.
// func tokenFromFile(file string) (*oauth2.Token, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	tok := &oauth2.Token{}
// 	err = json.NewDecoder(f).Decode(tok)
// 	return tok, err
// }

// // Saves a token to a file path.
// func saveToken(path string, token *oauth2.Token) {
// 	fmt.Printf("Saving credential file to: %s\n", path)
// 	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
// 	if err != nil {
// 		log.Fatalf("Unable to cache oauth token: %v", err)
// 	}
// 	defer f.Close()
// 	json.NewEncoder(f).Encode(token)
// }

// func InitCalendar() *calendar.Service {
// 	ctx := context.Background()
// 	b, err := os.ReadFile("credentials.json")
// 	if err != nil {
// 		log.Fatalf("Unable to read client secret file: %v", err)
// 	}

// 	// If modifying these scopes, delete your previously saved token.json.
// 	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
// 	if err != nil {
// 		log.Fatalf("Unable to parse client secret file to config: %v", err)
// 	}
// 	client := getClient(config)

// 	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve Calendar client: %v", err)
// 	}

// 	return srv
// }

// func CreateEvent(srv *calendar.Service, EventCore event.EventCore) error {
// 	event := &calendar.Event{
// 		Summary:     EventCore.Name,
// 		Location:    EventCore.Address,
// 		Description: EventCore.Description,
// 		Start: &calendar.EventDateTime{
// 			DateTime: EventCore.StartDate.String(),
// 			TimeZone: "Indonesia/Jakarta",
// 		},
// 		End: &calendar.EventDateTime{
// 			DateTime: EventCore.EndDate.String(),
// 			TimeZone: "Indonesia/Jakarta",
// 		},
// 	}

// 	calendarID := "primary"
// 	_, errCr := srv.Events.Insert(calendarID, event).Do()
// 	if errCr != nil {
// 		return errCr
// 	}

// 	return nil
// }

// func CreateEvents(oauth *oauth2.Config, code string, EventCore event.EventCore) error {
// 	token, err := oauth.Exchange(context.Background(), code)
// 	if err != nil {
// 		log.Println("code exchange failed")
// 		return fmt.Errorf("code exchange failed: %s", err.Error())
// 	}

// 	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

// 	calendarService, err := calendar.New(client)
// 	if err != nil {
// 		return err
// 	}

// 	event := &calendar.Event{
// 		Summary:     EventCore.Name,
// 		Location:    EventCore.Address,
// 		Description: EventCore.Description,
// 		Start: &calendar.EventDateTime{
// 			DateTime: EventCore.StartDate.String(),
// 			TimeZone: "Indonesia/Jakarta",
// 		},
// 		End: &calendar.EventDateTime{
// 			DateTime: EventCore.EndDate.String(),
// 			TimeZone: "Indonesia/Jakarta",
// 		},
// 	}

// 	calendarID := "primary"
// 	_, errCr := calendarService.Events.Insert(calendarID, event).Do()
// 	if errCr != nil {
// 		return errCr
// 	}

// 	return nil
// }
