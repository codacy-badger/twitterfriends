package twitterfriends

import (
	"encoding/json"
	"fmt"

	oauth "github.com/Betelgeuse1/twitteroauth"
)

// Debug Enable/Disable debugging
// If Debug is set to false
// The debug message will be printed anyway
var Debug = false

// PostTweet will send a tweet using the app.env variables
// And the json string that you'll give it.
func PostTweet(jsonInfo string) Tweet {
	tweet := createNewTweetFromJSON(jsonInfo)
	url := tweet.createPostTweetURL(false)

	client, request := tweet.createClientAndRequestForPostTweet(url)

	resp, err := client.Do(request)
	oauth.LogFatal(err)

	strResponse := oauth.GetDebugInfos(resp)
	errID := json.Unmarshal([]byte(strResponse), &tweet)
	oauth.LogFatal(errID)

	LetsDebug("TWEET POST...", tweet, request, strResponse)

	return tweet
}

// Destroy the given tweet, this is a method.
func (tweet Tweet) DestroyTweet() {
	url := fmt.Sprintf(DestroyTweetURL, tweet.TweetID)
	client, request := createSimpleClientAndRequest(url, "POST")

	resp, errResponse := client.Do(request)
	oauth.LogFatal(errResponse)

	LetsDebug("TWEET DELETED...", tweet, request, oauth.GetDebugInfos(resp))
}

// Destroy the tweet given with jsonInfo, this is a function.
func DestroyTweet(jsonInfo string) {
	tweet := createNewTweetFromJSON(jsonInfo)
	tweet.DestroyTweet()

}
