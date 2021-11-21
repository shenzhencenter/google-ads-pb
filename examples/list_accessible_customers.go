package examples

import (
	"context"
	"fmt"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
)

func ListAccessibleCustomers(ctx context.Context, conn *grpc.ClientConn) {
	customers, err := services.NewCustomerServiceClient(conn).
		ListAccessibleCustomers(ctx, &services.ListAccessibleCustomersRequest{})
	if err != nil {
		fmt.Printf("%+#v", err)
		return
	}

	for _, customer := range customers.ResourceNames {
		fmt.Println("ResourceName: " + customer)
	}
}
