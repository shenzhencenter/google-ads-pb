package account_management

import (
	"fmt"
	"os"

	"github.com/shenzhencenter/google-ads-pb/examples"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/protobuf/encoding/protojson"
)

func RunHTTPExample_GetAccountInformation() {
	GetAccountInformation_HTTP(os.Getenv("CUSTOMER_ID"))
}

func GetAccountInformation_HTTP(customerID string) {
	searchGoogleAdsRequest := services.SearchGoogleAdsRequest{
		CustomerId: customerID,
		Query:      "SELECT customer.descriptive_name FROM customer WHERE customer.id = " + customerID,
	}
	request, err := protojson.Marshal(&searchGoogleAdsRequest)
	if err != nil {
		panic(err)
	}
	searchGoogleAdsResponse := new(services.SearchGoogleAdsResponse)

	response := examples.HttpRequest(request,
		examples.GoogleAdsSearchMehtod,
		fmt.Sprintf(examples.GoogleAdsSearchPath, customerID),
		examples.WithHeader("authorization", os.Getenv("ACCESS_TOKEN")),
		examples.WithHeader("developer-token", os.Getenv("DEVELOPER_TOKEN")),
		examples.WithHeader("login-customer-id", os.Getenv("CUSTOMER_ID")),
	)
	if err := protojson.Unmarshal(response, searchGoogleAdsResponse); err != nil {
		panic(err)
	}

	for _, resource := range searchGoogleAdsResponse.Results {
		fmt.Println(resource.Customer.GetDescriptiveName())
	}
}
