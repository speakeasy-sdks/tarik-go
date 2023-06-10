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

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.BodyLoginAPIV1TokenPost](../../models/shared/bodyloginapiv1tokenpost.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.LoginAPIV1TokenPostResponse](../../models/operations/loginapiv1tokenpostresponse.md), error**

