
# Google Ads API Client Library for Golang

[![Go](https://github.com/shenzhencenter/google-ads-pb/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/shenzhencenter/google-ads-pb/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/shenzhencenter/google-ads-pb?status.svg)](https://pkg.go.dev/github.com/shenzhencenter/google-ads-pb)
[![Go Report Card](https://goreportcard.com/badge/github.com/shenzhencenter/google-ads-pb)](https://goreportcard.com/report/github.com/shenzhencenter/google-ads-pb)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is a Golang client library for accessing the [Google Ads API](https://developers.google.com/google-ads/api/docs/start). Most of the code is generated from the [googleapis](https://github.com/googleapis/googleapis/tree/master/google/ads/googleads) repository. We only make some modifications and examples. Detailed informations can be found on [workflows/generator](https://github.com/shenzhencenter/google-ads-pb/blob/main/.github/workflows/generator.yml).

Please note that this is not an official project. But we think it is very low risk because it is a very mature project, and we have been using it in production for more than two years. The only thing to note is the [sunset schedule](https://developers.google.com/google-ads/api/docs/sunset-dates) of the Google Ads API.

## Features

- Fully support for Google Ads API.
- Support for GRPC calls.
- Support for HTTP calls by using [protojson](https://google.golang.org/protobuf/encoding/protojson).
- Frequent updates based on [official repository](https://github.com/googleapis/googleapis).

## Version support

| google-ads-pb | Google Ads API | Sunset date |
|---|---|---|
| v1.5.0 | v14 | End of May 2024 |
| v1.4.1 | v13.1 | End of January 2024 |
| v1.4.0 | v13 | End of January 2024 | 
| v1.3.1 | v12 | <b>End of September 2023</b> |
| <del>v1.2.1</del> | <del>v11.1</del> | Deprecated |
| <del>v1.2.0</del> | <del>v11</del> | Deprecated |
| <del>v1.1.1</del> | <del>v10</del> | Deprecated |

## Requirements

- Go 20 or later.
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
if err != nil { panic(err) }
defer conn.Close()
```

3. Make the first call.

```go
customerServiceClient := services.NewCustomerServiceClient(conn)
accessibleCustomers, err := customerServiceClient.ListAccessibleCustomers(
  ctx, 
  &services.ListAccessibleCustomersRequest{},
)
if err != nil { panic(err) }

for _, customer := range accessibleCustomers.ResourceNames {
  fmt.Println("ResourceName: " + customer)
}
```

4. Make HTTP calls with using [protojson](https://google.golang.org/protobuf/encoding/protojson) is not recommended now. But it is available yet, for more information please refer to the code in examples.


## Examples


### Account management

- ListAccessiableCustomer [[GRPC](https://github.com/shenzhencenter/google-ads-pb/blob/main/examples/account_management/list_accessible_customers.go)] [[HTTP](https://github.com/shenzhencenter/google-ads-pb/blob/main/examples/account_management/http_list_accessible_customers.go)]
- GetAccountInformation [[GRPC](https://github.com/shenzhencenter/google-ads-pb/blob/main/examples/account_management/get_account_information.go)] [[HTTP](https://github.com/shenzhencenter/google-ads-pb/blob/main/examples/account_management/http_get_account_information.go)]

### Campaign management

- GetCampaignsByLabel [[GRPC](https://github.com/shenzhencenter/google-ads-pb/blob/main/examples/campaign_management/get_campaigns_by_label.go)] [[HTTP](https://github.com/shenzhencenter/google-ads-pb/blob/main/examples/campaign_management/http_get_campaigns_by_label.go)]

## Related

Here are some related projects

- [Google APIs](https://github.com/googleapis/googleapis)
- [Google Ads API Client Library for PHP](https://github.com/googleads/google-ads-php)

## Contributing

Welcome to contribute more examples and documentations.

### Supporting

If you like this project, please consider buying me a coffee. 😄

<a href="https://www.buymeacoffee.com/crazyfrog" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>
