# Harness SDK for Go

This project is a Harness SDK for the go programming language. It provides go client for interacting with the current gen GraphQL and Config-as-Code API's.

## Disclaimer

This product is not supported by the Harness Customer support team. If you have any questions please open a [new issue](https://github.com/harness-io/harness-go-sdk/issues/new) or join our slack [channel](https://harnesscommunity.slack.com/archives/C02G9CUNF1S).

## Getting Started

### Installing

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or project's Go module dependencies.

```bash
go get github.com/harness-io/harness-go-sdk
```

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

```bash
go get -u github.com/harness-io/harness-go-sdk
```

### Dependencies

The metadata of the SDK's dependencies can be found in the Go module file `go.mod`.

### Go Modules

If you are using Go modules, your go get will default to the latest tagged release version of the SDK. To get a specific release version of the SDK use @\<tag\> in your go get command.

```bash
go get github.com/harness-io/harness-go-sdk@v0.2.11
```

To get the latest SDK repository change use @latest.

```bash
go get github.com/harness-io/harness-go-sdk@latest
```

## Quick Examples

### Get an application by name

```go
client := NewClient()
app, err := client.ApplicationClient.GetApplicationByName("my-app)
```

### Create a Service

```go
svc, _ := ServiceFactory(app.Id, serviceName, cac.DeploymentTypes.Kubernetes, cac.ArtifactTypes.Docker)
svc.ApplicationId = app.Id

newService, err := client.Services().UpsertService(svc)
```

### Configuration

There are a few environment variables you can set to configure the api client.

- `HARNESS_ACCOUNT_ID`: (required) The ID of the harness account you are connecting to.
- `HARNESS_API_KEY`: (required) The API Key used for authentication.
- `HARNESS_BEARER_TOKEN`: (optional) The authentication bearer token. This is needed for certain API calls to the `config-as-code` API's. This will be deprecated in the near future once those endpoints are updated.
- `HARNESS_ENDPOINT`: (optional) The FQDN for contacting the Harness managers. Defaults to `https://app.harness.io`.

If you need to provide additional configuration you can create a client object from scratch.

```go
client := &Client{
    UserAgent:   getUserAgentString(),
    Endpoint:    utils.GetEnv(envvar.Endpoint, utils.DefaultApiUrl),
    AccountId:   os.Getenv(envvar.AccountId),
    APIKey:      os.Getenv(envvar.ApiKey),
    BearerToken: os.Getenv(envvar.BearerToken),
    HTTPClient: &retryablehttp.Client{
        RetryMax:     10,
        RetryWaitMin: 5 * time.Second,
        RetryWaitMax: 10 * time.Second,
        HTTPClient: &http.Client{
            Timeout: 10 * time.Second,
        },
        Backoff:    retryablehttp.DefaultBackoff,
        CheckRetry: retryablehttp.DefaultRetryPolicy,
    },
}
```

The `github.com/hashicorp/go-retryablehttp` is essentially a drop-in replacement for the `http` package and is used to handle retries when getting rate limited.
