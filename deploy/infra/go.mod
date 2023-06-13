module infra

go 1.18

require (
	NFTM/shared v0.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go-v2 v1.16.7
	github.com/aws/aws-sdk-go-v2/config v1.15.13
	github.com/aws/constructs-go/constructs/v10 v10.1.43
	github.com/aws/jsii-runtime-go v1.61.0
	github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2 v2.3.51
	github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2 v2.0.0-rc.48
)

require (
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.9 // indirect
	github.com/aws/smithy-go v1.12.0 // indirect
)

replace NFTM/shared => ../../back-end/shared
