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

func TestCreatorGet(t *testing.T) {
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
	_, err := client.Creators.Get(
		context.TODO(),
		"123e4567-e89b-12d3-a456-426614174000",
		influshipapi.CreatorGetParams{
			Include: []string{"profiles"},
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

func TestCreatorAutocompleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Creators.Autocomplete(context.TODO(), influshipapi.CreatorAutocompleteParams{
		Q:        "fitness",
		Limit:    influshipapi.Int(5),
		Platform: influshipapi.CreatorAutocompleteParamsPlatformInstagram,
		Scope:    influshipapi.CreatorAutocompleteParamsScopeAllPlatforms,
	})
	if err != nil {
		var apierr *influshipapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreatorLookalikeWithOptionalParams(t *testing.T) {
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
	_, err := client.Creators.Lookalike(context.TODO(), influshipapi.CreatorLookalikeParams{
		Seeds: []influshipapi.CreatorLookalikeParamsSeed{{
			CreatorID: influshipapi.String("123e4567-e89b-12d3-a456-426614174000"),
			Platform:  "instagram",
			Username:  influshipapi.String("fitness_coach_jane"),
			Weight:    influshipapi.Float(1),
		}},
		Cursor: influshipapi.String("cursor"),
		Filters: influshipapi.CreatorLookalikeParamsFilters{
			EngagementRate: influshipapi.CreatorLookalikeParamsFiltersEngagementRate{
				Max: influshipapi.Float(10),
				Min: influshipapi.Float(1.5),
			},
			Followers: influshipapi.CreatorLookalikeParamsFiltersFollowers{
				Max: influshipapi.Float(500000),
				Min: influshipapi.Float(10000),
			},
			Verified: influshipapi.Bool(true),
		},
		Limit: influshipapi.Int(25),
	})
	if err != nil {
		var apierr *influshipapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreatorMatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Creators.Match(context.TODO(), influshipapi.CreatorMatchParams{
		Creators: []influshipapi.CreatorMatchParamsCreator{{
			CreatorID: influshipapi.String("123e4567-e89b-12d3-a456-426614174000"),
			Platform:  "instagram",
			Username:  influshipapi.String("fitness_coach_jane"),
		}},
		Intent: influshipapi.CreatorMatchParamsIntent{
			Query:   "Looking for fitness influencers to promote our new protein bar",
			Context: influshipapi.String("Target audience is health-conscious millennials"),
		},
	})
	if err != nil {
		var apierr *influshipapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
