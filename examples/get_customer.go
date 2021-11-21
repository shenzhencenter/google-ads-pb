package examples

import (
	"context"
	"fmt"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
)

func GetCustomer(ctx context.Context, conn *grpc.ClientConn, customerID string) {
	request := services.GetCustomerRequest{
		ResourceName: "customers/" + customerID,
	}
	customer, err := services.NewCustomerServiceClient(conn).GetCustomer(ctx, &request)
	if err != nil {
		return
	}

	fmt.Println(*customer.DescriptiveName)
}
