# Authentication

### Available Operations

* [LoginAPIV1TokenPost](#loginapiv1tokenpost) - Login

## LoginAPIV1TokenPost

Login

### Example Usage

```go
package main

import(
	"context"
	"log"
	"LK"
	"LK/pkg/models/shared"
)

func main() {
    s := lk.New()

    ctx := context.Background()
    res, err := s.Authentication.LoginAPIV1TokenPost(ctx, shared.BodyLoginAPIV1TokenPost{
        ClientID: lk.String("deserunt"),
        ClientSecret: lk.String("suscipit"),
        GrantType: lk.String("iure"),
        Password: "magnam",
        Scope: lk.String("debitis"),
        Username: "Anahi38",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoginAPIV1TokenPost200ApplicationJSONAny != nil {
        // handle response
    }
}
```
