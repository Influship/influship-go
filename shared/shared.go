// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/Influship/influship-go/internal/apijson"
	"github.com/Influship/influship-go/packages/param"
	"github.com/Influship/influship-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Basic creator information
type CreatorBasic struct {
	// Creator unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Avatar URL
	AvatarURL string `json:"avatar_url" api:"required" format:"uri"`
	// Creator bio
	Bio string `json:"bio" api:"required"`
	// Creator display name
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AvatarURL   respjson.Field
		Bio         respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatorBasic) RawJSON() string { return r.JSON.raw }
func (r *CreatorBasic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Abbreviated profile information
type ProfileSummary struct {
	// Profile unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Engagement rate as a percentage, null if unknown (e.g. 3.5 means 3.5%)
	EngagementRate float64 `json:"engagement_rate" api:"required"`
	// Follower count (null if unknown)
	Followers int64 `json:"followers" api:"required"`
	// Whether the account is verified
	IsVerified bool `json:"is_verified" api:"required"`
	// Social media platform
	//
	// Any of "instagram".
	Platform ProfileSummaryPlatform `json:"platform" api:"required"`
	// Profile URL
	URL string `json:"url" api:"required" format:"uri"`
	// Profile username
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		EngagementRate respjson.Field
		Followers      respjson.Field
		IsVerified     respjson.Field
		Platform       respjson.Field
		URL            respjson.Field
		Username       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileSummary) RawJSON() string { return r.JSON.raw }
func (r *ProfileSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Social media platform
type ProfileSummaryPlatform string

const (
	ProfileSummaryPlatformInstagram ProfileSummaryPlatform = "instagram"
)
