<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	authlete "github.com/authlete/authlete-go-sdk"
	"github.com/authlete/authlete-go-sdk/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := authlete.New(
		authlete.WithServiceID("<id>"),
		authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
	)

	res, err := s.Service.Get(ctx, operations.ServiceGetAPIRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Service != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->