// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package influshipapi

import (
	"github.com/stainless-sdks/influship-api-go/option"
)

// RawService contains methods and other services that help with interacting with
// the Influship API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRawService] method instead.
type RawService struct {
	options []option.RequestOption
	// Fetch fresh data directly from social platforms in real-time. Use when you need
	// the most current information or data for profiles not yet in our database.
	Instagram RawInstagramService
	// Fetch fresh data directly from social platforms in real-time. Use when you need
	// the most current information or data for profiles not yet in our database.
	Youtube RawYoutubeService
}

// NewRawService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRawService(opts ...option.RequestOption) (r RawService) {
	r = RawService{}
	r.options = opts
	r.Instagram = NewRawInstagramService(opts...)
	r.Youtube = NewRawYoutubeService(opts...)
	return
}
