// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package influshipapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Influship/influship-go"
	"github.com/Influship/influship-go/internal/testutil"
	"github.com/Influship/influship-go/option"
)

func TestSearchNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := influshipapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Search.New(context.TODO(), influshipapi.SearchNewParams{
		Query: "fitness influencers with 100k+ followers who post workout videos",
		Filters: influshipapi.SearchNewParamsFilters{
			EngagementRate: influshipapi.SearchNewParamsFiltersEngagementRate{
				Max: influshipapi.Float(10),
				Min: influshipapi.Float(1.5),
			},
			Followers: influshipapi.SearchNewParamsFiltersFollowers{
				Max: influshipapi.Float(500000),
				Min: influshipapi.Float(10000),
			},
			Verified: influshipapi.Bool(true),
		},
		Limit:     influshipapi.Int(25),
		Platforms: []string{"instagram"},
	})
	if err != nil {
		var apierr *influshipapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := influshipapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Search.Get(
		context.TODO(),
		"123e4567-e89b-12d3-a456-426614174000",
		influshipapi.SearchGetParams{
			Cursor: influshipapi.String("cursor"),
			Limit:  influshipapi.Int(25),
		},
	)
	if err != nil {
		var apierr *influshipapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
