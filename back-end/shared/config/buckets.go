package config

import "os"

type Bucket int

type BucketConfig struct {
	Name string
}
type BucketRegistry map[Bucket]BucketConfig

const (
	Bucket_MEDIA Bucket = iota
)

func GetBuckets(env *Stage) BucketRegistry {

	SetEnv(env)

	if env == nil {
		envS := Stage(os.Getenv("env"))
		env = &envS
	}

	return BucketRegistry{
		Bucket_MEDIA: BucketConfig{
			Name: string(*env) + "-nftm-media",
		},
	}

}
