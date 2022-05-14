
# Google Ads API Client Library for Golang

This project is a client library for the [Google Ads API](https://developers.google.com/google-ads/api/docs/start). Most of the code is generated from [googleapis](https://github.com/googleapis/googleapis/tree/master/google/ads/googleads). And this project only adds some examples and minor adjustments. Please note this is an unofficial project.


## Features

- Fully support for Google Ads API.
- Support for GRPC calls.
- Support for HTTP calls by using [protojson](https://google.golang.org/protobuf/encoding/protojson).
- Frequent updates based on [official repository](https://github.com/googleapis/googleapis).


## Requirements

- Golang 1.4 or later.
- Before starting, recommend to read the [OAuth2 guide](https://developers.google.com/google-ads/api/docs/oauth/overview).
- Apply for a [developer token](https://developers.google.com/google-ads/api/docs/first-call/dev-token) if you don't already have one.

## Installation

```
go get -d github.com/shenzhencenter/google-ads-pb
```
    
## Getting started

1. Set the environment variables.

```
export ACCESS_TOKEN=<your access token>
export DEVELOPER_TOKEN=<your developer token>
export CUSTOMER_ID=<your customer id>
```

2. Create a GRPC connection.

```
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

```
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

4. Make HTTP calls with using [protojson](https://google.golang.org/protobuf/encoding/protojson).

```
// 4.1 Set request
listAccessibleCustomersRequest, err := protojson.Marshal(
  &services.ListAccessibleCustomersRequest{},
)
if err != nil { panic(err) }
request, err := http.NewRequest(
  "GET", "https://googleads.googleapis.com/v10/customers:listAccessibleCustomers",
  bytes.NewBuffer(listAccessibleCustomersRequest))
if err != nil { panic(err) }

// 4.2 Set request headers
header := make(http.Header)
header.Set("authorization", os.Getenv("ACCESS_TOKEN"))
header.Set("developer-token", os.Getenv("DEVELOPER_TOKEN"))
header.Set("login-customer-id", os.Getenv("CUSTOMER_ID"))
request.Header = header

// 4.3 Send request
client := &http.Client{}
response, err := client.Do(request)
if err != nil { panic(err) }
defer response.Body.Close()

// 4.4 Read the response
var responseBody []byte
if responseBody, err = ioutil.ReadAll(response.Body); err != nil {
  panic(err)
}

// 4.5 Print the response
listAccessibleCustomersResponse := new(services.ListAccessibleCustomersResponse)
if err := protojson.Unmarshal(responseBody, listAccessibleCustomersResponse); err != nil {
  panic(err)
}
for _, customer := range listAccessibleCustomersResponse.ResourceNames {
  fmt.Println("ResourceName: " + customer)
}
```


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
