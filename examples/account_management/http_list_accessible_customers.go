package account_management

import (
	"fmt"
	"os"

	"github.com/shenzhencenter/google-ads-pb/examples"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/protobuf/encoding/protojson"
)

func RunHTTPExample_ListAccessibleCustomers() {
	ListAccessibleCustomers_HTTP()
}

func ListAccessibleCustomers_HTTP() {
	listAccessibleCustomersRequest, err := protojson.Marshal(
		&services.ListAccessibleCustomersRequest{},
	)
	if err != nil {
		panic(err)
	}
	listAccessibleCustomersResponse := new(services.ListAccessibleCustomersResponse)

	response := examples.HttpRequest(listAccessibleCustomersRequest,
		examples.ListAccessibleCustomersMehtod,
		examples.ListAccessibleCustomersPath,
		examples.WithHeader("authorization", os.Getenv("ACCESS_TOKEN")),
		examples.WithHeader("developer-token", os.Getenv("DEVELOPER_TOKEN")),
	)
	if err := protojson.Unmarshal(response, listAccessibleCustomersResponse); err != nil {
		panic(err)
	}

	for _, customer := range listAccessibleCustomersResponse.ResourceNames {
		fmt.Println("ResourceName: " + customer)
	}
}
