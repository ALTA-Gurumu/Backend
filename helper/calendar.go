package helper

import (
	"Gurumu/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func GetClient(config *oauth2.Config) *http.Client {
	tokFile := "helper/temporary/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		SaveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func SaveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func CreateEvent(event *calendar.Event) string {
	ctx := context.Background()
	client_id := config.GOOGLE_OAUTH_CLIENT_ID1
	project := config.GOOGLE_PROJECT_ID1
	secret := config.GOOGLE_OAUTH_CLIENT_SECRET1
	b := `{"web":{"client_id":"` + client_id + `","project_id":"` + project + `","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"` + secret + `","redirect_uris":["https://devmyproject.site/callback"]}}`
	bt := []byte(b)

	config, err := google.ConfigFromJSON(bt, calendar.CalendarEventsScope, calendar.CalendarScope, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Println("unable to parse client secret file to config: ", err)
		return err.Error()
	}
	client := GetClient(config)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Println("unable to retrieve calendar client: ", err)
	}

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).SendUpdates("all").ConferenceDataVersion(1).Do()

	if err != nil {
		log.Println("unable to create event. ", err)
	}
	tautanGmet := "meet.google.com/" + event.ConferenceData.ConferenceId
	return tautanGmet
}
