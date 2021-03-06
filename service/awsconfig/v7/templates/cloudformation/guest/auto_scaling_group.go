package guest

const AutoScalingGroup = `{{define "autoscaling_group"}}
  {{ .ASGType }}AutoScalingGroup:
    Type: "AWS::AutoScaling::AutoScalingGroup"
    Properties:
      VPCZoneIdentifier:
        - !Ref PrivateSubnet
      AvailabilityZones: [{{ .WorkerAZ }}]
      MinSize: {{ .ASGMinSize }}
      MaxSize: {{ .ASGMaxSize }}
      LaunchConfigurationName: !Ref {{ .ASGType }}LaunchConfiguration
      LoadBalancerNames:
        - !Ref IngressLoadBalancer
      HealthCheckGracePeriod: {{ .HealthCheckGracePeriod }}
      MetricsCollection:
        - Granularity: "1Minute"
      Tags:
        - Key: Name
          Value: {{ .ClusterID }}-{{ .ASGType }}
          PropagateAtLaunch: true
    UpdatePolicy:
      AutoScalingRollingUpdate:
        # minimum amount of instances that must always be running during a rolling update
        MinInstancesInService: {{ .MinInstancesInService }}
        # only do a rolling update of this amount of instances max
        MaxBatchSize: {{ .MaxBatchSize }}
        # after creating a new instance, pause operations on the ASG for this amount of time
        PauseTime: {{ .RollingUpdatePauseTime }}
{{end}}`
