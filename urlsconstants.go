package twitterfriends

const (
	// All twitter api's url
	// PostTweetURL is used for Posting tweet (tweet, reply)
	PostTweetURL = "https://api.twitter.com/1.1/statuses/update.json"
	// DestroyTweetURL is used in order to delete a tweet
	DestroyTweetURL = "https://api.twitter.com/1.1/statuses/destroy/%d.json"
	// ShowSingleTweetURL is used in order to get a single tweet from an id
	ShowSingleTweetURL = "https://api.twitter.com/1.1/statuses/show.json?id=%d"
)
