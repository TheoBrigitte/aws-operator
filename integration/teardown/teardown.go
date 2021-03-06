// +build k8srequired

package teardown

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/giantswarm/e2e-harness/pkg/framework"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/aws-operator/integration/client"
	"github.com/giantswarm/aws-operator/integration/env"
)

func HostPeerVPC(c *client.AWS, g *framework.Guest, h *framework.Host) error {
	log.Printf("Deleting Host Peer VPC stack")

	_, err := c.CloudFormation.DeleteStack(&cloudformation.DeleteStackInput{
		StackName: aws.String("host-peer-" + env.ClusterID()),
	})
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func Teardown(c *client.AWS, g *framework.Guest, h *framework.Host) error {
	var err error

	{
		err = framework.HelmCmd("delete aws-operator --purge")
		if err != nil {
			return microerror.Mask(err)
		}
		err = framework.HelmCmd("delete cert-operator --purge")
		if err != nil {
			return microerror.Mask(err)
		}
		err = framework.HelmCmd("delete node-operator --purge")
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		err = framework.HelmCmd("delete cert-resource-lab --purge")
		if err != nil {
			return microerror.Mask(err)
		}
		err = framework.HelmCmd("delete aws-resource-lab --purge")
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
