package twitterfriends

import (
	"fmt"
	oauth "github.com/Betelgeuse1/twitteroauth"
)

// Debug Enable/Disable debugging
var Debug = false

// SendTweet will send a tweet using the app.env variables
func SendTweet(jsonTweet string) {
	tweet := createNewTweetFromJSON(jsonTweet)
	url := tweet.createURL(nil)
	fmt.Println(url)
	client, request := tweet.createClientAndRequest(url)
	fmt.Println(request, "\n")
	resp, err := client.Do(request)
	oauth.LogFatal(err)

	oauth.DebugInfos(resp)


}
