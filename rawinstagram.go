// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package influshipapi

import (
	"context"
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
)

// Fetch fresh data directly from social platforms in real-time. Use when you need
// the most current information or data for profiles not yet in our database.
//
// RawInstagramService contains methods and other services that help with
// interacting with the Influship API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRawInstagramService] method instead.
type RawInstagramService struct {
	options []option.RequestOption
}

// NewRawInstagramService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRawInstagramService(opts ...option.RequestOption) (r RawInstagramService) {
	r = RawInstagramService{}
	r.options = opts
	return
}

// Fetch fresh Instagram profile data directly from Instagram in real-time. Use
// this when you need the most current follower counts, bio, or recent activity.
//
// **When to use live scraping:**
//
// - Profile not found in our database
// - Need real-time follower/engagement data
// - Verifying current profile status before campaign
//
// **Note:** Live scraping is slower than cached data (2-5 seconds) and costs more.
// Use cached endpoints when freshness isn't critical.
//
// **Pricing**: 0.5 credits per profile scraped ($0.005)
func (r *RawInstagramService) GetProfile(ctx context.Context, username string, query RawInstagramGetProfileParams, opts ...option.RequestOption) (res *RawInstagramGetProfileResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/raw/instagram/profile/%s", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type RawInstagramGetProfileResponse struct {
	// Live scraped profile data
	Data RawInstagramGetProfileResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawInstagramGetProfileResponse) RawJSON() string { return r.JSON.raw }
func (r *RawInstagramGetProfileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Live scraped profile data
type RawInstagramGetProfileResponseData struct {
	// Profile unique identifier
	ID       string                                     `json:"id" api:"required" format:"uuid"`
	Activity RawInstagramGetProfileResponseDataActivity `json:"activity" api:"required"`
	// Avatar URL
	AvatarURL string `json:"avatar_url" api:"required" format:"uri"`
	// Profile bio
	Bio string `json:"bio" api:"required"`
	// Account category
	Category string `json:"category" api:"required"`
	// Creator unique identifier
	CreatorID string `json:"creator_id" api:"required" format:"uuid"`
	// Last data refresh timestamp
	DataUpdatedAt time.Time `json:"data_updated_at" api:"required" format:"date-time"`
	// Display name
	DisplayName string `json:"display_name" api:"required"`
	// External website URL
	ExternalURL string                                   `json:"external_url" api:"required" format:"uri"`
	Growth      RawInstagramGetProfileResponseDataGrowth `json:"growth" api:"required"`
	// Whether this is a business account
	IsBusiness bool `json:"is_business" api:"required"`
	// Whether the account is private
	IsPrivate bool `json:"is_private" api:"required"`
	// Whether the account is verified
	IsVerified bool                                      `json:"is_verified" api:"required"`
	Metrics    RawInstagramGetProfileResponseDataMetrics `json:"metrics" api:"required"`
	// Social media platform
	//
	// Any of "instagram".
	Platform string `json:"platform" api:"required"`
	// Listed pronouns
	Pronouns []string `json:"pronouns" api:"required"`
	// When this data was scraped
	ScrapedAt time.Time `json:"scraped_at" api:"required" format:"date-time"`
	// Profile URL
	URL string `json:"url" api:"required" format:"uri"`
	// Profile username
	Username string `json:"username" api:"required"`
	// Recent posts (only included when include_posts=true)
	Posts []RawInstagramGetProfileResponseDataPost `json:"posts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Activity      respjson.Field
		AvatarURL     respjson.Field
		Bio           respjson.Field
		Category      respjson.Field
		CreatorID     respjson.Field
		DataUpdatedAt respjson.Field
		DisplayName   respjson.Field
		ExternalURL   respjson.Field
		Growth        respjson.Field
		IsBusiness    respjson.Field
		IsPrivate     respjson.Field
		IsVerified    respjson.Field
		Metrics       respjson.Field
		Platform      respjson.Field
		Pronouns      respjson.Field
		ScrapedAt     respjson.Field
		URL           respjson.Field
		Username      respjson.Field
		Posts         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawInstagramGetProfileResponseData) RawJSON() string { return r.JSON.raw }
func (r *RawInstagramGetProfileResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawInstagramGetProfileResponseDataActivity struct {
	// Timestamp of last post
	LastPostAt time.Time `json:"last_post_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastPostAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawInstagramGetProfileResponseDataActivity) RawJSON() string { return r.JSON.raw }
func (r *RawInstagramGetProfileResponseDataActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawInstagramGetProfileResponseDataGrowth struct {
	// Follower growth percentage over 30 days (e.g. 2.5 means +2.5%)
	Followers30dPct float64 `json:"followers_30d_pct" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Followers30dPct respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawInstagramGetProfileResponseDataGrowth) RawJSON() string { return r.JSON.raw }
func (r *RawInstagramGetProfileResponseDataGrowth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawInstagramGetProfileResponseDataMetrics struct {
	// Average comments on recent posts
	AvgCommentsRecent float64 `json:"avg_comments_recent" api:"required"`
	// Average likes on recent posts
	AvgLikesRecent float64 `json:"avg_likes_recent" api:"required"`
	// Average views on recent posts
	AvgViewsRecent float64 `json:"avg_views_recent" api:"required"`
	// Engagement rate as a percentage (e.g. 3.5 means 3.5%)
	EngagementRate float64 `json:"engagement_rate" api:"required"`
	// Follower count
	Followers int64 `json:"followers" api:"required"`
	// Following count
	Following int64 `json:"following" api:"required"`
	// Total post count
	Posts int64 `json:"posts" api:"required"`
	// Posts in the last 30 days
	PostsLast30d int64 `json:"posts_last_30d" api:"required"`
	// Average posts per week
	PostsPerWeek float64 `json:"posts_per_week" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgCommentsRecent respjson.Field
		AvgLikesRecent    respjson.Field
		AvgViewsRecent    respjson.Field
		EngagementRate    respjson.Field
		Followers         respjson.Field
		Following         respjson.Field
		Posts             respjson.Field
		PostsLast30d      respjson.Field
		PostsPerWeek      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawInstagramGetProfileResponseDataMetrics) RawJSON() string { return r.JSON.raw }
func (r *RawInstagramGetProfileResponseDataMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Simplified post from live scrape
type RawInstagramGetProfileResponseDataPost struct {
	// Post unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Post caption
	Caption string `json:"caption" api:"required"`
	// Comment count
	CommentsCount int64 `json:"comments_count" api:"required"`
	// Like count
	LikesCount int64 `json:"likes_count" api:"required"`
	// Primary media URL
	MediaURL string `json:"media_url" api:"required" format:"uri"`
	// Platform-specific post ID
	PlatformID string `json:"platform_id" api:"required"`
	// Post timestamp
	PostedAt time.Time `json:"posted_at" api:"required" format:"date-time"`
	// Type of post
	//
	// Any of "image", "video", "carousel", "reel", "story".
	Type string `json:"type" api:"required"`
	// Post URL
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Caption       respjson.Field
		CommentsCount respjson.Field
		LikesCount    respjson.Field
		MediaURL      respjson.Field
		PlatformID    respjson.Field
		PostedAt      respjson.Field
		Type          respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RawInstagramGetProfileResponseDataPost) RawJSON() string { return r.JSON.raw }
func (r *RawInstagramGetProfileResponseDataPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RawInstagramGetProfileParams struct {
	// Include recent posts in response
	IncludePosts param.Opt[bool] `query:"include_posts,omitzero" json:"-"`
	// Number of posts to include
	PostLimit param.Opt[int64] `query:"post_limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RawInstagramGetProfileParams]'s query parameters as
// `url.Values`.
func (r RawInstagramGetProfileParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
