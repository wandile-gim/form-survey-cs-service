package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
)

const confDir = "internal/config/"

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	path := fmt.Sprintf("%s%s", confDir, tokFile)
	tok, err := tokenFromFile(path)
	if err != nil {
		// 토큰이 없으면 웹에서 인증 코드를 받아야 함
		tok = getTokenFromWeb(config)
		saveToken(path, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser and complete the authentication:\n%v\n", authURL)

	// 사용자로부터 인증 코드를 입력받음
	var authCode string
	fmt.Print("Enter the authorization code: ")
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	// 입력받은 인증 코드를 사용해 토큰을 교환
	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
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
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func OAuthForSheet() *http.Client {
	credential := "client_secret.json"
	b, err := os.ReadFile(fmt.Sprintf("%s%s", confDir, credential))
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// OAuth2 설정 로드
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return getClient(config)
}
