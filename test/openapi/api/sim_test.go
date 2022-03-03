package api

import (
	"context"
	"testing"

	testSandbox "github.com/soracom-labs/openapi-client-go/test/openapi/sandbox"
)

func TestListSims(t *testing.T) {
	testCases := []struct {
		Name       string
		HasError   bool
		Limit      int32
		StatusCode int
	}{
		{
			Name:       "nominal scenarios",
			HasError:   false,
			Limit:      10,
			StatusCode: 200,
		},
		{
			Name:       "non-nominal scenarios",
			HasError:   true,
			Limit:      -1,
			StatusCode: 500,
		},
	}

	// cache these for readability
	ctx := globalTestApiConfigs.sandboxConfigs.GetContext(context.Background())
	client := testSandbox.GetClient()

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			resp, r, err := client.SimApi.ListSims(ctx).Limit(testCase.Limit).Execute()
			if hasError := (err != nil); hasError != testCase.HasError {
				t.Errorf("got error %v, expected %v, err %v", hasError, testCase.HasError, err)
			}
			if r.StatusCode != testCase.StatusCode {
				t.Errorf("got unexpected status code, expected %v, actual %v", testCase.HasError, r.StatusCode)
			}
			if !testCase.HasError && r.StatusCode == 200 {
				// for nominal case only
				if int(testCase.Limit) < len(resp) {
					t.Errorf("got unexpected # of sims, actual got # of sims(=%v) should be equal to or less than limit one(=%v)", len(resp), testCase.Limit)
				}
			}
		})
	}
}
