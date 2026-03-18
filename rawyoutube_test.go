// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package influshipapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/influship-api-go"
	"github.com/stainless-sdks/influship-api-go/internal/testutil"
	"github.com/stainless-sdks/influship-api-go/option"
)

func TestRawYoutubeGetChannelWithOptionalParams(t *testing.T) {
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
	_, err := client.Raw.Youtube.GetChannel(
		context.TODO(),
		"@techreviews",
		influshipapi.RawYoutubeGetChannelParams{
			IncludeVideos: influshipapi.Bool(true),
			VideoLimit:    influshipapi.Int(12),
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

func TestRawYoutubeGetChannelTranscriptsWithOptionalParams(t *testing.T) {
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
	_, err := client.Raw.Youtube.GetChannelTranscripts(
		context.TODO(),
		"@techreviews",
		influshipapi.RawYoutubeGetChannelTranscriptsParams{
			IncludeSegments: influshipapi.Bool(false),
			Language:        influshipapi.String("en"),
			SortBy:          influshipapi.RawYoutubeGetChannelTranscriptsParamsSortByNewest,
			VideoLimit:      influshipapi.Int(5),
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

func TestRawYoutubeGetTranscriptWithOptionalParams(t *testing.T) {
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
	_, err := client.Raw.Youtube.GetTranscript(
		context.TODO(),
		"dQw4w9WgXcQ",
		influshipapi.RawYoutubeGetTranscriptParams{
			Language: influshipapi.String("en"),
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

func TestRawYoutubeSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Raw.Youtube.Search(context.TODO(), influshipapi.RawYoutubeSearchParams{
		Q:            "fitness workout",
		CountryCode:  influshipapi.String("US"),
		LanguageCode: influshipapi.String("en"),
		Limit:        influshipapi.Int(20),
	})
	if err != nil {
		var apierr *influshipapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
