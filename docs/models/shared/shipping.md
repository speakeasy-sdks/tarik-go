# Shipping

Represents info about shipping order


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Address`                                                          | [NPPackStation](../../models/shared/nppackstation.md)              | :heavy_check_mark:                                                 | Represents info about shipping via `Nova Posta PackStation` method |
| `Method`                                                           | [ShippingMethod](../../models/shared/shippingmethod.md)            | :heavy_check_mark:                                                 | N/A                                                                |
| `Price`                                                            | **float64*                                                         | :heavy_minus_sign:                                                 | N/A                                                                |
| `Recipient`                                                        | [Recipient](../../models/shared/recipient.md)                      | :heavy_check_mark:                                                 | Represents info about recipient                                    |