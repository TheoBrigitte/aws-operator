package template

// AWSOperatorChartValues values required by aws-operator-chart, the environment
// variables will be expanded before writing the contents to a file.
var AWSOperatorChartValues = `Installation:
  V1:
    Guest:
      Kubernetes:
        API:
          Auth:
            Provider:
              OIDC:
                ClientID: ""
                IssueURL: ""
                UsernameClaim: ""
                GroupsClaim: ""
      Update:
        Enabled: ${GUEST_UPDATE_ENABLED}
    Name: ci-aws-operator
    Provider:
      AWS:
        Region: ${AWS_REGION}
    Secret:
      AWSOperator:
        IDRSAPub: ${IDRSA_PUB}
        SecretYaml: |
          service:
            aws:
              accesskey:
                id: ${GUEST_AWS_ACCESS_KEY_ID}
                secret: ${GUEST_AWS_SECRET_ACCESS_KEY}
                token: ${GUEST_AWS_SESSION_TOKEN}
              hostaccesskey:
                id: ${HOST_AWS_ACCESS_KEY_ID}
                secret: ${HOST_AWS_SECRET_ACCESS_KEY}
                token: ${HOST_AWS_SESSION_TOKEN}
      Registry:
        PullSecret:
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"${REGISTRY_PULL_SECRET}\"}}}"
`
