package s3

import (
	"NFTM/shared/config"
	"aws/helper"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func SetMediaBucket(stack constructs.Construct) *awss3.Bucket {
	bucket := awss3.NewBucket(stack, helper.SetName("media-bucket"), &awss3.BucketProps{
		BucketName:       jsii.String(config.Conf.Buckets[config.Bucket_MEDIA].Name),
		PublicReadAccess: jsii.Bool(false),
	})

	awscdk.NewCfnOutput(stack, jsii.String("MediaBucket/ARN"), &awscdk.CfnOutputProps{
		Value: bucket.BucketArn(),
	})

	return &bucket
}
