# google-ads-pb

You can use the golang library to interact with the Google Ads API across grpc. This library is not the official Google Ads API library.

> Generated from [googleapis/ads/v9](https://github.com/googleapis/googleapis/tree/master/google/ads/googleads/v9)

## Installation

```
$ go get -d github.com/shenzhencenter/google-ads-pb
```

## Usage

main.go contains some example functions. It works with environment variables:

```
export GOOGLE_CLIENT_ID=<your client id>
export GOOGLE_CLIENT_SECRET=<your client secret>
export GOOGLE_DEVELOPER_TOKEN=<your developer token>
export GOOGLE_ACCESS_TOKEN=<your access token>
export GOOGLE_REFRESH_TOKEN=<your refresh token>
```

You can run the example functions with:

```
$ go run main.go
```

It will output the result to stdout:

```
# output example 1: list accessible Google Ads accounts
ResourceName: customers/1231231234

# output example 2: get a Google Ads account information
test-google-account-name

# output example 3: search all campaigns
customers/1231234/campaigns/12341234123
```

See [main.go](https://github.com/shenzhencenter/google-ads-pb/blob/main/main.go) and [examples](https://github.com/shenzhencenter/google-ads-pb/tree/main/examples) for more details.

## Documentation

Documentation is in progress. Welcome to contribute!
