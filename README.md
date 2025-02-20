
# Google Ads API Client Library for Golang

[![Go](https://github.com/shenzhencenter/google-ads-pb/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/shenzhencenter/google-ads-pb/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/shenzhencenter/google-ads-pb?status.svg)](https://pkg.go.dev/github.com/shenzhencenter/google-ads-pb)
[![Go Report Card](https://goreportcard.com/badge/github.com/shenzhencenter/google-ads-pb)](https://goreportcard.com/report/github.com/shenzhencenter/google-ads-pb)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This library provides a Golang client for the [Google Ads API](https://developers.google.com/google-ads/api/docs/start). It's fully generated from the [googleapis](https://github.com/googleapis/googleapis/tree/master/google/ads/googleads) repository. More information on the generation process can be found [here](https://github.com/shenzhencenter/google-ads-pb/blob/main/.github/workflows/generator.yml).

Although this project isn't official, we deem it as low-risk due to its maturity and our two years of using it in production. However, always consult the [sunset schedule](https://developers.google.com/google-ads/api/docs/sunset-dates) of the Google Ads API.

## Features

- Full support for Google Ads API.
- Source code generated from the official googleapis repository.
- Support for GRPC and HTTP calls using [protojson](https://google.golang.org/protobuf/encoding/protojson).
- Although we won't regularly update this based on the [official repository](https://github.com/googleapis/googleapis), we continually maintain it. Update frequency depends on our needs but usually occurs a month after the official release.

## Version support

| google-ads-pb     | Google Ads API   | Sunset date                  |
| ----------------- | ---------------- | ---------------------------- |
| v1.18.0           | v18              | September 2025               |
| v1.17.1           | v17.1            | May 2025                     |
| v1.17.0           | v17              | May 2025                     |
| v1.16.1           | v16.1            | January 2025                 |
| v1.7.0            | v16              | January 2025                 |
| <del>v1.6.0</del> | <del>v15</del>   | Deprecated                   |
| <del>v1.5.1</del> | <del>v14.1</del> | Deprecated                   |
| <del>v1.5.0</del> | <del>v14</del>   | Deprecated                   |

## Requirements

- Go 1.22.
- Familiarize yourself with the [OAuth2 guide](https://developers.google.com/google-ads/api/docs/oauth/overview).
- If needed, obtain a [developer token](https://developers.google.com/google-ads/api/docs/first-call/dev-token).

## Installation

```bash
go get github.com/shenzhencenter/google-ads-pb
```
    
## Getting started

1. Set your environment variables.

```bash
export ACCESS_TOKEN=<your access token>
export DEVELOPER_TOKEN=<your developer token>
export CUSTOMER_ID=<your customer id>
```

If you're using gRPC, you should attach the access token, developer token, and customer ID to the context.

```go
ctx := context.Background()
headers := metadata.Pairs(
    "authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"),
    "developer-token", os.Getenv("DEVELOPER_TOKEN"),
    "login-customer-id", os.Getenv("CUSTOMER_ID"),
)
ctx = metadata.NewOutgoingContext(ctx, headers)
```

If you're using HTTP, you should attach the access token, developer token, and customer ID to the header.

```go
header := make(http.Header)
header.Set("content-type", "application/json")
header.Set("authorization", os.Getenv("ACCESS_TOKEN"))
header.Set("developer-token", os.Getenv("DEVELOPER_TOKEN"))
```

2. Establish a GRPC connection.

```go
cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
conn, err := grpc.NewClient("googleads.googleapis.com:443", cred)
if err != nil {
    panic(err)
}
defer conn.Close()
```

3. Start making calls.

```go
customerServiceClient := services.NewCustomerServiceClient(conn)
accessibleCustomers, err := customerServiceClient.ListAccessibleCustomers(
    ctx, // be sure to use the context with the access token, developer token, and customer ID
    &services.ListAccessibleCustomersRequest{},
)
if err != nil {
    panic(err)
}

for _, customer := range accessibleCustomers.ResourceNames {
    fmt.Println("ResourceName: " + customer)
}
```

You can also make HTTP calls using [protojson](https://google.golang.org/protobuf/encoding/protojson), though it isn't recommended.

```go
const endpoint = "https://googleads.googleapis.com/v18/customers:listAccessibleCustomers"
req := services.ListAccessibleCustomersRequest{}
requestBody, _ := protojson.Marshal(&req)
request, _ := http.NewRequest("GET", endpoint, bytes.NewBuffer(requestBody))
header := make(http.Header)
header.Set("content-type", "application/json")
header.Set("authorization", os.Getenv("ACCESS_TOKEN"))
header.Set("developer-token", os.Getenv("DEVELOPER_TOKEN"))
request.Header = header
response, _ := http.DefaultClient.Do(request)
defer response.Body.Close()
var responseBody []byte
if responseBody, err = io.ReadAll(response.Body); err != nil {
    panic(err)
}
listAccessibleCustomersResponse := new(services.ListAccessibleCustomersResponse)
if err := protojson.Unmarshal(responseBody, listAccessibleCustomersResponse); err != nil {
    panic(err)
}
for _, customer := range listAccessibleCustomersResponse.ResourceNames {
    fmt.Println("ResourceName: " + customer)
}
```

> **Note**: The above examples are just a starting point. You should adjust them based on your needs.

## Examples

See [clients/internal/snippets](https://github.com/shenzhencenter/google-ads-pb/tree/main/clients/internal/snippets).

## Related

Here are some related projects

- [Google APIs](https://github.com/googleapis/googleapis)
- [Google Ads API Client Library for PHP](https://github.com/googleads/google-ads-php)

## Contributing

As this project is fully protoc-generated, we aren't accepting pull requests. However, we value your feedback and suggestions, so feel free to open an issue.
