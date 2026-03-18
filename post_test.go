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

func TestPostListWithOptionalParams(t *testing.T) {
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
	_, err := client.Posts.List(context.TODO(), influshipapi.PostListParams{
		CreatorID: influshipapi.String("123e4567-e89b-12d3-a456-426614174000"),
		Cursor:    influshipapi.String("cursor"),
		Limit:     influshipapi.Int(12),
		Platform:  influshipapi.PostListParamsPlatformInstagram,
		Sort:      influshipapi.PostListParamsSortRecent,
		Username:  influshipapi.String("fitness_coach_jane"),
	})
	if err != nil {
		var apierr *influshipapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
