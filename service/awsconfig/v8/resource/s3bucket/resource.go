package s3bucket

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	awsservice "github.com/giantswarm/aws-operator/service/aws"
	"github.com/giantswarm/aws-operator/service/awsconfig/v8/key"
)

const (
	// Name is the identifier of the resource.
	Name = "s3bucketv8"
)

// Config represents the configuration used to create a new s3bucket resource.
type Config struct {
	// Dependencies.
	AwsService *awsservice.Service
	Clients    Clients
	Logger     micrologger.Logger

	// Settings.
	InstallationName string
}

// DefaultConfig provides a default configuration to create a new s3bucket
// resource by best effort.
func DefaultConfig() Config {
	return Config{
		// Dependencies.
		AwsService: nil,
		Clients:    Clients{},
		Logger:     nil,
	}
}

// Resource implements the s3bucket resource.
type Resource struct {
	// Dependencies.
	awsService *awsservice.Service
	clients    Clients
	logger     micrologger.Logger

	// Settings.
	installationName string
}

// New creates a new configured s3bucket resource.
func New(config Config) (*Resource, error) {
	// Dependencies.
	if config.AwsService == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.AwsService must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	// Settings.
	if config.InstallationName == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.InstallationName must not be empty", config)
	}

	newResource := &Resource{
		// Dependencies.
		awsService: config.AwsService,
		clients:    config.Clients,
		logger:     config.Logger,

		// Settings.
		installationName: config.InstallationName,
	}

	return newResource, nil
}

func (r *Resource) Name() string {
	return Name
}

func toBucketState(v interface{}) (BucketState, error) {
	if v == nil {
		return BucketState{}, nil
	}

	bucketState, ok := v.(BucketState)
	if !ok {
		return BucketState{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", BucketState{}, v)
	}

	return bucketState, nil
}

func (r *Resource) getS3BucketTags(customObject v1alpha1.AWSConfig) []*s3.Tag {
	clusterTags := key.ClusterTags(customObject, r.installationName)
	s3Tags := []*s3.Tag{}

	for k, v := range clusterTags {
		tag := &s3.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		s3Tags = append(s3Tags, tag)
	}

	return s3Tags
}
