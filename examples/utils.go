package examples

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	EndPoint                      = "googleads.googleapis.com:443"
	ListAccessibleCustomersMehtod = "GET"
	ListAccessibleCustomersPath   = "https://googleads.googleapis.com/v11/customers:listAccessibleCustomers"
	GoogleAdsSearchMehtod         = "POST"
	GoogleAdsSearchPath           = "https://googleads.googleapis.com/v11/customers/%s/googleAds:search"
)

type (
	ContextOption func(*context.Context)
	HeaderOption  func(header *http.Header)
)

var conn *grpc.ClientConn

func GetGRPCConnection() *grpc.ClientConn {
	if conn != nil {
		return conn
	}
	cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	conn, err := grpc.Dial(EndPoint, cred)
	if err != nil {
		panic(err)
	}
	return conn
}

func WithContext(k string, v string) ContextOption {
	return func(ctx *context.Context) {
		*ctx = metadata.AppendToOutgoingContext(*ctx, k, v)
	}
}

func SetContext(ctx context.Context, opts ...ContextOption) context.Context {
	for _, opt := range opts {
		opt(&ctx)
	}
	return ctx
}

func WithHeader(k string, v string) HeaderOption {
	return func(header *http.Header) {
		header.Set(k, v)
	}
}

func HttpRequest(body []byte, method string, url string, opts ...HeaderOption) []byte {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	header := make(http.Header)
	header.Set("content-type", "application/json")
	header.Set("accept", "application/json")
	for _, opt := range opts {
		opt(&header)
	}
	request.Header = header

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 4.4 Read the response
	var responseBody []byte
	if responseBody, err = ioutil.ReadAll(response.Body); err != nil {
		panic(err)
	}
	return responseBody
}
