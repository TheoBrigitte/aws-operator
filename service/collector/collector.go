package collector

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	awsutil "github.com/giantswarm/aws-operator/client/aws"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

const (
	Namespace = "aws_operator"

	CidrLabel  = "cidr"
	IDLabel    = "id"
	NameLabel  = "name"
	StateLabel = "state"

	gaugeValue = float64(1)
)

type Config struct {
	Logger micrologger.Logger

	AwsConfig awsutil.Config
}

type Collector struct {
	logger micrologger.Logger

	awsClients awsutil.Clients

	vpcs *prometheus.Desc
}

func New(config Config) (*Collector, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Logger must not be empty")
	}

	var emptyAwsConfig awsutil.Config
	if config.AwsConfig == emptyAwsConfig {
		return nil, microerror.Maskf(invalidConfigError, "config.AwsConfig must not be empty")
	}

	awsClients := awsutil.NewClients(config.AwsConfig)

	c := &Collector{
		logger: config.Logger,

		awsClients: awsClients,

		vpcs: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "vpc_info"),
			"VPC information.",
			[]string{CidrLabel, IDLabel, NameLabel, StateLabel},
			nil,
		),
	}

	return c, nil
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.vpcs
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.logger.Log("debug", "collecting metrics")

	c.collectVPCs(ch)

	c.logger.Log("debug", "finished collecting metrics")
}

func (c *Collector) collectVPCs(ch chan<- prometheus.Metric) {
	c.logger.Log("debug", "collecting metrics for vpc")

	resp, err := c.awsClients.EC2.DescribeVpcs(&ec2.DescribeVpcsInput{})
	if err != nil {
		c.logger.Log("debug", "could not list vpcs: ", err)
	}

	for _, vpc := range resp.Vpcs {
		name := ""
		for _, tag := range vpc.Tags {
			if *tag.Key == "Name" {
				name = *tag.Value
			}
		}

		ch <- prometheus.MustNewConstMetric(
			c.vpcs, prometheus.GaugeValue, gaugeValue,
			*vpc.CidrBlock, *vpc.VpcId, name, *vpc.State,
		)
	}

	c.logger.Log("debug", "finished collecting metrics for vpc")
}
