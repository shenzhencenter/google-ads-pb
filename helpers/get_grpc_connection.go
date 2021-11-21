package helpers

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const EndPoint = "googleads.googleapis.com:443"

func NewGrpcConnection(ctx context.Context, token *oauth2.Token, developerToken string, customerID string) (*grpc.ClientConn, context.Context, error) {
	headers := metadata.Pairs(
		"Authorization", fmt.Sprintf("%s %s", token.TokenType, token.AccessToken),
		"developer-token", developerToken,
	)

	if len(customerID) > 0 {
		headers.Set("login-customer-id", customerID)
	}

	ctx = metadata.NewOutgoingContext(ctx, headers)

	cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	conn, err := grpc.Dial(EndPoint, cred)
	if err != nil {
		return nil, nil, err
	}
	return conn, ctx, nil
}

func ResetLoginCustomerID(ctx context.Context, customerID string) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, "login-customer-id", customerID)
	return ctx
}
