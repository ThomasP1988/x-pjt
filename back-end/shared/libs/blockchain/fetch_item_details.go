package blockchain

import (
	"NFTM/shared/entities/nft"
	repo_nftitem "NFTM/shared/repositories/nft-item"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func FetchItemDetailsByIndex(ctx context.Context, contract721 *ERC721, collection *nft.Collection, index int) (*nft.Item, error) {
	tokenId, err := contract721.ERC721Caller.TokenByIndex(&bind.CallOpts{
		Context: ctx,
	}, big.NewInt(int64(index)))

	if err != nil {
		fmt.Println("error getting token by index", err)
		return nil, err
	}
	return FetchItemDetails(ctx, contract721, collection, int(tokenId.Int64()))
}

func FetchItemDetails(ctx context.Context, contract721 *ERC721, collection *nft.Collection, tokenIdInt int) (*nft.Item, error) {

	tokenId := big.NewInt(int64(tokenIdInt))

	tokenURI, err := contract721.ERC721Caller.TokenURI(&bind.CallOpts{
		Context: ctx,
	}, tokenId)
	fmt.Printf("tokenURI: %v\n", tokenURI)

	if err != nil {
		fmt.Println("error getting URI", err)
		return nil, err
	}

	url := formatAddress(tokenURI)
	nftItem := &nft.Item{}

	err = getRarities(url, nftItem)

	if err != nil {
		fmt.Println("error getting NFT rarities", err)
		return nil, err
	}

	nftItem.CollectionAddress = collection.Address
	nftItem.TokenID = int(tokenId.Int64())
	nftItem.TokenURI = tokenURI
	nftItem.IsFetching = false

	SaveImages(ctx, nftItem)

	err = repo_nftitem.Create(ctx, nftItem)

	if err != nil {
		fmt.Println("error saving NFT Item", err)
		return nil, err
	}

	return nftItem, nil
}

func formatAddress(uri string) string {
	if strings.HasPrefix(uri, "ipfs://") {
		return "https://ipfs.io/ipfs/" + strings.Replace(uri, "ipfs://", "", 1)
	}
	return uri
}

func getRarities(url string, nftItem *nft.Item) error {
	fmt.Printf("url: %v\n", url)
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("error getting rarities: %v\n", err)
		return err
	}
	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(nftItem)

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	bodyString := string(bodyBytes)

	fmt.Printf("response: %v\n", bodyString)

	return nil
}
