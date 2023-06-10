# OrderStatus

### Available Operations

* [StatusRequestValidatedAPIV1OrdersOrderIDStatusRequestPost](#statusrequestvalidatedapiv1ordersorderidstatusrequestpost) - Send order status request
* [StatusSchemaAPIV1OrdersStatusSchemaGet](#statusschemaapiv1ordersstatusschemaget) - Get JSON schema for order status
* [StatusValidatedAPIV1OrdersOrderIDStatusPost](#statusvalidatedapiv1ordersorderidstatuspost) - Set order status

## StatusRequestValidatedAPIV1OrdersOrderIDStatusRequestPost

Send order status request

### Example Usage

```go
package main

import(
	"context"
	"log"
	"LK"
	"LK/pkg/models/operations"
	"LK/pkg/models/shared"
)

func main() {
    s := lk.New()

    ctx := context.Background()
    res, err := s.OrderStatus.StatusRequestValidatedAPIV1OrdersOrderIDStatusRequestPost(ctx, operations.StatusRequestValidatedAPIV1OrdersOrderIDStatusRequestPostRequest{
        OrderStatus: shared.OrderStatus{
            OrderID: 313218,
            Status: shared.OrderStatusesOne,
        },
        OrderID: 965417,
    }, operations.StatusRequestValidatedAPIV1OrdersOrderIDStatusRequestPostSecurity{
        OAuth2PasswordBearer: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.StatusRequestValidatedAPIV1OrdersOrderIDStatusRequestPostRequest](../../models/operations/statusrequestvalidatedapiv1ordersorderidstatusrequestpostrequest.md)   | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |
| `security`                                                                                                                                                                   | [operations.StatusRequestValidatedAPIV1OrdersOrderIDStatusRequestPostSecurity](../../models/operations/statusrequestvalidatedapiv1ordersorderidstatusrequestpostsecurity.md) | :heavy_check_mark:                                                                                                                                                           | The security requirements to use for the request.                                                                                                                            |


### Response

**[*operations.StatusRequestValidatedAPIV1OrdersOrderIDStatusRequestPostResponse](../../models/operations/statusrequestvalidatedapiv1ordersorderidstatusrequestpostresponse.md), error**


## StatusSchemaAPIV1OrdersStatusSchemaGet

Get JSON schema for order status

### Example Usage

```go
package main

import(
	"context"
	"log"
	"LK"
)

func main() {
    s := lk.New()

    ctx := context.Background()
    res, err := s.OrderStatus.StatusSchemaAPIV1OrdersStatusSchemaGet(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.StatusSchemaAPIV1OrdersStatusSchemaGetResponse](../../models/operations/statusschemaapiv1ordersstatusschemagetresponse.md), error**


## StatusValidatedAPIV1OrdersOrderIDStatusPost

Set order status

### Example Usage

```go
package main

import(
	"context"
	"log"
	"LK"
	"LK/pkg/models/operations"
	"LK/pkg/models/shared"
)

func main() {
    s := lk.New()

    ctx := context.Background()
    res, err := s.OrderStatus.StatusValidatedAPIV1OrdersOrderIDStatusPost(ctx, operations.StatusValidatedAPIV1OrdersOrderIDStatusPostRequest{
        OrderStatus: shared.OrderStatus{
            OrderID: 692532,
            Status: shared.OrderStatusesOne,
        },
        OrderID: 725255,
    }, operations.StatusValidatedAPIV1OrdersOrderIDStatusPostSecurity{
        OAuth2PasswordBearer: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Response != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `request`                                                                                                                                        | [operations.StatusValidatedAPIV1OrdersOrderIDStatusPostRequest](../../models/operations/statusvalidatedapiv1ordersorderidstatuspostrequest.md)   | :heavy_check_mark:                                                                                                                               | The request object to use for the request.                                                                                                       |
| `security`                                                                                                                                       | [operations.StatusValidatedAPIV1OrdersOrderIDStatusPostSecurity](../../models/operations/statusvalidatedapiv1ordersorderidstatuspostsecurity.md) | :heavy_check_mark:                                                                                                                               | The security requirements to use for the request.                                                                                                |


### Response

**[*operations.StatusValidatedAPIV1OrdersOrderIDStatusPostResponse](../../models/operations/statusvalidatedapiv1ordersorderidstatuspostresponse.md), error**

