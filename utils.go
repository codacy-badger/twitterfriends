package twitterfriends

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	oauth "github.com/Betelgeuse1/twitteroauth"
)

func createNewTweetFromJSON(JSON string) Tweet {
	// Create a []byte from the JSON string
	bJSON := []byte(JSON)
	// Create an empty tweet
	tweet := Tweet{
		// If latitude (and longitude) is ommited
		// Then there value would be 0
		// But 0, 0 is a correct coordinates
		// So we are setting default values here
		// Which will be overwritten if passed
		// inside the JSON string

		// Here the default value is set to 200
		// Because Latitude goes from -90 to 90
		// And Longitude from -180 to 180
		// So 200 isn't a valid value
		Latitude:  200,
		Longitude: 200,
	}
	err := json.Unmarshal(bJSON, &tweet)
	if err != nil {
		log.Fatal(err)
		return Tweet{}

	}

	return tweet
}

func (tweet Tweet) createPostTweetURL(toEncode interface{}) string {
	if len(tweet.Status) == 0 {
		log.Fatal(errNoStatus)
		return ""

	}

	var url string

	// Check if there is an ID to reply to
	// And if so, if the screen_name is present
	if tweet.InReplyTo > 0 && tweet.UserReply != "" {
		// Add the screen_name to the status in order to make the reply works
		// Optionnal if your replying to yourself
		tweet.Status = fmt.Sprintf("%s %s", tweet.UserReply, tweet.Status)
		// If we need to encode the whole parameters inside the url
		// Then we encode the tweet.Status
		if toEncode == true {
			tweet.Status = oauth.PercentEncode(tweet.Status)

		}
		// Add the reply_ID to the url string
		url = fmt.Sprintf("%s?status=%s&in_reply_to_status_id=%d", PostTweetURL, tweet.Status, tweet.InReplyTo)

	}

	// We land there if no reply_id where specified
	// So we just add the status to the url
	if url == "" {
		if toEncode == true {
			tweet.Status = oauth.PercentEncode(tweet.Status)
		}

		url = fmt.Sprintf("%s?status=%s", PostTweetURL, tweet.Status)

	}

	if tweet.IsSensitive {
		url = fmt.Sprintf("%s&possibly_sensitive=true", url)

	}

	if tweet.Latitude != 200 && tweet.Longitude != 200 {
		url = fmt.Sprintf("%s&lat=%f&long=%f", url, tweet.Latitude, tweet.Longitude)

	}

	if tweet.ShowCoords {
		url = fmt.Sprintf("%s&display_coordinates=true", url)
	}

	return url

}

func (tweet Tweet) createClientAndRequestForPostTweet(URL string) (http.Client, *http.Request) {
	req, errCreateRequest := http.NewRequest("POST", URL, nil)
	oauth.LogFatal(errCreateRequest)

	errHeaders := oauth.SetAuthHeaders(req)
	oauth.LogFatal(errHeaders)

	encodedURL := tweet.createPostTweetURL(true)

	newURL, err := url.Parse(encodedURL)
	oauth.LogFatal(err)
	req.URL = newURL

	return http.Client{}, req
}

func createSimpleClientAndRequest(URL string, method string) (http.Client, *http.Request) {
	req, errRequest := http.NewRequest(method, URL, nil)
	oauth.LogFatal(errRequest)

	errHeaders := oauth.SetAuthHeaders(req)
	oauth.LogFatal(errHeaders)

	return http.Client{}, req

}

func LetsDebug(message string, tweet Tweet, request *http.Request, strResponse string) {
	fmt.Println(message)
	if Debug {
		fmt.Println("\n", "Tweet:", tweet)
		fmt.Println("\n", "Request:", *request)
		fmt.Println("\n", "Response:", strResponse)
		fmt.Println()

	}
}