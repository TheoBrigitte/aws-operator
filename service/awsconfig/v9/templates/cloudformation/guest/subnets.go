package guest

const Subnets = `{{define "subnets"}}
  PublicSubnet:
    Type: AWS::EC2::Subnet
    Properties:
      AvailabilityZone: {{ .PublicSubnetAZ }}
      CidrBlock: {{ .PublicSubnetCIDR }}
      MapPublicIpOnLaunch: {{ .PublicSubnetMapPublicIPOnLaunch }}
      Tags:
      - Key: Name
        Value: {{ .PublicSubnetName }}
      VpcId: !Ref VPC

  PublicSubnetRouteTableAssociation:
    Type: AWS::EC2::SubnetRouteTableAssociation
    Properties:
      RouteTableId: !Ref PublicRouteTable
      SubnetId: !Ref PublicSubnet

  PrivateSubnet:
    Type: AWS::EC2::Subnet
    Properties:
      AvailabilityZone: {{ .PrivateSubnetAZ }}
      CidrBlock: {{ .PrivateSubnetCIDR }}
      MapPublicIpOnLaunch: {{ .PrivateSubnetMapPublicIPOnLaunch }}
      Tags:
      - Key: Name
        Value: {{ .PrivateSubnetName }}
      VpcId: !Ref VPC

  PrivateSubnetRouteTableAssociation:
    Type: AWS::EC2::SubnetRouteTableAssociation
    Properties:
      RouteTableId: !Ref PrivateRouteTable
      SubnetId: !Ref PrivateSubnet
{{end}}`
