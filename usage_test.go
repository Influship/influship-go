// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package influshipapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/Influship/influship-go"
	"github.com/Influship/influship-go/internal/testutil"
	"github.com/Influship/influship-go/option"
)

func TestUsage(t *testing.T) {
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
	search, err := client.Search.New(context.TODO(), influshipapi.SearchNewParams{
		Query: "sustainable fashion creators with engaged audiences",
		Limit: influshipapi.Int(25),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", search.SearchID)
}
