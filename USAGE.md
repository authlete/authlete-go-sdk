<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	authlete "github.com/authlete/authlete-go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := authlete.New(
		authlete.WithSecurity(os.Getenv("AUTHLETE_BEARER")),
	)

	res, err := s.Service.Get(ctx, "<id>")
	if err != nil {
		log.Fatal(err)
	}
	if res.Service != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->