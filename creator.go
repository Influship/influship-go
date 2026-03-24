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

// Retrieve creator profiles and discover new creators through search,
// autocomplete, and lookalike matching. Creators are cross-platform entities that
// may have profiles on multiple social networks.
//
// CreatorService contains methods and other services that help with interacting
// with the Influship API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreatorService] method instead.
type CreatorService struct {
	options []option.RequestOption
}

// NewCreatorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCreatorService(opts ...option.RequestOption) (r CreatorService) {
	r = CreatorService{}
	r.options = opts
	return
}

// Retrieve a creator's profile including AI-generated summary, content themes, and
// optionally their linked social profiles.
//
// **What is a Creator?** A creator is a cross-platform entity representing a
// person or brand. They may have profiles on multiple social networks (Instagram,
// YouTube, TikTok, etc.) that are linked together.
//
// **Include options:**
//
// - `profiles`: Include all linked social profiles with metrics
//
// **Pricing**: 0.1 credits per request ($0.001)
func (r *CreatorService) Get(ctx context.Context, id string, query CreatorGetParams, opts ...option.RequestOption) (res *CreatorGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/creators/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fast typeahead search for creators by name or username. Optimized for
// search-as-you-type UIs with sub-100ms response times.
//
// **Matching behavior:**
//
// - Matches against creator name, username, and display name
// - Results include which field matched and the matching value
// - Prefix matching (e.g., "fit" matches "fitness_coach")
//
// **Scope options:**
//
// - `creator_only`: Return only the creator entity
// - `matched_platforms`: Return only profiles that matched the query
// - `all_platforms`: Return all linked profiles (default)
//
// **Pricing**: 0.05 credits per request ($0.0005)
func (r *CreatorService) Autocomplete(ctx context.Context, query CreatorAutocompleteParams, opts ...option.RequestOption) (res *CreatorAutocompleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/creators/autocomplete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Find creators similar to provided seed creators using AI-powered similarity
// matching. Analyzes content themes, audience overlap, posting style, and
// engagement patterns.
//
// **Use cases:**
//
// - Expand campaigns with creators similar to proven performers
// - Find alternatives when preferred creators are unavailable
// - Discover emerging creators in the same niche
//
// **How it works:**
//
// 1. Provide 1-10 seed creators (by ID or platform/username)
// 2. Optionally weight seeds to prioritize certain creators
// 3. Get ranked results with similarity scores and shared traits
//
// **Pricing**: 1.5 credits per creator returned ($0.015)
func (r *CreatorService) Lookalike(ctx context.Context, body CreatorLookalikeParams, opts ...option.RequestOption) (res *pagination.BodyCursor[CreatorLookalikeResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/creators/lookalike"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
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

// Find creators similar to provided seed creators using AI-powered similarity
// matching. Analyzes content themes, audience overlap, posting style, and
// engagement patterns.
//
// **Use cases:**
//
// - Expand campaigns with creators similar to proven performers
// - Find alternatives when preferred creators are unavailable
// - Discover emerging creators in the same niche
//
// **How it works:**
//
// 1. Provide 1-10 seed creators (by ID or platform/username)
// 2. Optionally weight seeds to prioritize certain creators
// 3. Get ranked results with similarity scores and shared traits
//
// **Pricing**: 1.5 credits per creator returned ($0.015)
func (r *CreatorService) LookalikeAutoPaging(ctx context.Context, body CreatorLookalikeParams, opts ...option.RequestOption) *pagination.BodyCursorAutoPager[CreatorLookalikeResponse] {
	return pagination.NewBodyCursorAutoPager(r.Lookalike(ctx, body, opts...))
}

// Evaluate how well creators match a specific campaign using AI analysis. Returns
// a fit score (0-1), decision recommendation (good/neutral/avoid), and
// evidence-based explanations.
//
// **Use cases:**
//
// - Vet shortlisted creators before outreach
// - Rank candidates for a specific campaign
// - Get AI-generated talking points for why a creator fits
//
// **How it works:**
//
// 1. Describe your campaign intent and target audience
// 2. Provide up to 100 creators to evaluate
// 3. Get detailed scores with explanations and evidence
//
// **Pricing**: 1 credit per creator scored ($0.01)
func (r *CreatorService) Match(ctx context.Context, body CreatorMatchParams, opts ...option.RequestOption) (res *CreatorMatchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/creators/match"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CreatorGetResponse struct {
	// Full creator details
	Data CreatorGetResponseData `json:"data" api:"required"`
	// Present when partial results were returned because one or more linked profiles
	// were skipped for data integrity reasons.
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
func (r CreatorGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CreatorGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full creator details
type CreatorGetResponseData struct {
	// Creator unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// AI-generated summary of the creator
	AISummary string `json:"ai_summary" api:"required"`
	// Avatar URL
	AvatarURL string `json:"avatar_url" api:"required" format:"uri"`
	// Creator bio
	Bio string `json:"bio" api:"required"`
	// Content themes/topics
	ContentThemes []string `json:"content_themes" api:"required"`
	// Creator display name
	Name string `json:"name" api:"required"`
	// Social profiles (only included when include=profiles)
	Profiles []shared.ProfileSummary `json:"profiles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AISummary     respjson.Field
		AvatarURL     respjson.Field
		Bio           respjson.Field
		ContentThemes respjson.Field
		Name          respjson.Field
		Profiles      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *CreatorGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorAutocompleteResponse struct {
	// Autocomplete results
	Data []CreatorAutocompleteResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorAutocompleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CreatorAutocompleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorAutocompleteResponseData struct {
	// Creator ID
	ID string `json:"id" api:"required" format:"uuid"`
	// Avatar URL
	Avatar string `json:"avatar" api:"required" format:"uri"`
	// Creator name
	Name string `json:"name" api:"required"`
	// Matching platforms
	Platforms []CreatorAutocompleteResponseDataPlatform `json:"platforms" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Avatar      respjson.Field
		Name        respjson.Field
		Platforms   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorAutocompleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *CreatorAutocompleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorAutocompleteResponseDataPlatform struct {
	DisplayName string `json:"display_name" api:"required"`
	// The field value that matched
	MatchField string `json:"match_field" api:"required"`
	// How the query matched this profile
	//
	// Any of "name", "username", "display_name".
	MatchType string `json:"match_type" api:"required"`
	// Social media platform
	//
	// Any of "instagram".
	Platform string `json:"platform" api:"required"`
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		MatchField  respjson.Field
		MatchType   respjson.Field
		Platform    respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorAutocompleteResponseDataPlatform) RawJSON() string { return r.JSON.raw }
func (r *CreatorAutocompleteResponseDataPlatform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorLookalikeResponse struct {
	// Basic creator information
	Creator shared.CreatorBasic `json:"creator" api:"required"`
	// Abbreviated profile information
	PrimaryProfile shared.ProfileSummary `json:"primary_profile" api:"required"`
	// Similarity information for lookalike match
	Similarity CreatorLookalikeResponseSimilarity `json:"similarity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Creator        respjson.Field
		PrimaryProfile respjson.Field
		Similarity     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorLookalikeResponse) RawJSON() string { return r.JSON.raw }
func (r *CreatorLookalikeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Similarity information for lookalike match
type CreatorLookalikeResponseSimilarity struct {
	// Similarity score (0-1)
	Score float64 `json:"score" api:"required"`
	// Shared traits with seed creators
	SharedTraits []string `json:"shared_traits" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Score        respjson.Field
		SharedTraits respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorLookalikeResponseSimilarity) RawJSON() string { return r.JSON.raw }
func (r *CreatorLookalikeResponseSimilarity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorMatchResponse struct {
	Data []CreatorMatchResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorMatchResponse) RawJSON() string { return r.JSON.raw }
func (r *CreatorMatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorMatchResponseData struct {
	Creator CreatorMatchResponseDataCreator `json:"creator" api:"required"`
	Input   CreatorMatchResponseDataInput   `json:"input" api:"required"`
	Match   CreatorMatchResponseDataMatch   `json:"match" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Creator     respjson.Field
		Input       respjson.Field
		Match       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorMatchResponseData) RawJSON() string { return r.JSON.raw }
func (r *CreatorMatchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorMatchResponseDataCreator struct {
	ID        string `json:"id" api:"required" format:"uuid"`
	AvatarURL string `json:"avatar_url" api:"required" format:"uri"`
	Name      string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AvatarURL   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorMatchResponseDataCreator) RawJSON() string { return r.JSON.raw }
func (r *CreatorMatchResponseDataCreator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorMatchResponseDataInput struct {
	CreatorID string `json:"creator_id" format:"uuid"`
	// Social media platform
	//
	// Any of "instagram".
	Platform string `json:"platform"`
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatorID   respjson.Field
		Platform    respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorMatchResponseDataInput) RawJSON() string { return r.JSON.raw }
func (r *CreatorMatchResponseDataInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorMatchResponseDataMatch struct {
	// Match decision recommendation
	//
	// Any of "good", "neutral", "avoid".
	Decision string `json:"decision" api:"required"`
	// Structured reasons supporting the decision
	Reasons []CreatorMatchResponseDataMatchReason `json:"reasons" api:"required"`
	// Match score (0-1)
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Decision    respjson.Field
		Reasons     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorMatchResponseDataMatch) RawJSON() string { return r.JSON.raw }
func (r *CreatorMatchResponseDataMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorMatchResponseDataMatchReason struct {
	// Human-readable reason for the match
	Text string `json:"text" api:"required"`
	// ID of the supporting fact, if applicable
	FactID string `json:"fact_id"`
	// ID of the source post, if applicable
	SourcePostID string `json:"source_post_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text         respjson.Field
		FactID       respjson.Field
		SourcePostID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorMatchResponseDataMatchReason) RawJSON() string { return r.JSON.raw }
func (r *CreatorMatchResponseDataMatchReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorGetParams struct {
	// Additional data to include in response
	//
	// Any of "profiles".
	Include []string `query:"include,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [CreatorGetParams]'s query parameters as `url.Values`.
func (r CreatorGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CreatorAutocompleteParams struct {
	// Search query (min 2 characters)
	Q string `query:"q" api:"required" json:"-"`
	// Maximum results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by platform
	//
	// Any of "instagram".
	Platform CreatorAutocompleteParamsPlatform `query:"platform,omitzero" json:"-"`
	// Which platforms to include in results
	//
	// Any of "creator_only", "matched_platforms", "all_platforms".
	Scope CreatorAutocompleteParamsScope `query:"scope,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CreatorAutocompleteParams]'s query parameters as
// `url.Values`.
func (r CreatorAutocompleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by platform
type CreatorAutocompleteParamsPlatform string

const (
	CreatorAutocompleteParamsPlatformInstagram CreatorAutocompleteParamsPlatform = "instagram"
)

// Which platforms to include in results
type CreatorAutocompleteParamsScope string

const (
	CreatorAutocompleteParamsScopeCreatorOnly      CreatorAutocompleteParamsScope = "creator_only"
	CreatorAutocompleteParamsScopeMatchedPlatforms CreatorAutocompleteParamsScope = "matched_platforms"
	CreatorAutocompleteParamsScopeAllPlatforms     CreatorAutocompleteParamsScope = "all_platforms"
)

type CreatorLookalikeParams struct {
	// Seed creators to find similar creators for
	Seeds []CreatorLookalikeParamsSeed `json:"seeds,omitzero" api:"required"`
	// Pagination cursor for next page
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// Maximum results to return
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Additional filters
	Filters CreatorLookalikeParamsFilters `json:"filters,omitzero"`
	paramObj
}

func (r CreatorLookalikeParams) MarshalJSON() (data []byte, err error) {
	type shadow CreatorLookalikeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatorLookalikeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Seed creator for lookalike search
type CreatorLookalikeParamsSeed struct {
	// Creator ID (use this OR platform+username)
	CreatorID param.Opt[string] `json:"creator_id,omitzero" format:"uuid"`
	// Username (required with platform)
	Username param.Opt[string] `json:"username,omitzero"`
	// Weight for this seed (0-1)
	Weight param.Opt[float64] `json:"weight,omitzero"`
	// Platform (required with username)
	//
	// Any of "instagram".
	Platform string `json:"platform,omitzero"`
	paramObj
}

func (r CreatorLookalikeParamsSeed) MarshalJSON() (data []byte, err error) {
	type shadow CreatorLookalikeParamsSeed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatorLookalikeParamsSeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CreatorLookalikeParamsSeed](
		"platform", "instagram",
	)
}

// Additional filters
type CreatorLookalikeParamsFilters struct {
	// Filter by verified status
	Verified param.Opt[bool] `json:"verified,omitzero"`
	// Filter by engagement rate
	EngagementRate CreatorLookalikeParamsFiltersEngagementRate `json:"engagement_rate,omitzero"`
	// Filter by follower count
	Followers CreatorLookalikeParamsFiltersFollowers `json:"followers,omitzero"`
	paramObj
}

func (r CreatorLookalikeParamsFilters) MarshalJSON() (data []byte, err error) {
	type shadow CreatorLookalikeParamsFilters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatorLookalikeParamsFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by engagement rate
type CreatorLookalikeParamsFiltersEngagementRate struct {
	// Maximum engagement rate (%)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum engagement rate (%)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r CreatorLookalikeParamsFiltersEngagementRate) MarshalJSON() (data []byte, err error) {
	type shadow CreatorLookalikeParamsFiltersEngagementRate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatorLookalikeParamsFiltersEngagementRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by follower count
type CreatorLookalikeParamsFiltersFollowers struct {
	// Maximum follower count
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum follower count
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r CreatorLookalikeParamsFiltersFollowers) MarshalJSON() (data []byte, err error) {
	type shadow CreatorLookalikeParamsFiltersFollowers
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatorLookalikeParamsFiltersFollowers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatorMatchParams struct {
	// Creators to evaluate
	Creators []CreatorMatchParamsCreator `json:"creators,omitzero" api:"required"`
	// Campaign intent for creator matching
	Intent CreatorMatchParamsIntent `json:"intent,omitzero" api:"required"`
	paramObj
}

func (r CreatorMatchParams) MarshalJSON() (data []byte, err error) {
	type shadow CreatorMatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatorMatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Creator identifier for match endpoint
type CreatorMatchParamsCreator struct {
	// Creator ID (use this OR platform+username)
	CreatorID param.Opt[string] `json:"creator_id,omitzero" format:"uuid"`
	// Username (required with platform)
	Username param.Opt[string] `json:"username,omitzero"`
	// Platform (required with username)
	//
	// Any of "instagram".
	Platform string `json:"platform,omitzero"`
	paramObj
}

func (r CreatorMatchParamsCreator) MarshalJSON() (data []byte, err error) {
	type shadow CreatorMatchParamsCreator
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatorMatchParamsCreator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CreatorMatchParamsCreator](
		"platform", "instagram",
	)
}

// Campaign intent for creator matching
//
// The property Query is required.
type CreatorMatchParamsIntent struct {
	// Campaign description
	Query string `json:"query" api:"required"`
	// Additional context about the campaign
	Context param.Opt[string] `json:"context,omitzero"`
	paramObj
}

func (r CreatorMatchParamsIntent) MarshalJSON() (data []byte, err error) {
	type shadow CreatorMatchParamsIntent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatorMatchParamsIntent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
