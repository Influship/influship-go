// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package influshipapi

import (
	"github.com/Influship/influship-go/internal/apierror"
	"github.com/Influship/influship-go/packages/param"
	"github.com/Influship/influship-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Basic creator information
//
// This is an alias to an internal type.
type CreatorBasic = shared.CreatorBasic

// Abbreviated profile information
//
// This is an alias to an internal type.
type ProfileSummary = shared.ProfileSummary

// Social media platform
//
// This is an alias to an internal type.
type ProfileSummaryPlatform = shared.ProfileSummaryPlatform

// Equals "instagram"
const ProfileSummaryPlatformInstagram = shared.ProfileSummaryPlatformInstagram
