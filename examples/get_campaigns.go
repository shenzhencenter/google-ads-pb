package examples

import (
	"context"
	"fmt"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
)

func GetCampaigns(ctx context.Context, conn *grpc.ClientConn, customerId string) {
	request := services.SearchGoogleAdsRequest{
		CustomerId: customerId,
		Query:      "SELECT campaign.id, campaign.name FROM campaign",
	}
	response, err := services.NewGoogleAdsServiceClient(conn).Search(ctx, &request)
	if err != nil {
		panic(err)
	}

	for _, resource := range response.Results {
		fmt.Println(resource.Campaign.ResourceName)
	}
}
