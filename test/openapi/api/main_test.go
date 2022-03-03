package api

import (
	"log"
	"os"
	"testing"

	soracom "github.com/soracom-labs/openapi-client-go/openapi/api"
	testSandbox "github.com/soracom-labs/openapi-client-go/test/openapi/sandbox"
)

type testApiConfigs struct {
	subscriber     *soracom.Subscriber
	sandboxConfigs *testSandbox.TestSandboxConfigs
}

var (
	globalTestApiConfigs *testApiConfigs
)

//nolint:staticcheck, no need to call os.Exit() from go 1.15+
func TestMain(m *testing.M) {
	authKeyId := os.Getenv("SORACOM_AUTH_KEY_ID")
	authKey := os.Getenv("SORACOM_AUTH_KEY")

	configs, err := testSandbox.NewTestSandboxConfigs(authKeyId, authKey)
	if err != nil {
		log.Fatalf("NewConfigs() failed with %v", err)
	}

	createdSubscriber, err := configs.CreateSubscriber("plan01s")
	if err != nil {
		log.Fatalf("CreateSubscriber() failed with %v", err)
	}

	subscriber, err := configs.RegisterSubscriber(createdSubscriber)
	if err != nil {
		log.Fatalf("RegisterSubscriber() failed with %v", err)
	}

	globalTestApiConfigs = &testApiConfigs{
		subscriber:     subscriber,
		sandboxConfigs: configs,
	}

	m.Run()
}
