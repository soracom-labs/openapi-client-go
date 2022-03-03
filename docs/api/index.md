# SORACOM API

## Example
### Authentication
To execute API calls, you must first authenticate your SORACOM account using one of the following authentication methods:

- **Root Account** - authenticate using your SORACOM `email` and `password` login information
- **AuthKey** - authenticate using a `AuthKeyId` and `AuthKey`

Here's an example to get API key with **AuthKey**

```go
package main

import (
	"context"
	"fmt"
	"os"

	soracom "github.com/soracom-labs/openapi-client-go/openapi/api"
)


func main() {
	client := soracom.NewAPIClient(soracom.NewConfiguration())
	authKeyId := os.Getenv("SORACOM_AUTH_KEY_ID")
	authKey := os.Getenv("SORACOM_AUTH_KEY")

	authRequest := soracom.AuthRequest{
		AuthKeyId: &authKeyId,
		AuthKey:   &authKey,
	}
	resp, r, err := client.AuthApi.Auth(context.Background()).AuthRequest(authRequest).Execute()
	if err != nil {
		fmt.Println(err)
	}
	if r.StatusCode != 200 {
		fmt.Println(r)
	}
	fmt.Println(*resp.ApiKey, *resp.Token)
}
```

### API call example

Here's an example to list subscribers with known API keys

```go
package main

import (
	"context"
	"fmt"
	"os"

	soracom "github.com/soracom-labs/openapi-client-go/openapi/api"
)


func main() {
	apiKey := os.Getenv("SORACOM_API_KEY")
	token := os.Getenv("SORACOM_API_TOKEN")
	config := soracom.NewConfiguration()
	client := soracom.NewAPIClient(config)

	contextApiKeys := map[string]soracom.APIKey{
		"api_key": {
			Key: apiKey,
		},
		"api_token": {
			Key: token,
		},
	}
	ctx := context.WithValue(context.Background(), soracom.ContextAPIKeys, contextApiKeys)
	resp, r, err := client.SubscriberApi.ListSubscribers(ctx).Execute()
	if err != nil {
		fmt.Println(err)
	}
	if r.StatusCode != 200 {
		fmt.Println(r)
	}
	for _, subscriber := range resp {
		fmt.Println(subscriber.GetImsi(), subscriber.GetSimId())
	}
}
```
