package twitterfriends

import (
	"log"
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"

	oauth "github.com/Betelgeuse1/twitteroauth"
)

const (
	// All Twitters' URL
	sendTweetURL = "https://api.twitter.com/1.1/statuses/update.json"

)

// Tweet struct used inside all the twitter.go functions
type Tweet struct {
	Status  string `json:"status"`
	InReplyTo uint64 `json:"in_reply_to_status_id"`
	UserReply string `json:"screen_name"`
	IsSensitive bool `json:"possibly_sensitive"`

}

func createNewTweetFromJSON(JSON string) *Tweet {
	// Create a []byte from the JSON string
	bJSON := []byte(JSON)
	// Create an empty tweet
	tweet := &Tweet{}
	err := json.Unmarshal(bJSON, &tweet)
	if err != nil {
		log.Fatal(err)
		return nil

	}

	return tweet
}

func (tweet *Tweet) createURL(toEncode interface{}) (string) {
	if len(tweet.Status) == 0 {
		log.Fatal(errNoStatus)
		return ""

	}

	if toEncode == true {
		tweet.Status = oauth.PercentEncode(tweet.Status)
	}

	var url string

	if tweet.InReplyTo > 0 && tweet.UserReply != "" {
		url = fmt.Sprintf("%s?status=%%40%s%%20%s&in_reply_to_status_id=%v", sendTweetURL, tweet.UserReply, tweet.Status, tweet.InReplyTo)
	
	} else {
		url = fmt.Sprintf("%s?status=%s", sendTweetURL, tweet.Status)

	}

	if tweet.IsSensitive {
		url = fmt.Sprintf("%s&possibly_sensitive=true", url)

	}

	return url

}

func (tweet *Tweet) createClientAndRequest(URL string) (http.Client, *http.Request) {
	req, errCreateRequest := http.NewRequest("POST", URL, nil)
	oauth.LogFatal(errCreateRequest)

	errHeaders := oauth.SetAuthHeaders(req)
	oauth.LogFatal(errHeaders)

	encodedURL := tweet.createURL(true)

	newURL, err := url.Parse(encodedURL)
	oauth.LogFatal(err)
	req.URL = newURL

	return http.Client{}, req
}

