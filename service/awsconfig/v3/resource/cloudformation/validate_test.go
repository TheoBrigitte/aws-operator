package cloudformation

import (
	"testing"

	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/aws-operator/service/awsconfig/v3/resource/cloudformation/adapter"
	"github.com/giantswarm/micrologger/microloggertest"
)

func Test_validateHostPeeringRoutes(t *testing.T) {
	t.Parallel()
	customObject := v1alpha1.AWSConfig{
		Spec: v1alpha1.AWSConfigSpec{
			AWS: v1alpha1.AWSConfigSpecAWS{
				VPC: v1alpha1.AWSConfigSpecAWSVPC{
					PrivateSubnetCIDR: "172.31.0.0/16",
				},
			},
		},
	}

	testCases := []struct {
		description          string
		unexistentRouteTable bool
		expectedError        bool
	}{
		{
			description:          "route table doesn't exist, do not expect error",
			unexistentRouteTable: true,
			expectedError:        false,
		},
		{
			description:          "route table exists, expect error",
			unexistentRouteTable: false,
			expectedError:        true,
		},
	}

	for _, tc := range testCases {
		var err error
		var newResource *Resource
		{
			resourceConfig := DefaultConfig()
			resourceConfig.Clients = &adapter.Clients{
				EC2: &adapter.EC2ClientMock{},
				IAM: &adapter.IAMClientMock{},
				KMS: &adapter.KMSClientMock{},
			}
			ec2Mock := &adapter.EC2ClientMock{}
			ec2Mock.SetUnexistingRouteTable(tc.unexistentRouteTable)
			resourceConfig.HostClients = &adapter.Clients{
				EC2:            ec2Mock,
				CloudFormation: &adapter.CloudFormationMock{},
				IAM:            &adapter.IAMClientMock{},
			}

			resourceConfig.Logger = microloggertest.New()
			newResource, err = New(resourceConfig)
			if err != nil {
				t.Error("expected", nil, "got", err)
			}
		}

		t.Run(tc.description, func(t *testing.T) {
			err := newResource.validateHostPeeringRoutes(customObject)
			if tc.expectedError && err == nil {
				t.Errorf("expected error didn't happen")
			}
			if !tc.expectedError && err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}
