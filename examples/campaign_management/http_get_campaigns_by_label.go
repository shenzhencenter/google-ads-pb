package examples

import (
	"fmt"
	"os"

	"github.com/shenzhencenter/google-ads-pb/examples"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/protobuf/encoding/protojson"
)

func RunHTTPExample_GetCampaignsByLabel() {
	GetCampaignsByLabel_HTTP(os.Getenv("CUSTOMER_ID"), LabelID)
}

func GetCampaignsByLabel_HTTP(customerID string, labelID string) {
	searchGoogleAdsRequest := services.SearchGoogleAdsRequest{
		CustomerId: customerID,
		Query:      fmt.Sprintf(GAQL_GetCampaignsByLabel, labelID),
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
		fmt.Printf("Campaign ID: %d\n", resource.Campaign.GetId())
		fmt.Printf("Campaign name: %s\n", resource.Campaign.GetName())
		fmt.Printf("Label ID: %d\n", resource.Label.GetId())
		fmt.Printf("Label name: %s\n", resource.Label.GetName())
	}
}
