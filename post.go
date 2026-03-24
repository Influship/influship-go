// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package influshipapi

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Influship/influship-go/internal/apijson"
	"github.com/Influship/influship-go/internal/apiquery"
	"github.com/Influship/influship-go/internal/requestconfig"
	"github.com/Influship/influship-go/option"
	"github.com/Influship/influship-go/packages/pagination"
	"github.com/Influship/influship-go/packages/param"
	"github.com/Influship/influship-go/packages/respjson"
)

// Retrieve and analyze social media posts with engagement metrics, media content,
// and performance data.
//
// PostService contains methods and other services that help with interacting with
// the Influship API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostService] method instead.
type PostService struct {
	options []option.RequestOption
}

// NewPostService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPostService(opts ...option.RequestOption) (r PostService) {
	r = PostService{}
	r.options = opts
	return
}

// Retrieve posts for a creator or profile with engagement metrics and media data.
//
// **Query options:**
//
// - By creator: Use `creator_id` to get posts across all their profiles
// - By profile: Use `platform` + `username` for a specific profile's posts
//
// **Sort options:**
//
// - `recent`: Most recent posts first (default)
// - `top_engagement`: Highest engagement rate first
// - `most_likes`: Most likes first
// - `most_views`: Most views first (video content)
// - `most_comments`: Most comments first
//
// **Pricing**: 0.05 credits per post returned ($0.0005)
func (r *PostService) List(ctx context.Context, query PostListParams, opts ...option.RequestOption) (res *pagination.QueryCursor[PostListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/posts"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve posts for a creator or profile with engagement metrics and media data.
//
// **Query options:**
//
// - By creator: Use `creator_id` to get posts across all their profiles
// - By profile: Use `platform` + `username` for a specific profile's posts
//
// **Sort options:**
//
// - `recent`: Most recent posts first (default)
// - `top_engagement`: Highest engagement rate first
// - `most_likes`: Most likes first
// - `most_views`: Most views first (video content)
// - `most_comments`: Most comments first
//
// **Pricing**: 0.05 credits per post returned ($0.0005)
func (r *PostService) ListAutoPaging(ctx context.Context, query PostListParams, opts ...option.RequestOption) *pagination.QueryCursorAutoPager[PostListResponse] {
	return pagination.NewQueryCursorAutoPager(r.List(ctx, query, opts...))
}

// Full post details
type PostListResponse struct {
	// Post unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Post caption
	Caption string `json:"caption" api:"required"`
	// Hashtags used in the post
	Hashtags []string `json:"hashtags" api:"required"`
	// Post location information
	Location PostListResponseLocation `json:"location" api:"required"`
	// Post media information
	Media PostListResponseMedia `json:"media" api:"required"`
	// Usernames mentioned in the post
	Mentions []string `json:"mentions" api:"required"`
	// Post engagement metrics
	Metrics PostListResponseMetrics `json:"metrics" api:"required"`
	// Social media platform
	//
	// Any of "instagram".
	Platform PostListResponsePlatform `json:"platform" api:"required"`
	// Platform-specific post ID
	PlatformID string `json:"platform_id" api:"required"`
	// Post timestamp
	PostedAt time.Time `json:"posted_at" api:"required" format:"date-time"`
	// Profile unique identifier
	ProfileID string `json:"profile_id" api:"required" format:"uuid"`
	// Type of post
	//
	// Any of "image", "video", "carousel", "reel", "story".
	Type PostListResponseType `json:"type" api:"required"`
	// Post URL
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Caption     respjson.Field
		Hashtags    respjson.Field
		Location    respjson.Field
		Media       respjson.Field
		Mentions    respjson.Field
		Metrics     respjson.Field
		Platform    respjson.Field
		PlatformID  respjson.Field
		PostedAt    respjson.Field
		ProfileID   respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostListResponse) RawJSON() string { return r.JSON.raw }
func (r *PostListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Post location information
type PostListResponseLocation struct {
	// Location name
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostListResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *PostListResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Post media information
type PostListResponseMedia struct {
	// Video duration in seconds
	DurationSeconds float64 `json:"duration_seconds" api:"required"`
	// Thumbnail URL
	ThumbnailURL string `json:"thumbnail_url" api:"required" format:"uri"`
	// Media URL
	URL string `json:"url" api:"required" format:"uri"`
	// Video URL (for video content)
	VideoURL string `json:"video_url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationSeconds respjson.Field
		ThumbnailURL    respjson.Field
		URL             respjson.Field
		VideoURL        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostListResponseMedia) RawJSON() string { return r.JSON.raw }
func (r *PostListResponseMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Post engagement metrics
type PostListResponseMetrics struct {
	// Comment count
	Comments int64 `json:"comments" api:"required"`
	// Engagement rate for this post as a percentage (e.g. 3.8 means 3.8%)
	EngagementRate float64 `json:"engagement_rate" api:"required"`
	// Like count
	Likes int64 `json:"likes" api:"required"`
	// Share count
	Shares int64 `json:"shares" api:"required"`
	// View count (for video content)
	Views int64 `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comments       respjson.Field
		EngagementRate respjson.Field
		Likes          respjson.Field
		Shares         respjson.Field
		Views          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostListResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *PostListResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Social media platform
type PostListResponsePlatform string

const (
	PostListResponsePlatformInstagram PostListResponsePlatform = "instagram"
)

// Type of post
type PostListResponseType string

const (
	PostListResponseTypeImage    PostListResponseType = "image"
	PostListResponseTypeVideo    PostListResponseType = "video"
	PostListResponseTypeCarousel PostListResponseType = "carousel"
	PostListResponseTypeReel     PostListResponseType = "reel"
	PostListResponseTypeStory    PostListResponseType = "story"
)

type PostListParams struct {
	// Creator ID (use this OR platform+username)
	CreatorID param.Opt[string] `query:"creator_id,omitzero" format:"uuid" json:"-"`
	// Pagination cursor for next page
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum posts to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Username (required with platform)
	Username param.Opt[string] `query:"username,omitzero" json:"-"`
	// Platform (required with username)
	//
	// Any of "instagram".
	Platform PostListParamsPlatform `query:"platform,omitzero" json:"-"`
	// Sort order
	//
	// Any of "recent", "top_engagement", "most_likes", "most_views", "most_comments".
	Sort PostListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PostListParams]'s query parameters as `url.Values`.
func (r PostListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Platform (required with username)
type PostListParamsPlatform string

const (
	PostListParamsPlatformInstagram PostListParamsPlatform = "instagram"
)

// Sort order
type PostListParamsSort string

const (
	PostListParamsSortRecent        PostListParamsSort = "recent"
	PostListParamsSortTopEngagement PostListParamsSort = "top_engagement"
	PostListParamsSortMostLikes     PostListParamsSort = "most_likes"
	PostListParamsSortMostViews     PostListParamsSort = "most_views"
	PostListParamsSortMostComments  PostListParamsSort = "most_comments"
)
