package main

// Response is the top-level structure
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Tweet   Tweet  `json:"tweet"`
}

// Tweet represents the tweet information
type Tweet struct {
	URL               string       `json:"url"`
	ID                string       `json:"id"`
	Text              string       `json:"text"`
	Author            Author       `json:"author"`
	Replies           int          `json:"replies"`
	Retweets          int          `json:"retweets"`
	Likes             int          `json:"likes"`
	CreatedAt         string       `json:"created_at"`
	CreatedTimestamp  int64        `json:"created_timestamp"`
	PossiblySensitive bool         `json:"possibly_sensitive"`
	Views             int          `json:"views"`
	IsNoteTweet       bool         `json:"is_note_tweet"`
	Lang              string       `json:"lang"`
	ReplyingTo        *string      `json:"replying_to"`
	ReplyingToStatus  *interface{} `json:"replying_to_status"`
	Media             Media        `json:"media"`
	Quote             *Quote       `json:"quote"`
	Source            string       `json:"source"`
	TwitterCard       string       `json:"twitter_card"`
	Color             *interface{} `json:"color"`
}

// Author represents the author of the tweet
type Author struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	ScreenName  string       `json:"screen_name"`
	AvatarURL   string       `json:"avatar_url"`
	BannerURL   string       `json:"banner_url"`
	Description string       `json:"description"`
	Location    string       `json:"location"`
	URL         string       `json:"url"`
	Followers   int          `json:"followers"`
	Following   int          `json:"following"`
	Joined      string       `json:"joined"`
	Likes       int          `json:"likes"`
	Website     Website      `json:"website"`
	Tweets      int          `json:"tweets"`
	AvatarColor *interface{} `json:"avatar_color"`
}

// Website represents the website information
type Website struct {
	URL        string `json:"url"`
	DisplayURL string `json:"display_url"`
}

// Media represents media content in the tweet
type Media struct {
	All    []MediaContent `json:"all"`
	Photos []Photo        `json:"photos"`
	Videos []Video        `json:"videos"`
	Mosaic Mosaic         `json:"mosaic"`
}

// MediaContent represents each media content
type MediaContent struct {
	Type         string  `json:"type"`
	URL          string  `json:"url"`
	Width        int     `json:"width"`
	Height       int     `json:"height"`
	AltText      string  `json:"altText,omitempty"`
	ThumbnailURL string  `json:"thumbnail_url,omitempty"`
	Duration     float64 `json:"duration,omitempty"`
	Format       string  `json:"format,omitempty"`
}

type Photo struct {
	Type    string `json:"type"`
	URL     string `json:"url"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	AltText string `json:"altText"`
}

type Video struct {
	URL          string  `json:"url"`
	ThumbnailURL string  `json:"thumbnail_url"`
	Duration     float64 `json:"duration"`
	Width        int     `json:"width"`
	Height       int     `json:"height"`
	Format       string  `json:"format"`
	Type         string  `json:"type"`
}

// Mosaic represents the mosaic view of media
type Mosaic struct {
	Type    string        `json:"type"`
	Formats MosaicFormats `json:"formats"`
}

// MosaicFormats represents the formats of the mosaic
type MosaicFormats struct {
	JPEG string `json:"jpeg"`
	WEBP string `json:"webp"`
}

// Quote represents the quoted tweet
type Quote struct {
	URL              string       `json:"url"`
	ID               string       `json:"id"`
	Text             string       `json:"text"`
	Author           Author       `json:"author"`
	Replies          int          `json:"replies"`
	Retweets         int          `json:"retweets"`
	Likes            int          `json:"likes"`
	CreatedAt        string       `json:"created_at"`
	CreatedTimestamp int64        `json:"created_timestamp"`
	Views            int          `json:"views"`
	IsNoteTweet      bool         `json:"is_note_tweet"`
	Lang             string       `json:"lang"`
	ReplyingTo       string       `json:"replying_to"`
	ReplyingToStatus *interface{} `json:"replying_to_status"`
	Media            Media        `json:"media"`
	Source           string       `json:"source"`
	TwitterCard      string       `json:"twitter_card"`
	Color            *interface{} `json:"color"`
}
