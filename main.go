package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// Set the OAuth2 authorization and token exchange endpoints for Twitter (X.com)
var twitterEndpoint = oauth2.Endpoint{
	AuthURL:  "https://twitter.com/i/oauth2/authorize",
	TokenURL: "https://api.twitter.com/2/oauth2/token",
}

var config *oauth2.Config

// Step 1: Redirect user to Twitter for authorization
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Println("OAuth URL generated:", url) // Log the generated OAuth URL
	http.Redirect(w, r, url, http.StatusFound)
}

// Step 2: Handle callback and exchange the code for an access token
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// Get the authorization code from the callback URL
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Authorization code not found", http.StatusBadRequest)
		fmt.Println("Authorization code not found")
		return
	}

	// Exchange the authorization code for an access token
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		fmt.Printf("Error during token exchange: %v\n", err)
		return
	}

	// Use the access token to post a tweet
	postTweet(token)
	fmt.Fprintf(w, "Tweet posted successfully!")
}

// Function to post a tweet using the access token
func postTweet(token *oauth2.Token) {
	client := config.Client(context.Background(), token)

	// Tweet data
	tweet := map[string]string{
		"text": "Hello from X.com API using Go!",
	}
	tweetData, err := json.Marshal(tweet)
	if err != nil {
		log.Fatalf("Error marshalling tweet data: %v", err)
	}

	// Make the POST request to create the tweet
	resp, err := client.Post("https://api.twitter.com/2/tweets", "application/json", bytes.NewBuffer(tweetData))
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to post tweet, Status: %d, Error: %v", resp.StatusCode, err)
	}

	defer resp.Body.Close()
	fmt.Println("Tweet posted successfully!")
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set up the OAuth2 config for Twitter
	config = &oauth2.Config{
		ClientID:     os.Getenv("TWITTER_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITTER_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"tweet.write", "tweet.read", "users.read"}, // Removed offline.access for simplicity
		Endpoint:     twitterEndpoint,
	}

	// Set up HTTP handlers for OAuth flow
	http.HandleFunc("/", redirectHandler)         // Step 1: Start OAuth flow
	http.HandleFunc("/callback", callbackHandler) // Step 2: Handle callback

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
