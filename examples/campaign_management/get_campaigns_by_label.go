package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/shenzhencenter/google-ads-pb/examples"
	"github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
)

const GAQL_GetCampaignsByLabel = `SELECT campaign.id, campaign.name, label.id, label.name FROM campaign_label WHERE label.id = '%s' ORDER BY campaign.id`

const LabelID = "INSERT_LABEL_ID_HERE"

func RunGRPCExample_GetCampaignsByLabel() {
	conn := examples.GetGRPCConnection()
	defer conn.Close()
	ctx := examples.SetContext(context.Background(),
		examples.WithContext("authorization", os.Getenv("ACCESS_TOKEN")),
		examples.WithContext("developer-token", os.Getenv("DEVELOPER_TOKEN")),
		examples.WithContext("login-customer-id", os.Getenv("CUSTOMER_ID")),
	)
	GetCampaignsByLabel(ctx, conn, os.Getenv("CUSTOMER_ID"), LabelID)
}

func GetCampaignsByLabel(ctx context.Context, conn *grpc.ClientConn, customerID string, labelID string) {
	request := services.SearchGoogleAdsRequest{
		CustomerId: customerID,
		Query:      fmt.Sprintf(GAQL_GetCampaignsByLabel, labelID),
	}
	response, err := services.NewGoogleAdsServiceClient(conn).Search(ctx, &request)
	if err != nil {
		panic(err)
	}

	for _, resource := range response.Results {
		fmt.Printf("Campaign ID: %d\n", *resource.Campaign.Id)
		fmt.Printf("Campaign name: %s\n", *resource.Campaign.Name)
		fmt.Printf("Label ID: %d\n", *resource.Label.Id)
		fmt.Printf("Label name: %s\n", *resource.Label.Name)
	}
}
