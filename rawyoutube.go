// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package influshipapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Influship/influship-go/internal/apijson"
	"github.com/Influship/influship-go/internal/apiquery"
	"github.com/Influship/influship-go/internal/requestconfig"
	"github.com/Influship/influship-go/option"
	"github.com/Influship/influship-go/packages/param"
	"github.com/Influship/influship-go/packages/respjson"
	"github.com/Influship/influship-go/shared/constant"
)

// Fetch fresh data directly from social platforms in real-time. Use when you need
// the most current information or data for profiles not yet in our database.
//
// RawYoutubeService contains methods and other services that help with interacting
// with the Influship API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRawYoutubeService] method instead.
type RawYoutubeService struct {
	options []option.RequestOption
}

// NewRawYoutubeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRawYoutubeService(opts ...option.RequestOption) (r RawYoutubeService) {
	r = RawYoutubeService{}
	r.options = opts
	return
}

// Fetch fresh YouTube channel data including subscriber count, video count, and
// total views.
//
// **Pricing**: 0.5 credits per channel scraped ($0.005)
func (r *RawYoutubeService) GetChannel(ctx context.Context, handle string, query RawYoutubeGetChannelParams, opts ...option.RequestOption) (res *RawYoutubeGetChannelResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if handle == "" {
		err = errors.New("missing required handle parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/raw/youtube/channel/%s", url.PathEscape(handle))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch transcripts for multiple videos from a YouTube channel. Videos can be
// sorted by popularity, newest, or oldest before selection.
//
// **Features:**
//
// - Fetches up to 20 video transcripts per request
// - Sort by popular (most views), newest, or oldest
// - Partial success — individual video failures don't block the response
// - Optional timestamped segments for each transcript
//
// **Pricing**: 0.5 credits per transcript fetched ($0.005)
func (r *RawYoutubeService) GetChannelTranscripts(ctx context.Context, handle string, query RawYoutubeGetChannelTranscriptsParams, opts ...option.RequestOption) (res *RawYoutubeGetChannelTranscriptsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if handle == "" {
		err = errors.New("missing required handle parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/raw/youtube/channel-transcripts/%s", url.PathEscape(handle))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch YouTube video transcript/captions. Returns timestamped segments and full
// text. Useful for content analysis.
//
// **Supported sources:**
//
// - Manual captions (highest quality)
// - Auto-generated captions
// - Multiple language tracks
//
// **Pricing**: 0.5 credits per transcript ($0.005)
func (r *RawYoutubeService) GetTranscript(ctx context.Context, videoID string, query RawYoutubeGetTranscriptParams, opts ...option.RequestOption) (res *RawYoutubeGetTranscriptResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if videoID == "" {
		err = errors.New("missing required video_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/raw/youtube/transcript/%s", url.PathEscape(videoID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search YouTube videos and channels.
//
// **Pricing**: 0.5 credits per result returned ($0.005)
func (r *RawYoutubeService) Search(ctx context.Context, query RawYoutubeSearchParams, opts ...option.RequestOption) (res *RawYoutubeSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/raw/youtube/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type TranscriptSegment struct {
	// Duration in seconds
	Duration float64 `json:"duration" api:"required"`
	// Start time in seconds
	Start float64 `json:"start" api:"required"`
	// Segment text
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Start       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptSegment) RawJSON() string { return r.JSON.raw }
func (r *TranscriptSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeGetChannelResponse struct {
	Data RawYoutubeGetChannelResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeGetChannelResponse) RawJSON() string { return r.JSON.raw }
func (r *RawYoutubeGetChannelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeGetChannelResponseData struct {
	// Channel avatar URL
	AvatarURL string `json:"avatar_url" api:"required" format:"uri"`
	// Channel description
	Description string `json:"description" api:"required"`
	// Channel handle
	Handle string `json:"handle" api:"required"`
	// Channel name
	Name string `json:"name" api:"required"`
	// When this data was scraped
	ScrapedAt time.Time `json:"scraped_at" api:"required" format:"date-time"`
	// Subscriber count
	Subscribers int64 `json:"subscribers" api:"required"`
	// Total video count
	VideosCount int64 `json:"videos_count" api:"required"`
	// Total view count
	ViewsTotal int64 `json:"views_total" api:"required"`
	// Recent videos (only included when include_videos=true)
	Videos []RawYoutubeGetChannelResponseDataVideo `json:"videos"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvatarURL   respjson.Field
		Description respjson.Field
		Handle      respjson.Field
		Name        respjson.Field
		ScrapedAt   respjson.Field
		Subscribers respjson.Field
		VideosCount respjson.Field
		ViewsTotal  respjson.Field
		Videos      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeGetChannelResponseData) RawJSON() string { return r.JSON.raw }
func (r *RawYoutubeGetChannelResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeGetChannelResponseDataVideo struct {
	// Video ID
	ID string `json:"id" api:"required"`
	// Comment count
	Comments int64 `json:"comments" api:"required"`
	// Video duration in seconds
	DurationSeconds int64 `json:"duration_seconds" api:"required"`
	// Like count
	Likes int64 `json:"likes" api:"required"`
	// Publish timestamp
	PublishedAt time.Time `json:"published_at" api:"required" format:"date-time"`
	// Thumbnail URL
	ThumbnailURL string `json:"thumbnail_url" api:"required" format:"uri"`
	// Video title
	Title string `json:"title" api:"required"`
	// Video URL
	URL string `json:"url" api:"required" format:"uri"`
	// View count
	Views int64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Comments        respjson.Field
		DurationSeconds respjson.Field
		Likes           respjson.Field
		PublishedAt     respjson.Field
		ThumbnailURL    respjson.Field
		Title           respjson.Field
		URL             respjson.Field
		Views           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeGetChannelResponseDataVideo) RawJSON() string { return r.JSON.raw }
func (r *RawYoutubeGetChannelResponseDataVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeGetChannelTranscriptsResponse struct {
	Data RawYoutubeGetChannelTranscriptsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeGetChannelTranscriptsResponse) RawJSON() string { return r.JSON.raw }
func (r *RawYoutubeGetChannelTranscriptsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeGetChannelTranscriptsResponseData struct {
	// YouTube channel ID
	ChannelID string `json:"channel_id" api:"required"`
	// Channel name
	ChannelName string `json:"channel_name" api:"required"`
	// Channel handle
	Handle string `json:"handle" api:"required"`
	// Per-video transcript results
	Items []RawYoutubeGetChannelTranscriptsResponseDataItem `json:"items" api:"required"`
	// When this data was scraped
	ScrapedAt string `json:"scraped_at" api:"required"`
	// Number of transcripts that failed to fetch
	TranscriptsFailed int64 `json:"transcripts_failed" api:"required"`
	// Number of transcripts successfully fetched
	TranscriptsFetched int64 `json:"transcripts_fetched" api:"required"`
	// Total videos found on channel
	VideosFound int64 `json:"videos_found" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelID          respjson.Field
		ChannelName        respjson.Field
		Handle             respjson.Field
		Items              respjson.Field
		ScrapedAt          respjson.Field
		TranscriptsFailed  respjson.Field
		TranscriptsFetched respjson.Field
		VideosFound        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeGetChannelTranscriptsResponseData) RawJSON() string { return r.JSON.raw }
func (r *RawYoutubeGetChannelTranscriptsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeGetChannelTranscriptsResponseDataItem struct {
	// Error message if transcript fetch failed for this video
	Error string `json:"error" api:"required"`
	// Full transcript as plain text
	FullText string `json:"full_text" api:"required"`
	// Transcript language code
	Language string `json:"language" api:"required"`
	// Relative publish time
	PublishedText string `json:"published_text" api:"required"`
	// Caption source type
	//
	// Any of "manual", "auto_generated".
	Source string `json:"source" api:"required"`
	// Video title
	Title string `json:"title" api:"required"`
	// Timestamped segments (only if include_segments=true)
	Transcript []TranscriptSegment `json:"transcript" api:"required"`
	// Video URL
	URL string `json:"url" api:"required"`
	// YouTube video ID
	VideoID string `json:"video_id" api:"required"`
	// View count
	ViewCount float64 `json:"view_count" api:"required"`
	// Word count of transcript
	WordCount float64 `json:"word_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error         respjson.Field
		FullText      respjson.Field
		Language      respjson.Field
		PublishedText respjson.Field
		Source        respjson.Field
		Title         respjson.Field
		Transcript    respjson.Field
		URL           respjson.Field
		VideoID       respjson.Field
		ViewCount     respjson.Field
		WordCount     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeGetChannelTranscriptsResponseDataItem) RawJSON() string { return r.JSON.raw }
func (r *RawYoutubeGetChannelTranscriptsResponseDataItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeGetTranscriptResponse struct {
	Data RawYoutubeGetTranscriptResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeGetTranscriptResponse) RawJSON() string { return r.JSON.raw }
func (r *RawYoutubeGetTranscriptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeGetTranscriptResponseData struct {
	// Available transcript languages
	AvailableLanguages []string `json:"available_languages" api:"required"`
	// Full transcript as plain text
	FullText string `json:"full_text" api:"required"`
	// Transcript language code
	Language string `json:"language" api:"required"`
	// Video title
	Title string `json:"title" api:"required"`
	// Transcript segments
	Transcript []TranscriptSegment `json:"transcript" api:"required"`
	// Video URL
	URL string `json:"url" api:"required" format:"uri"`
	// Video ID
	VideoID string `json:"video_id" api:"required"`
	// Total word count
	WordCount int64 `json:"word_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableLanguages respjson.Field
		FullText           respjson.Field
		Language           respjson.Field
		Title              respjson.Field
		Transcript         respjson.Field
		URL                respjson.Field
		VideoID            respjson.Field
		WordCount          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeGetTranscriptResponseData) RawJSON() string { return r.JSON.raw }
func (r *RawYoutubeGetTranscriptResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeSearchResponse struct {
	Data RawYoutubeSearchResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *RawYoutubeSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeSearchResponseData struct {
	// Estimated total results count
	EstimatedResults float64 `json:"estimated_results" api:"required"`
	// The search query
	Query string `json:"query" api:"required"`
	// Search results (videos and channels)
	Results []RawYoutubeSearchResponseDataResultUnion `json:"results" api:"required"`
	// When this search was performed
	ScrapedAt string `json:"scraped_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EstimatedResults respjson.Field
		Query            respjson.Field
		Results          respjson.Field
		ScrapedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeSearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *RawYoutubeSearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RawYoutubeSearchResponseDataResultUnion contains all possible properties and
// values from [RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult],
// [RawYoutubeSearchResponseDataResultYouTubeSearchChannelResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type RawYoutubeSearchResponseDataResultUnion struct {
	ChannelHandle   string `json:"channel_handle"`
	ChannelID       string `json:"channel_id"`
	ChannelName     string `json:"channel_name"`
	ChannelVerified bool   `json:"channel_verified"`
	Description     string `json:"description"`
	// This field is from variant
	// [RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult].
	DurationSeconds float64 `json:"duration_seconds"`
	// This field is from variant
	// [RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult].
	DurationText string `json:"duration_text"`
	// This field is from variant
	// [RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult].
	PublishedAt string `json:"published_at"`
	// This field is from variant
	// [RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult].
	PublishedText string `json:"published_text"`
	ThumbnailURL  string `json:"thumbnail_url"`
	// This field is from variant
	// [RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult].
	Title string `json:"title"`
	Type  string `json:"type"`
	URL   string `json:"url"`
	// This field is from variant
	// [RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult].
	VideoID string `json:"video_id"`
	// This field is from variant
	// [RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult].
	ViewCount float64 `json:"view_count"`
	// This field is from variant
	// [RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult].
	ViewCountText string `json:"view_count_text"`
	// This field is from variant
	// [RawYoutubeSearchResponseDataResultYouTubeSearchChannelResult].
	SubscriberCount float64 `json:"subscriber_count"`
	// This field is from variant
	// [RawYoutubeSearchResponseDataResultYouTubeSearchChannelResult].
	VideoCount int64 `json:"video_count"`
	JSON       struct {
		ChannelHandle   respjson.Field
		ChannelID       respjson.Field
		ChannelName     respjson.Field
		ChannelVerified respjson.Field
		Description     respjson.Field
		DurationSeconds respjson.Field
		DurationText    respjson.Field
		PublishedAt     respjson.Field
		PublishedText   respjson.Field
		ThumbnailURL    respjson.Field
		Title           respjson.Field
		Type            respjson.Field
		URL             respjson.Field
		VideoID         respjson.Field
		ViewCount       respjson.Field
		ViewCountText   respjson.Field
		SubscriberCount respjson.Field
		VideoCount      respjson.Field
		raw             string
	} `json:"-"`
}

func (u RawYoutubeSearchResponseDataResultUnion) AsRawYoutubeSearchResponseDataResultYouTubeSearchVideoResult() (v RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RawYoutubeSearchResponseDataResultUnion) AsRawYoutubeSearchResponseDataResultYouTubeSearchChannelResult() (v RawYoutubeSearchResponseDataResultYouTubeSearchChannelResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RawYoutubeSearchResponseDataResultUnion) RawJSON() string { return u.JSON.raw }

func (r *RawYoutubeSearchResponseDataResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult struct {
	// Channel handle (e.g., @username)
	ChannelHandle string `json:"channel_handle" api:"required"`
	// Channel ID
	ChannelID string `json:"channel_id" api:"required"`
	// Channel name
	ChannelName string `json:"channel_name" api:"required"`
	// Whether channel is verified
	ChannelVerified bool `json:"channel_verified" api:"required"`
	// Video description snippet
	Description string `json:"description" api:"required"`
	// Video duration in seconds
	DurationSeconds float64 `json:"duration_seconds" api:"required"`
	// Video duration text (e.g., "12:34")
	DurationText string `json:"duration_text" api:"required"`
	// Approximate publish timestamp
	PublishedAt string `json:"published_at" api:"required"`
	// Relative publish time (e.g., "2 days ago")
	PublishedText string `json:"published_text" api:"required"`
	// Video thumbnail URL
	ThumbnailURL string `json:"thumbnail_url" api:"required"`
	// Video title
	Title string         `json:"title" api:"required"`
	Type  constant.Video `json:"type" api:"required"`
	// Full YouTube video URL
	URL string `json:"url" api:"required"`
	// YouTube video ID
	VideoID string `json:"video_id" api:"required"`
	// Number of views
	ViewCount float64 `json:"view_count" api:"required"`
	// View count text (e.g., "1.2M views")
	ViewCountText string `json:"view_count_text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelHandle   respjson.Field
		ChannelID       respjson.Field
		ChannelName     respjson.Field
		ChannelVerified respjson.Field
		Description     respjson.Field
		DurationSeconds respjson.Field
		DurationText    respjson.Field
		PublishedAt     respjson.Field
		PublishedText   respjson.Field
		ThumbnailURL    respjson.Field
		Title           respjson.Field
		Type            respjson.Field
		URL             respjson.Field
		VideoID         respjson.Field
		ViewCount       respjson.Field
		ViewCountText   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult) RawJSON() string {
	return r.JSON.raw
}
func (r *RawYoutubeSearchResponseDataResultYouTubeSearchVideoResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeSearchResponseDataResultYouTubeSearchChannelResult struct {
	// Channel handle (e.g., @username)
	ChannelHandle string `json:"channel_handle" api:"required"`
	// Channel ID
	ChannelID string `json:"channel_id" api:"required"`
	// Channel name
	ChannelName string `json:"channel_name" api:"required"`
	// Whether channel is verified
	ChannelVerified bool `json:"channel_verified" api:"required"`
	// Channel description snippet
	Description string `json:"description" api:"required"`
	// Subscriber count
	SubscriberCount float64 `json:"subscriber_count" api:"required"`
	// Channel avatar URL
	ThumbnailURL string           `json:"thumbnail_url" api:"required"`
	Type         constant.Channel `json:"type" api:"required"`
	// Full YouTube channel URL
	URL string `json:"url" api:"required"`
	// Number of videos on the channel
	VideoCount int64 `json:"video_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelHandle   respjson.Field
		ChannelID       respjson.Field
		ChannelName     respjson.Field
		ChannelVerified respjson.Field
		Description     respjson.Field
		SubscriberCount respjson.Field
		ThumbnailURL    respjson.Field
		Type            respjson.Field
		URL             respjson.Field
		VideoCount      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawYoutubeSearchResponseDataResultYouTubeSearchChannelResult) RawJSON() string {
	return r.JSON.raw
}
func (r *RawYoutubeSearchResponseDataResultYouTubeSearchChannelResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawYoutubeGetChannelParams struct {
	// Include recent videos in response
	IncludeVideos param.Opt[bool] `query:"include_videos,omitzero" json:"-"`
	// Number of videos to include
	VideoLimit param.Opt[int64] `query:"video_limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RawYoutubeGetChannelParams]'s query parameters as
// `url.Values`.
func (r RawYoutubeGetChannelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RawYoutubeGetChannelTranscriptsParams struct {
	// Include timestamped transcript segments in response
	IncludeSegments param.Opt[bool] `query:"include_segments,omitzero" json:"-"`
	// Language code for transcripts
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Number of videos to fetch transcripts for (max 20)
	VideoLimit param.Opt[int64] `query:"video_limit,omitzero" json:"-"`
	// How to sort channel videos before selecting
	//
	// Any of "popular", "newest", "oldest".
	SortBy RawYoutubeGetChannelTranscriptsParamsSortBy `query:"sort_by,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RawYoutubeGetChannelTranscriptsParams]'s query parameters
// as `url.Values`.
func (r RawYoutubeGetChannelTranscriptsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// How to sort channel videos before selecting
type RawYoutubeGetChannelTranscriptsParamsSortBy string

const (
	RawYoutubeGetChannelTranscriptsParamsSortByPopular RawYoutubeGetChannelTranscriptsParamsSortBy = "popular"
	RawYoutubeGetChannelTranscriptsParamsSortByNewest  RawYoutubeGetChannelTranscriptsParamsSortBy = "newest"
	RawYoutubeGetChannelTranscriptsParamsSortByOldest  RawYoutubeGetChannelTranscriptsParamsSortBy = "oldest"
)

type RawYoutubeGetTranscriptParams struct {
	// Language code or "auto" for automatic detection
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RawYoutubeGetTranscriptParams]'s query parameters as
// `url.Values`.
func (r RawYoutubeGetTranscriptParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RawYoutubeSearchParams struct {
	// Search query
	Q string `query:"q" api:"required" json:"-"`
	// Country code for localized results (ISO 3166-1 alpha-2)
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Language code for results
	LanguageCode param.Opt[string] `query:"language_code,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RawYoutubeSearchParams]'s query parameters as `url.Values`.
func (r RawYoutubeSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
