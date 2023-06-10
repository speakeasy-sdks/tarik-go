# OrderTicket

### Available Operations

* [NewTicketAPIV1TicketPost](#newticketapiv1ticketpost) - Add new order ticket
* [TicketStatusAPIV1TicketTicketIDCommentPost](#ticketstatusapiv1ticketticketidcommentpost) - Add ticket comment

## NewTicketAPIV1TicketPost

Add new order ticket

### Example Usage

```go
package main

import(
	"context"
	"log"
	"LK"
	"LK/pkg/models/shared"
	"LK/pkg/types"
	"LK/pkg/models/operations"
)

func main() {
    s := lk.New()

    ctx := context.Background()
    res, err := s.OrderTicket.NewTicketAPIV1TicketPost(ctx, shared.OrderTicket{
        Decision: "id",
        OrderID: 501324,
        SolutionTime: types.MustTimeFromString("2021-02-02T01:24:52.629Z"),
        Status: shared.OrderTicketStatusSevene9f1204F46b1410Fb9a0050ba5d6c38,
        TicketID: "deserunt",
    }, operations.NewTicketAPIV1TicketPostSecurity{
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [shared.OrderTicket](../../models/shared/orderticket.md)                                                   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.NewTicketAPIV1TicketPostSecurity](../../models/operations/newticketapiv1ticketpostsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |


### Response

**[*operations.NewTicketAPIV1TicketPostResponse](../../models/operations/newticketapiv1ticketpostresponse.md), error**


## TicketStatusAPIV1TicketTicketIDCommentPost

Add ticket comment

### Example Usage

```go
package main

import(
	"context"
	"log"
	"LK"
	"LK/pkg/models/operations"
	"LK/pkg/models/shared"
	"LK/pkg/types"
)

func main() {
    s := lk.New()

    ctx := context.Background()
    res, err := s.OrderTicket.TicketStatusAPIV1TicketTicketIDCommentPost(ctx, operations.TicketStatusAPIV1TicketTicketIDCommentPostRequest{
        OrderTicket: shared.OrderTicket{
            Decision: "nisi",
            OrderID: 423855,
            SolutionTime: types.MustTimeFromString("2021-10-15T07:59:26.631Z"),
            Status: shared.OrderTicketStatusThreee7f420cF46b1410Fc9a0050ba5d6c38,
            TicketID: "perferendis",
        },
        TicketID: "nihil",
    }, operations.TicketStatusAPIV1TicketTicketIDCommentPostSecurity{
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

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.TicketStatusAPIV1TicketTicketIDCommentPostRequest](../../models/operations/ticketstatusapiv1ticketticketidcommentpostrequest.md)   | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `security`                                                                                                                                     | [operations.TicketStatusAPIV1TicketTicketIDCommentPostSecurity](../../models/operations/ticketstatusapiv1ticketticketidcommentpostsecurity.md) | :heavy_check_mark:                                                                                                                             | The security requirements to use for the request.                                                                                              |


### Response

**[*operations.TicketStatusAPIV1TicketTicketIDCommentPostResponse](../../models/operations/ticketstatusapiv1ticketticketidcommentpostresponse.md), error**

