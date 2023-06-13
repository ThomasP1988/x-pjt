package main

import (
	"NFTM/shared/config"
	"context"
	"fmt"
	"log"

	"NFTM/apis/graphql/resolvers/collection"
	nftitem "NFTM/apis/graphql/resolvers/item"
	"NFTM/apis/graphql/resolvers/notification"
	"NFTM/apis/graphql/resolvers/search"
	"NFTM/apis/graphql/resolvers/token"
	"NFTM/apis/graphql/resolvers/user"
	"NFTM/apis/graphql/resolvers/wallet"
	utils "NFTM/apis/graphql/utils"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	config.GetConfig(nil)
	lambda.Start(Handler)
}

var Resolvers map[string]map[string]utils.Resolver = map[string]map[string]utils.Resolver{
	"Query": {
		"listCollections":      collection.List,
		"listCollectionsByIds": collection.ListByIds,
		"listItemsByIds":       nftitem.ListByKeys,
		"listNotifications":    notification.List,
		"me":                   user.Me,
		"searchCollections":    search.SearchCollection,
		"getItem":              nftitem.GetItem,
		"getCollection":        collection.GetItem,
	},
	"Mutation": {
		"connectWallet":           wallet.ConnectWallet,
		"createToken":             token.CreateNonce,
		"inviteUser":              user.Invite,
		"setLastSeenNotification": user.SetLastSeenNotification,
		"submitCollection":        collection.Submit,
		"validateCollection":      collection.Validate,
	},
	"Subscription": {
		"onUpdatedItem": nftitem.Subscribe,
		"notification":  notification.Subscribe,
	},
	"Collection": {
		"items": nftitem.ListByIds,
	},
}

// func Handler(ctx context.Context, evt map[string]interface{}) (interface{}, error) {
// 	evtBytes, err := json.Marshal(evt)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	fmt.Printf("evt: %+v\n", string(evtBytes))
// 	event := utils.ResolverEvent{}
// 	err = json.Unmarshal(evtBytes, &event)

//	if err != nil {
//		fmt.Printf("err: %v\n", err)
//	}
//
// useful for debugging
func Handler(ctx context.Context, event *utils.ResolverEvent) (interface{}, error) {

	log.Printf("event: %+v\n", event)
	fmt.Printf("ctx: %+v\n", ctx)

	var isAdmin bool = false

	if event.Identity.CognitoIdentityAuthType == utils.AuthType_Authenticated && event.Identity.CognitoIdentityID == config.Conf.Admin.UserPoolIdentity {
		isAdmin = true
		fmt.Printf("isAdmin: %v\n", isAdmin)
	}

	// err := notif_service.Create(ctx, notif_entity.Notification{
	// 	UserID:  event.Identity.CognitoIdentityID,
	// 	Type:    notif_entity.Welcome_NotificationType,
	// 	Message: "Welcome to NFT Quant" + time.Now().String(),
	// })
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	return Resolvers[string(event.Info.ParentTypeName)][event.Info.FieldName](ctx, utils.ResolverArgs{
		Args:   &event.Arguments,
		UserID: &event.Identity.CognitoIdentityID,
		Event:  event,
	})
}
