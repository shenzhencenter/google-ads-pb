package account_management

import (
	"context"
	"fmt"
	"os"

	"github.com/shenzhencenter/google-ads-pb/examples"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
)

func RunGRPCExample_GetAccountInformation() {
	conn := examples.GetGRPCConnection()
	defer conn.Close()
	ctx := examples.SetContext(context.Background(),
		examples.WithContext("authorization", os.Getenv("ACCESS_TOKEN")),
		examples.WithContext("developer-token", os.Getenv("DEVELOPER_TOKEN")),
		examples.WithContext("login-customer-id", os.Getenv("CUSTOMER_ID")),
	)
	GetAccountInformation(ctx, conn, os.Getenv("CUSTOMER_ID"))
}

func GetAccountInformation(ctx context.Context, conn *grpc.ClientConn, customerID string) {
	request := services.SearchGoogleAdsRequest{
		CustomerId: customerID,
		Query:      "SELECT customer.descriptive_name FROM customer WHERE customer.id = " + customerID,
	}
	search, err := services.NewGoogleAdsServiceClient(conn).Search(ctx, &request)
	if err != nil {
		return
	}

	for _, resource := range search.Results {
		fmt.Println(resource.Customer.GetDescriptiveName())
	}
}
