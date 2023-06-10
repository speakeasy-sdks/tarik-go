# OrderTicket


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Decision`                                                    | *string*                                                      | :heavy_check_mark:                                            | Decision                                                      |
| `OrderID`                                                     | *int64*                                                       | :heavy_check_mark:                                            | Marketplace Order ID                                          |
| `SolutionTime`                                                | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | Format YYYY-MM-DD[T]HH:MM                                     |
| `Status`                                                      | [OrderTicketStatus](../../models/shared/orderticketstatus.md) | :heavy_check_mark:                                            | Status                                                        |
| `TicketID`                                                    | *string*                                                      | :heavy_check_mark:                                            | BPM ticket ID                                                 |