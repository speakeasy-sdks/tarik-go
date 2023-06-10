# LK

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/tarik-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Authentication](docs/sdks/authentication/README.md)

* [LoginAPIV1TokenPost](docs/sdks/authentication/README.md#loginapiv1tokenpost) - Login

### [Order](docs/sdks/order/README.md)

* [OrderSchemaAPIV1OrdersSchemaGet](docs/sdks/order/README.md#orderschemaapiv1ordersschemaget) - Get JSON schema for order
* [OrderValidatedAPIV1OrdersPost](docs/sdks/order/README.md#ordervalidatedapiv1orderspost) - Add new order

### [OrderStatus](docs/sdks/orderstatus/README.md)

* [StatusRequestValidatedAPIV1OrdersOrderIDStatusRequestPost](docs/sdks/orderstatus/README.md#statusrequestvalidatedapiv1ordersorderidstatusrequestpost) - Send order status request
* [StatusSchemaAPIV1OrdersStatusSchemaGet](docs/sdks/orderstatus/README.md#statusschemaapiv1ordersstatusschemaget) - Get JSON schema for order status
* [StatusValidatedAPIV1OrdersOrderIDStatusPost](docs/sdks/orderstatus/README.md#statusvalidatedapiv1ordersorderidstatuspost) - Set order status

### [OrderTicket](docs/sdks/orderticket/README.md)

* [NewTicketAPIV1TicketPost](docs/sdks/orderticket/README.md#newticketapiv1ticketpost) - Add new order ticket
* [TicketStatusAPIV1TicketTicketIDCommentPost](docs/sdks/orderticket/README.md#ticketstatusapiv1ticketticketidcommentpost) - Add ticket comment
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
