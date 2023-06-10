# Order

### Available Operations

* [OrderSchemaAPIV1OrdersSchemaGet](#orderschemaapiv1ordersschemaget) - Get JSON schema for order
* [OrderValidatedAPIV1OrdersPost](#ordervalidatedapiv1orderspost) - Add new order

## OrderSchemaAPIV1OrdersSchemaGet

Get JSON schema for order

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
    res, err := s.Order.OrderSchemaAPIV1OrdersSchemaGet(ctx)
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

**[*operations.OrderSchemaAPIV1OrdersSchemaGetResponse](../../models/operations/orderschemaapiv1ordersschemagetresponse.md), error**


## OrderValidatedAPIV1OrdersPost

Checks if JSON has valid schema and adds request to create new order. No multiple orders will be created for the same **order_id**, even if request is accepted. Only first order request for **order_id** is created.

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
    res, err := s.Order.OrderValidatedAPIV1OrdersPost(ctx, shared.Order{
        CreatedAt: types.MustTimeFromString("2022-03-18T00:29:19.137Z"),
        Customer: shared.Customer{
            Email: "Junior.Kshlerin@hotmail.com",
            ErpID: 925597,
            FirstName: "Rocky",
            HumanID: 71036,
            ID: 337396,
            LastName: "Bogisich",
            MiddleName: "deserunt",
            Phone: "489-818-8947",
        },
        Items: []shared.Item{
            shared.Item{
                Name: "Deanna Sauer MD",
                Price: 6399.21,
                Quantity: 582020,
                Sku: 143353,
            },
            shared.Item{
                Name: "Irvin Rosenbaum III",
                Price: 4736,
                Quantity: 264555,
                Sku: 186332,
            },
            shared.Item{
                Name: "Jonathon Klocko",
                Price: 1352.18,
                Quantity: 18789,
                Sku: 324141,
            },
            shared.Item{
                Name: "Louis Moore",
                Price: 3864.89,
                Quantity: 943749,
                Sku: 902599,
            },
        },
        MerchantID: 681820,
        OrderID: 449950,
        Payment: shared.Payment{
            Method: shared.PaymentMethodZero,
            Status: shared.PaymentStatusZero,
        },
        Shipping: shared.Shipping{
            Address: shared.NPPackStation{
                City: "Kertzmannside",
                CityID: "b10faaa2-352c-4595-9907-aff1a3a2fa94",
                Comment: "commodi",
                Region: "quam",
                RegionID: "739251aa-52c3-4f5a-9019-da1ffe78f097",
                SettlementDescription: "cum",
                SettlementType: "0074f154-71b5-4e6e-93b9-9d488e1e91e4",
                Street: shared.Street{
                    Name: "Elizabeth Orn",
                },
                WarehouseID: "abd44269-802d-4502-a94b-b4f63c969e9a",
                WarehouseNumber: 223081,
            },
            Method: shared.ShippingMethodNovaposhtaPackstation,
            Price: lk.Float64(8915.55),
            Recipient: shared.Recipient{
                FirstName: "Veda",
                LastName: "Parisian",
                MiddleName: "in",
                Phone: "896.227.8436 x825",
            },
        },
    }, operations.OrderValidatedAPIV1OrdersPostSecurity{
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.Order](../../models/shared/order.md)                                                                         | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.OrderValidatedAPIV1OrdersPostSecurity](../../models/operations/ordervalidatedapiv1orderspostsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |


### Response

**[*operations.OrderValidatedAPIV1OrdersPostResponse](../../models/operations/ordervalidatedapiv1orderspostresponse.md), error**

