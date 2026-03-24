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
	"github.com/Influship/influship-go/internal/requestconfig"
	"github.com/Influship/influship-go/option"
	"github.com/Influship/influship-go/packages/param"
	"github.com/Influship/influship-go/packages/respjson"
)

// Access individual social media profiles with detailed metrics, growth data, and
// activity information. Profiles are platform-specific accounts linked to
// creators.
//
// ProfileService contains methods and other services that help with interacting
// with the Influship API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.options = opts
	return
}

// Retrieve detailed profile data including metrics, growth statistics, and
// activity information from our database.
//
// **Response includes:**
//
// - Basic info (bio, avatar, verification status)
// - Performance metrics (followers, engagement rate, avg likes/comments)
// - Growth data (30-day follower growth, monthly rate)
// - Activity data (last post date, posting frequency)
//
// **Pricing**: 0.1 credits per request ($0.001)
func (r *ProfileService) Get(ctx context.Context, username string, query ProfileGetParams, opts ...option.RequestOption) (res *ProfileGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.Platform == "" {
		err = errors.New("missing required platform parameter")
		return nil, err
	}
	if username == "" {
		err = errors.New("missing required username parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/profiles/%s/%s", url.PathEscape(query.Platform), url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Look up multiple profiles in a single request. Efficiently retrieve data for up
// to 100 profiles at once.
//
// **Response includes:**
//
//   - `found`: Array of profiles that exist in our database
//   - `not_found`: Array of profiles that weren't found (consider live scraping
//     these)
//
// **Pricing**: 0.1 credits per profile ($0.001)
func (r *ProfileService) Lookup(ctx context.Context, body ProfileLookupParams, opts ...option.RequestOption) (res *ProfileLookupResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/profiles/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Profile activity information
type ProfileActivity struct {
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
func (r ProfileActivity) RawJSON() string { return r.JSON.raw }
func (r *ProfileActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile growth statistics
type ProfileGrowth struct {
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
func (r ProfileGrowth) RawJSON() string { return r.JSON.raw }
func (r *ProfileGrowth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Profile performance metrics
type ProfileMetrics struct {
	// Average comments on recent posts
	AvgCommentsRecent float64 `json:"avg_comments_recent" api:"required"`
	// Average likes on recent posts
	AvgLikesRecent float64 `json:"avg_likes_recent" api:"required"`
	// Average views on recent posts (for video content)
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
func (r ProfileMetrics) RawJSON() string { return r.JSON.raw }
func (r *ProfileMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full profile details
type ProfileResponseData struct {
	// Profile unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Profile activity information
	Activity ProfileActivity `json:"activity" api:"required"`
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
	// External website URL from bio
	ExternalURL string `json:"external_url" api:"required" format:"uri"`
	// Profile growth statistics
	Growth ProfileGrowth `json:"growth" api:"required"`
	// Whether this is a business account
	IsBusiness bool `json:"is_business" api:"required"`
	// Whether the account is private
	IsPrivate bool `json:"is_private" api:"required"`
	// Whether the account is verified
	IsVerified bool `json:"is_verified" api:"required"`
	// Profile performance metrics
	Metrics ProfileMetrics `json:"metrics" api:"required"`
	// Social media platform
	//
	// Any of "instagram".
	Platform ProfileResponseDataPlatform `json:"platform" api:"required"`
	// Listed pronouns
	Pronouns []string `json:"pronouns" api:"required"`
	// Profile URL
	URL string `json:"url" api:"required" format:"uri"`
	// Profile username
	Username string `json:"username" api:"required"`
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
		URL           respjson.Field
		Username      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Social media platform
type ProfileResponseDataPlatform string

const (
	ProfileResponseDataPlatformInstagram ProfileResponseDataPlatform = "instagram"
)

type ProfileGetResponse struct {
	// Full profile details
	Data ProfileResponseData `json:"data" api:"required"`
	// Present when partial results were returned because profile metrics/data were
	// skipped due to integrity issues.
	Warning string `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Warning     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileLookupResponse struct {
	// Profiles that were found
	Data []ProfileResponseData `json:"data" api:"required"`
	// Profiles that were not found
	NotFound []ProfileLookupResponseNotFound `json:"not_found" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NotFound    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileLookupResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileLookupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileLookupResponseNotFound struct {
	// Social media platform
	//
	// Any of "instagram".
	Platform string `json:"platform" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Platform    respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileLookupResponseNotFound) RawJSON() string { return r.JSON.raw }
func (r *ProfileLookupResponseNotFound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileGetParams struct {
	// Platform name
	Platform string `path:"platform" api:"required" json:"-"`
	paramObj
}

type ProfileLookupParams struct {
	// Profiles to lookup
	Profiles []ProfileLookupParamsProfile `json:"profiles,omitzero" api:"required"`
	paramObj
}

func (r ProfileLookupParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileLookupParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileLookupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Platform, Username are required.
type ProfileLookupParamsProfile struct {
	// Social media platform
	//
	// Any of "instagram".
	Platform string `json:"platform,omitzero" api:"required"`
	// Username to lookup
	Username string `json:"username" api:"required"`
	paramObj
}

func (r ProfileLookupParamsProfile) MarshalJSON() (data []byte, err error) {
	type shadow ProfileLookupParamsProfile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileLookupParamsProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ProfileLookupParamsProfile](
		"platform", "instagram",
	)
}
