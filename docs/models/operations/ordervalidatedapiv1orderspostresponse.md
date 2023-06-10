# OrderValidatedAPIV1OrdersPostResponse


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ContentType`                                                             | *string*                                                                  | :heavy_check_mark:                                                        | N/A                                                                       |
| `HTTPValidationError`                                                     | [*shared.HTTPValidationError](../../models/shared/httpvalidationerror.md) | :heavy_minus_sign:                                                        | Validation Error                                                          |
| `Response`                                                                | [*shared.Response](../../models/shared/response.md)                       | :heavy_minus_sign:                                                        | Successful Response                                                       |
| `StatusCode`                                                              | *int*                                                                     | :heavy_check_mark:                                                        | N/A                                                                       |
| `RawResponse`                                                             | [*http.Response](https://pkg.go.dev/net/http#Response)                    | :heavy_minus_sign:                                                        | N/A                                                                       |