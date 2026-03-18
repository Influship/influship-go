// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package influshipapi

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/influship-api-go/internal/apijson"
	"github.com/stainless-sdks/influship-api-go/internal/requestconfig"
	"github.com/stainless-sdks/influship-api-go/option"
	"github.com/stainless-sdks/influship-api-go/packages/respjson"
)

// API health and status endpoints
//
// HealthService contains methods and other services that help with interacting
// with the Influship API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.options = opts
	return
}

// Check API health status. No authentication required.
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Health check response
type HealthCheckResponse struct {
	// Service health status
	Ok bool `json:"ok" api:"required"`
	// Current server timestamp
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ok          respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
