package sandbox

import (
	"context"
	"fmt"
	"strconv"
	"time"

	soracom "github.com/soracom-labs/openapi-client-go/openapi/api"
	"github.com/soracom-labs/openapi-client-go/openapi/sandbox"
)

type TestSandboxConfigs struct {
	apiKey *string
	token  *string
}

func (c *TestSandboxConfigs) GetContext(ctx context.Context) context.Context {
	contextApiKeys := map[string]soracom.APIKey{
		"api_key": {
			Key: *c.apiKey,
		},
		"api_token": {
			Key: *c.token,
		},
	}
	return context.WithValue(context.Background(), soracom.ContextAPIKeys, contextApiKeys)
}

func GetClient() *soracom.APIClient {
	soracomConfiguration := soracom.NewConfiguration()
	// fixme: adhoc impl to overwrite endpoint url
	soracomConfiguration.Servers = soracom.ServerConfigurations{
		soracom.ServerConfiguration{
			URL: sandbox.NewConfiguration().Servers[0].URL, // fixme: verify index access for safety
		},
	}
	return soracom.NewAPIClient(soracomConfiguration)
}

func NewTestSandboxConfigs(authKeyId, authKey string) (*TestSandboxConfigs, error) {
	randStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	sandboxInitRequest := *sandbox.NewSandboxInitRequest(authKey, authKeyId, fmt.Sprintf("openapi-client-go-test+%s@soracom.jp", randStr), fmt.Sprintf("Password%s", randStr))
	sandboxInitRequest.CoverageTypes = append(sandboxInitRequest.CoverageTypes, "g", "jp")
	configuration := sandbox.NewConfiguration()
	client := sandbox.NewAPIClient(configuration)
	resp, r, err := client.OperatorApi.SandboxInitializeOperator(context.Background()).SandboxInitRequest(sandboxInitRequest).Execute()
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid HTTP response: %v", r.StatusCode)
	}
	return &TestSandboxConfigs{
		apiKey: resp.ApiKey,
		token:  resp.Token,
	}, nil
}

func (c *TestSandboxConfigs) CreateSubscriber(subscription string) (*sandbox.SandboxCreateSubscriberResponse, error) {
	sandboxCreateSubscriberRequest := *sandbox.NewSandboxCreateSubscriberRequest()
	sandboxCreateSubscriberRequest.Subscription = &subscription
	configuration := sandbox.NewConfiguration()
	client := sandbox.NewAPIClient(configuration)
	ctx := c.GetContext(context.Background())
	resp, r, err := client.SubscriberApi.SandboxCreateSubscriber(ctx).SandboxCreateSubscriberRequest(sandboxCreateSubscriberRequest).Execute()
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid HTTP response: %v", r.StatusCode)
	}
	return resp, nil
}

func (c *TestSandboxConfigs) RegisterSubscriber(subscriber *sandbox.SandboxCreateSubscriberResponse) (*soracom.Subscriber, error) {
	registerSubscribersRequest := *soracom.NewRegisterSubscribersRequest(*subscriber.RegistrationSecret)
	ctx := c.GetContext(context.Background())
	client := GetClient()
	resp, r, err := client.SubscriberApi.RegisterSubscriber(ctx, *subscriber.Imsi).RegisterSubscribersRequest(registerSubscribersRequest).Execute()
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid HTTP response: %v", r.StatusCode)
	}
	return resp, nil
}
