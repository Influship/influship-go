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

func TestAutoPagination(t *testing.T) {
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
	iter := client.Search.GetAutoPaging(
		context.TODO(),
		"search_abc123",
		influshipapi.SearchGetParams{
			Cursor: influshipapi.String("eyJvZmZzZXQiOjEwfQ=="),
			Limit:  influshipapi.Int(10),
		},
	)
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		search := iter.Current()
		t.Logf("%+v\n", search.Creator)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
