package blockchain

import (
	"NFTM/shared/entities/nft"
	bucketmedia "NFTM/shared/repositories/bucket-media"
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"strings"

	"github.com/nfnt/resize"
)

const Delimitator string = "/"

func SaveImages(ctx context.Context, nftItem *nft.Item) error {
	fmt.Printf("\"save images\": %v\n", "save images")
	imgURL := formatAddress(nftItem.Image)

	response, err := http.Get(imgURL)

	if err != nil {
		fmt.Printf("nftItem.Image: %v\n", nftItem.Image)
		fmt.Printf("imgURL: %v\n", imgURL)
		fmt.Printf("error getting image: %v\n", err)
		return err
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	filetype := http.DetectContentType(bodyBytes)

	if strings.HasPrefix(filetype, "image") {
		ext := strings.Replace(filetype, "image/", "", 1)
		path := strings.Join([]string{
			"protected",
			nftItem.CollectionAddress,
			fmt.Sprint(nftItem.TokenID),
		}, Delimitator)

		mainPath, err := saveMain(ctx, path, ext, &bodyBytes, filetype)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return nil
		}

		thumbnailPath, err := saveThumbnail(ctx, path, ext, &bodyBytes, filetype, 180)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return nil
		}

		fmt.Printf("thumbnailPath: %v\n", thumbnailPath)

		nftItem.ThumbnailPath = thumbnailPath
		nftItem.ImagePath = mainPath
	}

	return nil
}

func saveMain(ctx context.Context, path string, ext string, body *[]byte, filetype string) (string, error) {
	mainPath := strings.Join([]string{
		path,
		"main." + ext,
	}, Delimitator)

	err := bucketmedia.Add(ctx, mainPath, body, filetype)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}
	return mainPath, nil
}

func saveThumbnail(ctx context.Context, path string, ext string, body *[]byte, filetype string, maxDimension uint) (string, error) {
	thumbnailPath := strings.Join([]string{
		path,
		"thumbnail.jpeg",
	}, Delimitator)

	img, format, err := image.Decode(bytes.NewBuffer(*body))
	fmt.Printf("format: %v\n", format)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}

	thumbnailImg := resize.Thumbnail(maxDimension, maxDimension, img, resize.Lanczos3)

	thumbnailBody := &[]byte{}
	thumbnailWriter := bytes.NewBuffer(*thumbnailBody)
	err = jpeg.Encode(thumbnailWriter, thumbnailImg, nil)

	if err != nil {
		fmt.Printf("error encoding thumbnail jpeg: %v\n", err)
		return "", err
	}

	thumbnailBytes := thumbnailWriter.Bytes()

	err = bucketmedia.Add(ctx, thumbnailPath, &thumbnailBytes, "image/jpeg")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}

	return thumbnailPath, nil
}
