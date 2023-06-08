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
