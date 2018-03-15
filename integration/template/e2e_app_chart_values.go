package template

// E2EAppChartValues values required by e2e-app-chart, the environment variables
// will be expanded before writing the contents to a file.
var E2EAppChartValues = `Installation:
  V1:
    Secret:
      Registry:
        PullSecret:
          DockerConfigJSON: "{\"auths\":{\"quay.io\":{\"auth\":\"$REGISTRY_PULL_SECRET\"}}}"
`
