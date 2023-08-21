
# Google Ads API Client Library for Golang

[![Go](https://github.com/shenzhencenter/google-ads-pb/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/shenzhencenter/google-ads-pb/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/shenzhencenter/google-ads-pb?status.svg)](https://pkg.go.dev/github.com/shenzhencenter/google-ads-pb)
[![Go Report Card](https://goreportcard.com/badge/github.com/shenzhencenter/google-ads-pb)](https://goreportcard.com/report/github.com/shenzhencenter/google-ads-pb)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is a Golang client library for accessing the [Google Ads API](https://developers.google.com/google-ads/api/docs/start). The code is fully generated from the [googleapis](https://github.com/googleapis/googleapis/tree/master/google/ads/googleads) repository. Detailed informations can be found on [workflows/generator](https://github.com/shenzhencenter/google-ads-pb/blob/main/.github/workflows/generator.yml).

Please note that this is not an official project. But we think it is very low risk because it is a very mature project, and we have been using it in production for more than two years. The only thing to note is the [sunset schedule](https://developers.google.com/google-ads/api/docs/sunset-dates) of the Google Ads API.

## Features

- Fully support for Google Ads API.
- Code is fully generated from the official googleapis repository.
- Support for GRPC calls.
- Support for HTTP calls by using [protojson](https://google.golang.org/protobuf/encoding/protojson).
- <del>Frequent updates based on [official repository](https://github.com/googleapis/googleapis).</del> We are sorry that we will not update it frequently in the future, but we will continue to maintain it. The frequency of updates depends on our needs (abount later than the official release one month).

## Version support

| google-ads-pb     | Google Ads API   | Sunset date                  |
| ----------------- | ---------------- | ---------------------------- |
| v1.5.1            | v14.1            | End of May 2024              |
| v1.5.0            | v14              | End of May 2024              |
| v1.4.1            | v13.1            | End of January 2024          |
| v1.4.0            | v13              | End of January 2024          |
| v1.3.1            | v12              | <b>End of September 2023</b> |
| <del>v1.2.1</del> | <del>v11.1</del> | Deprecated                   |
| <del>v1.2.0</del> | <del>v11</del>   | Deprecated                   |
| <del>v1.1.1</del> | <del>v10</del>   | Deprecated                   |

## Requirements

- Go 1.20 or later.
- Before starting, recommend to read the [OAuth2 guide](https://developers.google.com/google-ads/api/docs/oauth/overview).
- Apply for a [developer token](https://developers.google.com/google-ads/api/docs/first-call/dev-token) if you don't already have one.

## Installation

```bash
$ go get github.com/shenzhencenter/google-ads-pb
```
    
## Getting started

1. Set the environment variables.

```bash
$ export ACCESS_TOKEN=<your access token>
$ export DEVELOPER_TOKEN=<your developer token>
$ export CUSTOMER_ID=<your customer id>
```

2. Create a GRPC connection.

```go
ctx := context.Background()

headers := metadata.Pairs(
  "authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"),
  "developer-token", os.Getenv("DEVELOPER_TOKEN"),
  "login-customer-id", os.Getenv("CUSTOMER_ID"),
)
ctx = metadata.NewOutgoingContext(ctx, headers)

cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
conn, err := grpc.Dial("googleads.googleapis.com:443", cred)
if err != nil {
  panic(err)
}
defer conn.Close()
```

3. Make the first call.

```go
customerServiceClient := services.NewCustomerServiceClient(conn)
accessibleCustomers, err := customerServiceClient.ListAccessibleCustomers(
  ctx, 
  &services.ListAccessibleCustomersRequest{},
)
if err != nil {
  panic(err)
}

for _, customer := range accessibleCustomers.ResourceNames {
  fmt.Println("ResourceName: " + customer)
}
```

4. Make HTTP calls with using [protojson](https://google.golang.org/protobuf/encoding/protojson) is not recommended now, but it is available yet.

```go
req := services.ListAccessibleCustomersRequest{}
requestBody, err := protojson.Marshal(&req)
if err != nil {
  panic(err)
}
request, err := http.NewRequest("GET", "https://googleads.googleapis.com/v14/customers:listAccessibleCustomers", bytes.NewBuffer(requestBody))
if err != nil {
  panic(err)
}
header := make(http.Header)
header.Set("content-type", "application/json")
header.Set("authorization", os.Getenv("ACCESS_TOKEN"))
header.Set("developer-token", os.Getenv("DEVELOPER_TOKEN"))
request.Header = header
client := &http.Client{}
response, err := client.Do(request)
if err != nil {
  panic(err)
}
defer response.Body.Close()
var responseBody []byte
if responseBody, err = io.ReadAll(response.Body); err != nil {
  panic(err)
}
listAccessibleCustomersResponse := new(services.ListAccessibleCustomersResponse)
if err := protojson.Unmarshal(response, listAccessibleCustomersResponse); err != nil {
  panic(err)
}
for _, customer := range listAccessibleCustomersResponse.ResourceNames {
  fmt.Println("ResourceName: " + customer)
}
```

## Examples

See [clients/internal/snippets](https://github.com/shenzhencenter/google-ads-pb/tree/main/clients/internal/snippets).

## Related

Here are some related projects

- [Google APIs](https://github.com/googleapis/googleapis)
- [Google Ads API Client Library for PHP](https://github.com/googleads/google-ads-php)

## Contributing

This project is fully generated by protoc, so we can't accept pull requests. But we are very happy to receive your feedback and suggestions. Please feel free to create an issue.
