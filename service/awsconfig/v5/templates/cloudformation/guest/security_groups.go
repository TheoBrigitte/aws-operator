package guest

const SecurityGroups = `{{define "security_groups" }}
  MasterSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: {{ .MasterSecurityGroupName }}
      VpcId: !Ref VPC
      SecurityGroupIngress:
      {{ range .MasterSecurityGroupRules }}
      -
        IpProtocol: {{ .Protocol }}
        FromPort: {{ .Port }}
        ToPort: {{ .Port }}
        CidrIp: {{ .SourceCIDR }}
      {{ end }}
      Tags:
        - Key: Name
          Value:  {{ .MasterSecurityGroupName }}

  WorkerSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: {{ .WorkerSecurityGroupName }}
      VpcId: !Ref VPC
      SecurityGroupIngress:
      {{ range .WorkerSecurityGroupRules }}
      -
        IpProtocol: {{ .Protocol }}
        FromPort: {{ .Port }}
        ToPort: {{ .Port }}
        {{ if .SourceCIDR }}
        CidrIp: {{ .SourceCIDR }}
        {{ else }}
        SourceSecurityGroupId: !Ref {{ .SourceSecurityGroup }}
        {{ end }}
      {{ end }}
      Tags:
        - Key: Name
          Value:  {{ .WorkerSecurityGroupName }}

  IngressSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: {{ .IngressSecurityGroupName }}
      VpcId: !Ref VPC
      SecurityGroupIngress:
      {{ range .IngressSecurityGroupRules }}
      -
        IpProtocol: {{ .Protocol }}
        FromPort: {{ .Port }}
        ToPort: {{ .Port }}
        CidrIp: {{ .SourceCIDR }}
      {{ end }}
      Tags:
        - Key: Name
          Value: {{ .IngressSecurityGroupName }}

  # Allow all access between masters and workers for calico. This is done after
  # the other rules to avoid circular dependencies.
  MasterAllowCalicoIngressRule:
    Type: AWS::EC2::SecurityGroupIngress
    DependsOn: MasterSecurityGroup
    Properties:
      GroupId: !Ref MasterSecurityGroup
      IpProtocol: -1
      FromPort: -1
      ToPort: -1
      SourceSecurityGroupId: !Ref MasterSecurityGroup

  MasterAllowWorkerCalicoIngressRule:
    Type: AWS::EC2::SecurityGroupIngress
    DependsOn: MasterSecurityGroup
    Properties:
      GroupId: !Ref MasterSecurityGroup
      IpProtocol: -1
      FromPort: -1
      ToPort: -1
      SourceSecurityGroupId: !Ref WorkerSecurityGroup

  WorkerAllowCalicoIngressRule:
    Type: AWS::EC2::SecurityGroupIngress
    DependsOn: WorkerSecurityGroup
    Properties:
      GroupId: !Ref WorkerSecurityGroup
      IpProtocol: -1
      FromPort: -1
      ToPort: -1
      SourceSecurityGroupId: !Ref WorkerSecurityGroup

  WorkerAllowMasterCalicoIngressRule:
    Type: AWS::EC2::SecurityGroupIngress
    DependsOn: WorkerSecurityGroup
    Properties:
      GroupId: !Ref WorkerSecurityGroup
      IpProtocol: -1
      FromPort: -1
      ToPort: -1
      SourceSecurityGroupId: !Ref MasterSecurityGroup

{{end}}`
