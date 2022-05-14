package account_management

import (
	"context"
	"fmt"
	"os"

	"github.com/shenzhencenter/google-ads-pb/examples"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
)

func RunGRPCExample_ListAccessibleCustomers() {
	conn := examples.GetGRPCConnection()
	defer conn.Close()
	ctx := examples.SetContext(context.Background(),
		examples.WithContext("authorization", os.Getenv("ACCESS_TOKEN")),
		examples.WithContext("developer-token", os.Getenv("DEVELOPER_TOKEN")),
	)
	ListAccessibleCustomers(ctx, conn)
}

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
