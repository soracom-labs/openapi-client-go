package sandbox

import (
	"log"
	"os"
	"testing"
)

//nolint:staticcheck, no need to call os.Exit() from go 1.15+
func TestMain(m *testing.M) {
	authKeyId := os.Getenv("SORACOM_AUTH_KEY_ID")
	authKey := os.Getenv("SORACOM_AUTH_KEY")

	configs, err := NewTestSandboxConfigs(authKeyId, authKey)
	if err != nil {
		log.Fatalf("NewTestSandboxConfigs() failed with %v", err)
	}

	createdSubscriber, err := configs.CreateSubscriber("plan01s")
	if err != nil {
		log.Fatalf("CreateSubscriber() failed with %v", err)
	}

	_, err = configs.RegisterSubscriber(createdSubscriber)
	if err != nil {
		log.Fatalf("RegisterSubscriber() failed with %v", err)
	}

	m.Run()
}
