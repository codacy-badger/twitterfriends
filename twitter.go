package twitterfriends

import (
	// Golang packages
	"fmt"
	"net/http"
	"net/url"

	// Own package
	oauth "github.com/Betelgeuse1/twitteroauth"
)

// Enable/Disable debugging
var Debug bool = false

// All Twitters' URL
const (
	sendTweetURL = "https://api.twitter.com/1.1/statuses/update.json"
)

// Send a tweet with the application codes inside app.env file
func SendTweet(status string) {
	// Create a Client
	client := &http.Client{}
	// Add the status to the url but not encoded
	// Because if you send: "TEST TEST" and we encoded the status here
	// The signature will used %2520 as %20 which is the space
	// And the same thing goes for every special character that will be encoded
	joinURL := fmt.Sprintf("%s?status=%s", sendTweetURL, status)

	// Create the request and handle errors
	req, err := http.NewRequest("POST", joinURL, nil)
	// Handle err in a cleaner way
	oauth.LogFatal(err)

	// Use oauth package in order to add the
	// Authorization Header to the request
	headerErr := oauth.SetAuthHeaders(req)
	oauth.LogFatal(headerErr)

	// Here we create the encoded URL
	encodeURL := fmt.Sprintf("%s?status=%s", sendTweetURL, oauth.PercentEncode(status))
	// Create a new url.URL object to replace the one in req (our request)
	newURL, urlErr := url.Parse(encodeURL)
	oauth.LogFatal(urlErr)

	// Change the non-encoded url with the encoded one
	req.URL = newURL

	// Make the requets
	resp, reqErr := client.Do(req)
	oauth.LogFatal(reqErr)

	// If Debug is set, show twitter's response
	if Debug == true {
		fmt.Println(req)
		fmt.Println()
		oauth.DebugInfos(resp)

	}

}
