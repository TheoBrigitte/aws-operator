package guest

const Instance = `{{define "instance"}}
  MasterInstance:
    Type: "AWS::EC2::Instance"
    Description: Master instance
    Properties:
      AvailabilityZone: {{ .MasterAZ }}
      IamInstanceProfile: !Ref MasterInstanceProfile
      ImageId: {{ .MasterImageID }}
      InstanceType: {{ .MasterInstanceType }}
      SecurityGroupIds:
      - !Ref MasterSecurityGroup
      SubnetId: !Ref PrivateSubnet
      UserData: {{ .MasterSmallCloudConfig }}
      Tags:
      - Key: Name
        Value: {{ .ClusterID }}-master
  EtcdVolume:
    Type: AWS::EC2::Volume
    DependsOn:
    - MasterInstance
    Properties:
      Size: 100
      VolumeType: gp2
      AvailabilityZone: !GetAtt MasterInstance.AvailabilityZone
      Tags:
      - Key: Name
        Value: {{ .ClusterID }}-etcd
  MountPoint:
    Type: AWS::EC2::VolumeAttachment
    Properties:
      InstanceId: !Ref MasterInstance
      VolumeId: !Ref EtcdVolume
      Device: /dev/sdh
{{end}}`
