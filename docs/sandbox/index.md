# SORACOM API Sandbox

## Example

### Initialize Sandbox Operator

Here's an example to describe how to initialize Sandbox operator.

```go
package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/soracom-labs/openapi-client-go/openapi/sandbox"
)


func main() {
	authKeyId := os.Getenv("SORACOM_AUTH_KEY_ID")
	authKey := os.Getenv("SORACOM_AUTH_KEY")

	randStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	sandboxInitRequest := *sandbox.NewSandboxInitRequest(authKey, authKeyId, fmt.Sprintf("openapi-client-go-test+%s@soracom.jp", randStr), fmt.Sprintf("Password%s", randStr))
	sandboxInitRequest.CoverageTypes = append(sandboxInitRequest.CoverageTypes, "g", "jp")
	configuration := sandbox.NewConfiguration()
	client := sandbox.NewAPIClient(configuration)
	resp, r, err := client.OperatorApi.SandboxInitializeOperator(context.Background()).SandboxInitRequest(sandboxInitRequest).Execute()
	if err != nil {
		fmt.Println(err)
	}
	if r.StatusCode != 200 {
		fmt.Println(r)
	}
	fmt.Println(*resp.ApiKey, *resp.Token)
}
```
