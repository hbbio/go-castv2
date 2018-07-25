package controllers

import castv2 "github.com/AndreasAbdi/go-castv2"

type MediaCommand struct {
	castv2.PayloadHeaders
	MediaSessionID int `json:"mediaSessionId"`
}

type MediaStatusResponse struct {
	castv2.PayloadHeaders
	Status []*MediaStatus `json:"status,omitempty"`
}

//Generic enum type for media data

//MediaData is data format for message to send to chromecast to play a (vid/image/tvshow/music video/etc) via generic media player.
//https://developers.google.com/cast/docs/reference/messages#MediaData is the general info.
type MediaData struct {
	castv2.PayloadHeaders
	//ContentID is the identifier for the content to be loaded by the current receiver application in the chromecast.
	//Usually this is just the URL.
	//ContentType is the MIME type of the media
	ContentID   string                 `json:"contentId"`
	ContentType string                 `json:"contentType"` //data MIME
	StreamType  string                 `json:"streamType"`  // (NONE, BUFFERED, or LIVE)
	duration    float64                `json:"duration,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"` //stores a mediadata
	CustomData  map[string]interface{} `json:"customData,omitempty"`
}

type StandardMediaMetadata struct {
	MetadataType int
	Title        string
}

//TODO
type GenericMediaMetadata struct {
	StandardMediaMetadata
	Images      []string
	Subtitle    string
	ReleaseDate string
}

//TODO
type MovieMediaMetadata struct {
	StandardMediaMetadata
	Studio      string
	Subtitle    string
	Images      []string
	ReleaseDate string
}

//TODO
type TvShowMediaMetadata struct {
	StandardMediaMetadata
	Images          []string
	SeriesTitle     string
	Season          int
	Episode         int
	OriginalAirDate string
}

//TODO
type MusicTrackMediaMetadata struct {
	StandardMediaMetadata
	Images      []string
	AlbumName   string
	AlbumArtist string
	Artist      string
	Composer    string
	TrackNumber int
	DiscNumber  int
	ReleaseDate string
}

//TODO
type PhotoTrackMediaMetadata struct {
	StandardMediaMetadata
	Artist           string
	Location         string
	Latitude         float64
	Longitude        float64
	Width            int64
	Height           int64
	CreationDateTime string
}

type MediaStatus struct {
	castv2.PayloadHeaders
	MediaSessionID         int                    `json:"mediaSessionId"`
	PlaybackRate           float64                `json:"playbackRate"`
	PlayerState            string                 `json:"playerState"`
	CurrentTime            float64                `json:"currentTime"`
	SupportedMediaCommands int                    `json:"supportedMediaCommands"`
	Volume                 *Volume                `json:"volume,omitempty"`
	CustomData             map[string]interface{} `json:"customData"`
	IdleReason             string                 `json:"idleReason"`
}
