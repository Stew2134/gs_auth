package main

import(
	"context"
	"fmt"
	"log"
	"os"
	"golang.org/x/oauth2/google"
)

func main() {
	if len(os.Args)<2 {
		log.Fatal("Usage: gs_auth ./application_credentials.json")
	}
	filename := os.Args[1]
	ctx := context.Background()
	jsonKey, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read service account file: %v", err)
	}
	creds, err := google.CredentialsFromJSON(ctx, jsonKey, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Failed to parse credentials from JSON: %v", err)
	}
	ts := creds.TokenSource
	token, err := ts.Token()
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}
	fmt.Println(token.AccessToken)
}
