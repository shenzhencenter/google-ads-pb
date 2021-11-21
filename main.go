package main

import (
	"context"
	"github.com/caarlos0/env"
	"github.com/shenzhencenter/google-ads-pb/examples"
	"github.com/shenzhencenter/google-ads-pb/helpers"
	"time"
)

type Args struct {
	ClientID       string `env:"GOOGLE_CLIENT_ID,required"`
	ClientSecret   string `env:"GOOGLE_CLIENT_SECRET,required"`
	DeveloperToken string `env:"GOOGLE_DEVELOPER_TOKEN,required"`
	AccessToken    string `env:"GOOGLE_ACCESS_TOKEN,required"`
	TokenType      string `env:"GOOGLE_TOKEN_TYPE" envDefault:"Bearer"`
	RefreshToken   string `env:"GOOGLE_REFRESH_TOKEN,required"`
}

const (
	TestCustomerID = "1231231234"
)

func main() {
	// Load environment variables
	args := Args{}
	if err := env.Parse(&args); err != nil {
		panic(err)
	}
	ctx := context.Background()

	// Create a token
	token, err := helpers.NewToken(ctx, args.AccessToken, args.TokenType, args.RefreshToken, time.Unix(0, 0))
	if err != nil {
		panic(err)
	}

	// Create a client
	conn, ctx, err := helpers.NewGrpcConnection(ctx, token, args.DeveloperToken, "")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// example 1: list accessible customer IDs
	examples.ListAccessibleCustomers(ctx, conn)

	// example 2: get descriptive name by customer ID
	examples.GetCustomer(ctx, conn, TestCustomerID)

	// example 3: get all campaigns
	ctx = helpers.ResetLoginCustomerID(ctx, TestCustomerID)
	examples.GetCampaigns(ctx, conn, TestCustomerID)
}
