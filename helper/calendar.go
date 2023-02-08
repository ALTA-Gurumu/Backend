package helper

import (
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
	tokFile := "token.json"
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

// Saves a token to a file path.
func SaveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func Calendar(email, date, address string) (string, error) {
	event := &calendar.Event{
		Summary:     "Gurumu -  ",
		Location:    address,
		Description: "Gurumu - ",
		Start: &calendar.EventDateTime{
			Date: date,

			TimeZone: "Asia/Jakarta",
		},
		End: &calendar.EventDateTime{
			Date:     date,
			TimeZone: "Asia/Jakarta",
		},

		Attendees: []*calendar.EventAttendee{
			{Email: email},
		},
	}

	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// client_id := os.Getenv("GOOGLE_OAUTH_CLIENT_ID1")
	// project := os.Getenv("GOOGLE_PROJECT_ID1")
	// secret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET1")
	// b := `{"installed":{"client_id":"` + client_id + `","project_id":"` + project + `","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"` + secret + `","redirect_uris":["http://localhost"]}}`
	// bt := []byte(b)

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unablae to parse client secret file to config: %v", err)
	}
	client := GetClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	calendarId := "primary"
	event_notification := srv.Events.Insert(calendarId, event).SendUpdates("all")
	event, err = event_notification.Do()
	if err != nil {
		log.Fatalf("Unable to create event: %v", err)
	}

	return event.HtmlLink, nil
}

func CreateEvent(event *calendar.Event) {
	// Baca token JSON
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope, calendar.CalendarScope, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	// event := &calendar.Event{
	// 	Summary:     "Test Event",
	// 	Location:    "Somewhere",
	// 	Description: "This is a test event.",
	// 	Start: &calendar.EventDateTime{
	// 		DateTime: time.Now().Format(time.RFC3339),
	// 		TimeZone: "Asia/Jakarta",
	// 	},
	// 	End: &calendar.EventDateTime{
	// 		DateTime: time.Now().Add(time.Hour * 2).Format(time.RFC3339),
	// 		TimeZone: "Asia/Jakarta",
	// 	},
	// }

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)

}

func CreateEvent(event *calendar.Event) {
	// Baca token JSON
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope, calendar.CalendarScope, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	// event := &calendar.Event{
	// 	Summary:     "Test Event",
	// 	Location:    "Somewhere",
	// 	Description: "This is a test event.",
	// 	Start: &calendar.EventDateTime{
	// 		DateTime: time.Now().Format(time.RFC3339),
	// 		TimeZone: "Asia/Jakarta",
	// 	},
	// 	End: &calendar.EventDateTime{
	// 		DateTime: time.Now().Add(time.Hour * 2).Format(time.RFC3339),
	// 		TimeZone: "Asia/Jakarta",
	// 	},
	// }

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)

}
