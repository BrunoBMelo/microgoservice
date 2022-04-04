package appconfig

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/joho/godotenv"
)

type Config struct {
	PortApp          string
	AwsConfig        *aws.Config
	isDevelopment    bool
	localstackRegion string
	awsPartitionId   string
	localstackUrl    string
}

func LoadConfig() Config {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	configApp := Config{
		PortApp:          os.Getenv("PORT"),
		AwsConfig:        &aws.Config{},
		isDevelopment:    false,
		localstackRegion: os.Getenv("LOCALSTACK_AWS_REGION"),
		awsPartitionId:   os.Getenv("LOCALSTACK_PARTITION_ID"),
		localstackUrl:    os.Getenv("LOCALSTACK_URL"),
	}

	if os.Getenv("ENVIRONMENT") == "env" || os.Getenv("ENVIRONMENT") == "" {
		configApp.isDevelopment = true
	}

	envDev := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   configApp.awsPartitionId,
			URL:           configApp.localstackUrl,
			SigningRegion: configApp.AwsConfig.Region,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(envDev))
	if err != nil {
		panic(err)
	}

	configApp.AwsConfig = &cfg

	return configApp
}
