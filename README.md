
# Google Ads API Client Library for Golang

[![Go](https://github.com/Optable/google-ads-pb/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/Optable/google-ads-pb/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/Optable/google-ads-pb?status.svg)](https://pkg.go.dev/github.com/Optable/google-ads-pb)
[![Go Report Card](https://goreportcard.com/badge/github.com/Optable/google-ads-pb)](https://goreportcard.com/report/github.com/Optable/google-ads-pb)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This library provides a Golang client for the [Google Ads API](https://developers.google.com/google-ads/api/docs/start). It's fully generated from the [googleapis](https://github.com/googleapis/googleapis/tree/master/google/ads/googleads) repository. More information on the generation process can be found [here](https://github.com/Optable/google-ads-pb/blob/main/Makefile).

Although this project isn't official, we deem it as low-risk due to its maturity and our two years of using it in production. However, always consult the [sunset schedule](https://developers.google.com/google-ads/api/docs/sunset-dates) of the Google Ads API.

## Features

- Full support for Google Ads API.
- Source code generated from the official googleapis repository.
- Support for GRPC and HTTP calls using [protojson](https://google.golang.org/protobuf/encoding/protojson).
- Although we won't regularly update this based on the [official repository](https://github.com/googleapis/googleapis), we continually maintain it. Update frequency depends on our needs but usually occurs a month after the official release.

## Version support



| google-ads-pb     | Google Ads API   | Sunset date                  |
| ----------------- | ---------------- | ---------------------------- |
| Not generated yet            | v18              | End of September 2025       |
| Not generated yet            | v17.1            | End of May 2025             |
| Not generated yet            | v17              | End of May 2025             |
| Not generated yet            | v16.1            | End of January 2025         |
| Not generated yet            | v16              | End of January 2025         |
| Not generated yet            | v15              | End of September 2024       |
| v1.0.0            | v14              | End of May 2024       |

## Requirements

- Go 1.21.
- Familiarize yourself with the [OAuth2 guide](https://developers.google.com/google-ads/api/docs/oauth/overview).
- If needed, obtain a [developer token](https://developers.google.com/google-ads/api/docs/first-call/dev-token).

## Installation

```bash
$ go get github.com/Optable/google-ads-pb
```

## Getting started

1. Set your environment variables.

```bash
$ export ACCESS_TOKEN=<your access token>
$ export DEVELOPER_TOKEN=<your developer token>
$ export CUSTOMER_ID=<your customer id>
```

2. Establish a GRPC connection.

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

3. Start making calls.

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

You can also make HTTP calls using [protojson](https://google.golang.org/protobuf/encoding/protojson), though it isn't recommended.

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

See [clients/internal/snippets](https://github.com/Optable/google-ads-pb/tree/main/clients/internal/snippets).

## Note for the updating process

1. Find the version number in [this folder](https://github.com/googleapis/googleapis/tree/master/google/ads/googleads).
2. Checkout the branch with the version number. Example: For v14, checkout the branch `v14`.
3. Execute the following command.

```bash
make generate
```
4. It will prompt you to enter the version number. Enter the version number you found in step 1. Example: For v14, enter 14.
5. Once the process is done, it should have modified some files in `/clients` and `/protogen`
6. Modify this file to add the version number to the table with the next tag. Please also update the table with the infos here: https://developers.google.com/google-ads/api/docs/sunset-dates
7. Create a PR with the changes, merge it and create a new tag with the version number added in step 6.
```bash
git tag v1.0.0
```
8. Push the tag
```bash
git push origin v1.0.0
```
9. Create a new release on Github with the tag created in step 7.
10. Publish the realease on Golang
```bash
GOPROXY=proxy.golang.org go list -m github.com/Optable/google-ads-pb@v1.0.0
```

## Related

Here are some related projects

- [Google APIs](https://github.com/googleapis/googleapis)
- [Google Ads API Client Library for PHP](https://github.com/googleads/google-ads-php)

## Contributing

As this project is fully protoc-generated, we aren't accepting pull requests. However, we value your feedback and suggestions, so feel free to open an issue.
