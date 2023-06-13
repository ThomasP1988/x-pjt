package bucketmedia

import (
	commonAws "NFTM/shared/common/aws"
	"NFTM/shared/config"
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Add(ctx context.Context, path string, body *[]byte, contentType string) error {
	fmt.Printf("path: %v\n", path)
	bucketName := config.Conf.Buckets[config.Bucket_MEDIA].Name
	result, err := commonAws.GetS3Client().PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &path,
		Body:        bytes.NewReader(*body),
		ContentType: &contentType,
	})

	if err != nil {
		fmt.Printf("s3 media bucket CreateIOReader: %v\n", err)
		return err
	}

	fmt.Printf("result: %v\n", result)

	return nil
}
