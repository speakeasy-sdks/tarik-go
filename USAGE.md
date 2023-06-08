<!-- Start SDK Example Usage -->
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
        ClientID: lk.String("corrupti"),
        ClientSecret: lk.String("provident"),
        GrantType: lk.String("distinctio"),
        Password: "quibusdam",
        Scope: lk.String("unde"),
        Username: "Ryan.Little62",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoginAPIV1TokenPost200ApplicationJSONAny != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->