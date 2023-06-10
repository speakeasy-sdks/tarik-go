# Order

Represents full info about order


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `CreatedAt`                                 | [time.Time](https://pkg.go.dev/time#Time)   | :heavy_check_mark:                          | Format YYYY-MM-DD[T]HH:MM                   |
| `Customer`                                  | [Customer](../../models/shared/customer.md) | :heavy_check_mark:                          | Represents info about order customer        |
| `Items`                                     | [][Item](../../models/shared/item.md)       | :heavy_check_mark:                          | N/A                                         |
| `MerchantID`                                | *int64*                                     | :heavy_check_mark:                          | This field could be numeric string          |
| `OrderID`                                   | *int64*                                     | :heavy_check_mark:                          | This field could be numeric string          |
| `Payment`                                   | [Payment](../../models/shared/payment.md)   | :heavy_check_mark:                          | Represents info about order payment         |
| `Shipping`                                  | [Shipping](../../models/shared/shipping.md) | :heavy_check_mark:                          | Represents info about shipping order        |