
# Google Ads API Client Library for Golang

[![Go](https://github.com/shenzhencenter/google-ads-pb/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/shenzhencenter/google-ads-pb/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/shenzhencenter/google-ads-pb?status.svg)](https://pkg.go.dev/github.com/shenzhencenter/google-ads-pb)
[![Go Report Card](https://goreportcard.com/badge/github.com/shenzhencenter/google-ads-pb)](https://goreportcard.com/report/github.com/shenzhencenter/google-ads-pb)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This library provides a Golang client for the [Google Ads API](https://developers.google.com/google-ads/api/docs/start). It's fully generated from the [googleapis](https://github.com/googleapis/googleapis/tree/master/google/ads/googleads) repository. More information on the generation process can be found [here](https://github.com/shenzhencenter/google-ads-pb/blob/main/.github/workflows/generator.yml).

Although this project isn't official, we deem it as low-risk due to its maturity and our many years of using it in production. However, always consult the [sunset schedule](https://developers.google.com/google-ads/api/docs/sunset-dates) of the Google Ads API.

## Features

- Full support for Google Ads API.
- Support for gRPC and REST Interface.
- Source code generated from the official googleapis repository.

## Version support

| google-ads-pb      | Google Ads API   | Sunset date                  |
| ------------------ | ---------------- | ---------------------------- |
| v1.18.0            | v18              | September 2025               |
| v1.17.1            | v17.1            | May 2025                     |
| v1.17.0            | v17              | May 2025                     |
| <del>v1.16.1</del> | <del>v16.1</del> | Deprecated                   |
| <del>v1.7.0</del>  | <del>v16</del>   | Deprecated                   |

## Requirements

- Go 1.22.
- Familiarize yourself with the [OAuth2 guide](https://developers.google.com/google-ads/api/docs/oauth/overview).
- If needed, obtain a [developer token](https://developers.google.com/google-ads/api/docs/first-call/dev-token).

## Installation

```bash
go get github.com/shenzhencenter/google-ads-pb
```
    
## Getting started

1. (Optional) Set your environment variables.

```bash
export ACCESS_TOKEN=<your-access-token>
export DEVELOPER_TOKEN=<your-developer-token>
export CUSTOMER_ID=<your-customer-id>
```

> Tips: For testing purposes, you can use [google/oauth2l](https://github.com/google/oauth2l) to obtain an access token.

2. gRPC example.

```go
var (
    customerID     = os.Getenv("CUSTOMER_ID")
    developerToken = os.Getenv("DEVELOPER_TOKEN")
    accessToken    = os.Getenv("ACCESS_TOKEN")
)

cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
conn, err := grpc.NewClient("googleads.googleapis.com:443", cred)
if err != nil {
    panic(err)
}
defer conn.Close()

ctx := context.Background()
headers := metadata.Pairs(
    "authorization", "Bearer "+accessToken,
    "developer-token", developerToken,
    "login-customer-id", customerID,
)
ctx = metadata.NewOutgoingContext(ctx, headers)

req := &services.SearchGoogleAdsRequest{
    CustomerId: customerID,
    Query: "SELECT user_list.name, user_list.resource_name FROM user_list",
}
svc := services.NewGoogleAdsServiceClient(conn)
var header metadata.MD
result, err := svc.Search(ctx, req, grpc.Header(&header))
requestId := header.Get("request-id")
log.Printf("request-id: %v, login-customer-id: %v", requestId[0], customerID)
if err != nil {
    apiErr := status.Convert(err)
    log.Fatalf("code: %s, message: %s, details: %v", apiErr.Code(), apiErr.Message(), apiErr.Details())
}

for _, row := range result.Results {
    if row.UserList == nil {
        continue
    }
    log.Print("resource_name: ", row.UserList.GetResourceName(), " name: ", row.UserList.GetName())
}
```

3. HTTP example.

```go
var (
    customerID     = os.Getenv("CUSTOMER_ID")
    developerToken = os.Getenv("DEVELOPER_TOKEN")
    accessToken    = os.Getenv("ACCESS_TOKEN")
)

var endpoint = fmt.Sprintf("https://googleads.googleapis.com/v18/customers/%s/googleAds:search", customerID)
req := services.SearchGoogleAdsRequest{
    Query: "SELECT user_list.name, user_list.resource_name FROM user_list",
}
reqBody, _ := protojson.Marshal(&req)

reqHttp, _ := http.NewRequest("POST", endpoint, bytes.NewReader(reqBody))
header := make(http.Header)
header.Set("content-type", "application/json")
header.Set("authorization", "Bearer "+accessToken)
header.Set("developer-token", developerToken)
header.Set("login-customer-id", customerID)
reqHttp.Header = header

httpResp, err := http.DefaultClient.Do(reqHttp)
if os.Getenv("DEBUG") != "" {
    reqId := httpResp.Header.Get("request-id")
    log.Print("request-id: ", reqId, " login-customer-id: ", reqHttp.Header.Get("login-customer-id"))
}
if err != nil {
    log.Panic(err)
}
defer httpResp.Body.Close()
var httpRespBody []byte
if httpRespBody, err = io.ReadAll(httpResp.Body); err != nil {
    log.Panic(err)
}
result := new(services.SearchGoogleAdsResponse)
if err := protojson.Unmarshal(httpRespBody, result); err != nil {
    log.Fatal("failed to unmarshal response: ", err, "response: ", string(httpRespBody))
}
for _, row := range result.Results {
    if row.UserList == nil {
        continue
    }
    log.Print("resource_name: ", row.UserList.GetResourceName(), " name: ", row.UserList.GetName())
}
```

## TODO

[ ] [clients/internal/snippets](https://github.com/shenzhencenter/google-ads-pb/tree/main/clients/internal/snippets).

## References

- [Google APIs](https://github.com/googleapis/googleapis)
- [Google Ads API Client Library for PHP](https://github.com/googleads/google-ads-php)
- [google/oauth2l](https://github.com/google/oauth2l)
- [How to Use a Service Account to Call the Google Ads API with Golang](https://www.liaozhen.cn/posts/how-to-use-service-account-call-google-ads-api/)

## Contributing

As this project is fully protoc-generated, we aren't accepting pull requests. However, we value your feedback and suggestions, so feel free to open an issue.
