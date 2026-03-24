// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package influshipapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Influship/influship-go/internal/apijson"
	"github.com/Influship/influship-go/internal/apiquery"
	"github.com/Influship/influship-go/internal/requestconfig"
	"github.com/Influship/influship-go/option"
	"github.com/Influship/influship-go/packages/pagination"
	"github.com/Influship/influship-go/packages/param"
	"github.com/Influship/influship-go/packages/respjson"
	"github.com/Influship/influship-go/shared"
)

// AI-powered semantic search to find creators using natural language queries.
// Understands intent and context to match creators based on content themes,
// audience, and style.
//
// SearchService contains methods and other services that help with interacting
// with the Influship API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.options = opts
	return
}

// Search for creators using natural language queries. The AI understands intent
// and context to match creators based on content themes, audience demographics,
// and style.
//
// The response includes a `search_id` that can be used with `GET /v1/search/{id}`
// to paginate through results for free.
//
// **Use cases:**
//
//   - Find creators in a specific niche ("vegan food bloggers in LA")
//   - Discover creators with specific audience characteristics ("fitness influencers
//     with millennial audience")
//   - Search by content style ("creators who post cinematic travel videos")
//
// **Pricing**: 25 credits base + 2 credits per creator returned
func (r *SearchService) New(ctx context.Context, body SearchNewParams, opts ...option.RequestOption) (res *SearchNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Paginate through results from a previous search. Use the `search_id` returned by
// `POST /v1/search` to fetch additional pages.
//
// Search sessions expire after 1 hour. After expiry, a new search must be run.
//
// **Pricing**: 0 credits (included with initial search)
func (r *SearchService) Get(ctx context.Context, id string, query SearchGetParams, opts ...option.RequestOption) (res *pagination.QueryCursor[SearchGetResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/search/%s", id)
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

// Paginate through results from a previous search. Use the `search_id` returned by
// `POST /v1/search` to fetch additional pages.
//
// Search sessions expire after 1 hour. After expiry, a new search must be run.
//
// **Pricing**: 0 credits (included with initial search)
func (r *SearchService) GetAutoPaging(ctx context.Context, id string, query SearchGetParams, opts ...option.RequestOption) *pagination.QueryCursorAutoPager[SearchGetResponse] {
	return pagination.NewQueryCursorAutoPager(r.Get(ctx, id, query, opts...))
}

// Search match information
type MatchInfo struct {
	// Human-readable match reasons
	Reasons []string `json:"reasons" api:"required"`
	// Match relevance score (0-1)
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reasons     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MatchInfo) RawJSON() string { return r.JSON.raw }
func (r *MatchInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchNewResponse struct {
	Data []SearchNewResponseData `json:"data" api:"required"`
	// Whether more results are available
	HasMore bool `json:"has_more" api:"required"`
	// Cursor for the next page
	NextCursor string `json:"next_cursor" api:"required"`
	// Search ID. Use with GET /v1/search/{id} for free pagination.
	SearchID string `json:"search_id" api:"required" format:"uuid"`
	// Total number of results across all pages
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		SearchID    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchNewResponseData struct {
	// Basic creator information
	Creator shared.CreatorBasic `json:"creator" api:"required"`
	// Search match information
	Match MatchInfo `json:"match" api:"required"`
	// Abbreviated profile information
	PrimaryProfile shared.ProfileSummary `json:"primary_profile" api:"required"`
	// Abbreviated profile information
	RelevantProfile shared.ProfileSummary `json:"relevant_profile" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Creator         respjson.Field
		Match           respjson.Field
		PrimaryProfile  respjson.Field
		RelevantProfile respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *SearchNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchGetResponse struct {
	// Basic creator information
	Creator shared.CreatorBasic `json:"creator" api:"required"`
	// Search match information
	Match MatchInfo `json:"match" api:"required"`
	// Abbreviated profile information
	PrimaryProfile shared.ProfileSummary `json:"primary_profile" api:"required"`
	// Abbreviated profile information
	RelevantProfile shared.ProfileSummary `json:"relevant_profile" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Creator         respjson.Field
		Match           respjson.Field
		PrimaryProfile  respjson.Field
		RelevantProfile respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchNewParams struct {
	// Natural language search query
	Query string `json:"query" api:"required"`
	// Maximum results to return
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Additional filters
	Filters SearchNewParamsFilters `json:"filters,omitzero"`
	// Filter results to specific platforms
	//
	// Any of "instagram".
	Platforms []string `json:"platforms,omitzero"`
	paramObj
}

func (r SearchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SearchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional filters
type SearchNewParamsFilters struct {
	// Filter by verified status
	Verified param.Opt[bool] `json:"verified,omitzero"`
	// Filter by engagement rate
	EngagementRate SearchNewParamsFiltersEngagementRate `json:"engagement_rate,omitzero"`
	// Filter by follower count
	Followers SearchNewParamsFiltersFollowers `json:"followers,omitzero"`
	paramObj
}

func (r SearchNewParamsFilters) MarshalJSON() (data []byte, err error) {
	type shadow SearchNewParamsFilters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchNewParamsFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by engagement rate
type SearchNewParamsFiltersEngagementRate struct {
	// Maximum engagement rate (%)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum engagement rate (%)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SearchNewParamsFiltersEngagementRate) MarshalJSON() (data []byte, err error) {
	type shadow SearchNewParamsFiltersEngagementRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchNewParamsFiltersEngagementRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by follower count
type SearchNewParamsFiltersFollowers struct {
	// Maximum follower count
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum follower count
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SearchNewParamsFiltersFollowers) MarshalJSON() (data []byte, err error) {
	type shadow SearchNewParamsFiltersFollowers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchNewParamsFiltersFollowers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchGetParams struct {
	// Pagination cursor for next page
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SearchGetParams]'s query parameters as `url.Values`.
func (r SearchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
