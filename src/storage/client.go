package storage

import (
	"context"
	"log"
	"net/url"

	"example.com/review/v2/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	transport "github.com/aws/smithy-go/endpoints"
)

type Resolver struct {
	URL *url.URL
}

func (r *Resolver) ResolveEndpoint(_ context.Context, params s3.EndpointParameters) (transport.Endpoint, error) {
	u := *r.URL
	u.Path += "/" + *params.Bucket
	return transport.Endpoint{URI: u}, nil
}

func NewClient(cfg *config.Config) (*s3.Client, error) {
	var client *s3.Client

	if cfg.Provider == "minio" {
		endpoint, err := url.Parse(cfg.Endpoint)
		// fmt.Println(endpoint)
		if err != nil {
			log.Fatalf("Unable to parse MinIO endpoint: %v", err)
		}
		client = s3.New(s3.Options{
			EndpointResolverV2: &Resolver{URL: endpoint},
			Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     cfg.AccessKey,
					SecretAccessKey: cfg.SecretKey,
				}, nil
			}),
		})
	} else {
		client = s3.New(s3.Options{
			Region: cfg.Region,
			Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     cfg.AccessKey,
					SecretAccessKey: cfg.SecretKey,
				}, nil
			}),
		})
	}

	return client, nil
}
