package examples

import (
	"context"
	"fmt"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
)

func GetCustomer(ctx context.Context, conn *grpc.ClientConn, customerID string) {
	request := services.SearchGoogleAdsRequest{
		CustomerId: customerID,
		Query: 	"SELECT customer.descriptive_name FROM customer WHERE customer.id = " + customerID,
	}
	search, err := services.NewGoogleAdsServiceClient(conn).Search(ctx, &request)
	if err != nil {
		return
	}

	for _, resource := range search.Results {
		fmt.Println(*resource.Customer.DescriptiveName)
	}
}
