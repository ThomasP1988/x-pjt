package nftitem

import (
	"NFTM/shared/config"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type FetchItemMessage struct {
	CollectionAddress      string `json:"collectionAddress"`
	IndexItem              *int   `json:"indexItem"`
	TokenID                *int   `json:"tokenId"`
	ShouldUpdateCollection bool   `json:"shouldUpdateCollection"`
}

func FetchItem(ctx context.Context, args *FetchItemMessage) (*sns.PublishOutput, error) {
	nfts := GetNFTItemService()
	jsonArgs, err := json.Marshal(args)

	if err != nil {
		fmt.Printf("error marshalling FetchItems SNS arguments: %v\n", err)
		return nil, err
	}

	stringArgs := string(jsonArgs)

	topicArn := fmt.Sprintf("arn:aws:sns:%s:%s:%s", *config.Conf.Region, config.Conf.AWSAccount, config.Conf.SNS.FetchItemTopic)

	return nfts.SNSClient.Publish(ctx, &sns.PublishInput{
		Message:  &stringArgs,
		TopicArn: &topicArn,
	})
}
