package twitterfriends

// Tweet struct used inside all the twitter.go functions
type Tweet struct {
	// Field initialize after major process or jsonUnmarshal
	TweetID uint64 `json:"id"`
	// Require field
	Status string `json:"status"`
	// Reply related fields
	InReplyTo uint64 `json:"in_reply_to_status_id"`
	UserReply string `json:"screen_name"`
	// NSFW and medias related fields
	IsSensitive bool `json:"possibly_sensitive"`
	// Position fields
	Latitude   float32  `json:"lat"`
	Longitude  float32  `json:"long"`
	ShowCoords bool `json:"display_coordinates"`
}


