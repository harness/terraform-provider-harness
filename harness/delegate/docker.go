package delegate

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	harnessDockerImage = "harness/delegate:latest"
)

type DockerDelegateConfig struct {
	AccountId              string
	AccountSecret          string
	DelegateName           string
	ContainerName          string
	Image                  string
	ProfileId              string
	EnvVars                map[string]string
	ClientOptions          client.Opt
	ContainerConfig        *container.Config
	HostConfig             *container.HostConfig
	NetworkingConfig       *network.NetworkingConfig
	Platform               *specs.Platform
	ContainerStartOptions  types.ContainerStartOptions
	ContainerRemoveOptions types.ContainerRemoveOptions
}

func getDefaultEnvConfig(cfg *DockerDelegateConfig) map[string]string {
	return map[string]string{
		"ACCOUNT_ID":                         cfg.AccountId,
		"ACCOUNT_SECRET":                     cfg.AccountSecret,
		"CDN_URL":                            "https://app.harness.io",
		"CF_CLI6_PATH":                       "",
		"CF_CLI7_PATH":                       "",
		"CLIENT_TOOLS_DOWNLOAD_DISABLED":     "false",
		"DELEGATE_CHECK_LOCATION":            "delegatefree.txt",
		"DELEGATE_NAME":                      cfg.DelegateName,
		"DELEGATE_PROFILE":                   cfg.ProfileId,
		"DELEGATE_STORAGE_URL":               "https://app.harness.io",
		"DELEGATE_TYPE":                      "DOCKER",
		"DEPLOY_MODE":                        "KUBERNETES",
		"HELM_DESIRED_VERSION":               "",
		"HELM_PATH":                          "",
		"HELM3_PATH":                         "",
		"INSTALL_CLIENT_TOOLS_IN_BACKGROUND": "true",
		"JRE_VERSION":                        "1.8.0_242",
		"KUBECTL_PATH":                       "",
		"KUSTOMIZE_PATH":                     "",
		"MANAGER_HOST_AND_PORT":              "https://app.harness.io/gratis",
		"NO_PROXY":                           "",
		"OC_PATH":                            "",
		"POLL_FOR_TASKS":                     "false",
		"PROXY_HOST":                         "",
		"PROXY_MANAGER":                      "true",
		"PROXY_PASSWORD":                     "",
		"PROXY_PORT":                         "",
		"PROXY_SCHEME":                       "",
		"PROXY_USER":                         "",
		"REMOTE_WATCHER_URL_CDN":             "https://app.harness.io/public/shared/watchers/builds",
		"USE_CDN":                            "true",
		"VERSION_CHECK_DISABLED":             "false",
		"WATCHER_CHECK_LOCATION":             "current.version",
		"WATCHER_STORAGE_URL":                "https://app.harness.io/public/free/freemium/watchers",
	}
}

func getDefaultContainerConfig(cfg *DockerDelegateConfig) *container.Config {
	envConfig := getDefaultEnvConfig(cfg)

	// Merge user defined env vars with defaults
	if cfg.EnvVars != nil && len(cfg.EnvVars) > 0 {
		for k, v := range cfg.EnvVars {
			envConfig[k] = v
		}
	}

	return &container.Config{
		Image: cfg.Image,
		Env:   utils.MapToStringSlice(envConfig, "="),
	}
}

// Returns the container Id of the delegate
func RunDelegateContainer(ctx context.Context, cfg *DockerDelegateConfig) (string, error) {
	if cfg.Image == "" {
		cfg.Image = harnessDockerImage
	}

	var clientOpts client.Opt
	if cfg.ClientOptions != nil {
		clientOpts = cfg.ClientOptions
	} else {
		clientOpts = client.FromEnv
	}

	var containerConfig *container.Config
	if cfg.ContainerConfig != nil {
		containerConfig = cfg.ContainerConfig
	} else {
		containerConfig = getDefaultContainerConfig(cfg)
	}

	cli, err := client.NewClientWithOpts(clientOpts)
	if err != nil {
		return "", errors.Wrap(err, "failed to create docker client")
	}

	log.Infof("Pulling docker image %s", cfg.Image)
	reader, err := cli.ImagePull(ctx, cfg.Image, types.ImagePullOptions{All: true})
	if err != nil {
		return "", errors.Wrap(err, "failed to pull docker image")
	}
	io.Copy(os.Stdout, reader)

	log.Infof("Creating docker container %s", cfg.ContainerName)
	cont, err := cli.ContainerCreate(ctx, containerConfig, cfg.HostConfig, cfg.NetworkingConfig, cfg.Platform, cfg.ContainerName)
	if err != nil {
		return "", errors.Wrap(err, "failed to create docker container")
	}

	log.Infof("Starting docker container %s", cont.ID)
	err = cli.ContainerStart(ctx, cont.ID, cfg.ContainerStartOptions)
	if err != nil {
		return "", errors.Wrap(err, "failed to start delegate container: %s")
	}

	return cont.ID, nil
}

func RemoveDelegateContainer(ctx context.Context, cfg *DockerDelegateConfig, containerId string) error {
	cli, err := client.NewClientWithOpts(cfg.ClientOptions)
	if err != nil {
		return errors.Wrap(err, "failed to create docker client")
	}

	log.Infof("Stopping docker container %s", containerId)
	err = cli.ContainerStop(ctx, containerId, nil)
	if err != nil {
		return errors.Wrap(err, "failed to stop delegate container")
	}

	log.Infof("Removing docker container %s", containerId)
	err = cli.ContainerRemove(ctx, containerId, cfg.ContainerRemoveOptions)
	if err != nil {
		return errors.Wrap(err, "failed to remove delegate container")
	}

	return nil
}
