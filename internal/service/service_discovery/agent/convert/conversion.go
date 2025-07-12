package convert

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
)

// ExpandAgentConfig converts Terraform configuration to the API model
func ExpandAgentConfig(config []interface{}) (*svcdiscovery.DatabaseAgentConfiguration, error) {
	if len(config) == 0 || config[0] == nil {
		return nil, nil
	}

	conf := config[0].(map[string]interface{})
	agentConfig := &svcdiscovery.DatabaseAgentConfiguration{}

	// Set basic configuration
	var err error
	agentConfig.CollectorImage, err = getString(conf, "collector_image", false)
	if err != nil {
		return nil, fmt.Errorf("collector_image: %w", err)
	}

	agentConfig.LogWatcherImage, err = getString(conf, "log_watcher_image", false)
	if err != nil {
		return nil, fmt.Errorf("log_watcher_image: %w", err)
	}

	agentConfig.SkipSecureVerify, err = getBool(conf, "skip_secure_verify", false)
	if err != nil {
		return nil, fmt.Errorf("skip_secure_verify: %w", err)
	}

	// Set image pull secrets
	if secrets, ok := conf["image_pull_secrets"].([]interface{}); ok && len(secrets) > 0 {
		agentConfig.ImagePullSecrets = make([]string, len(secrets))
		for i, secret := range secrets {
			if s, ok := secret.(string); ok {
				agentConfig.ImagePullSecrets[i] = s
			}
		}
	}

	// Expand Kubernetes configuration
	if k8s, ok := conf["kubernetes"].([]interface{}); ok && len(k8s) > 0 {
		k8sConfig, err := ExpandKubernetesConfig(k8s)
		if err != nil {
			return nil, fmt.Errorf("kubernetes: %w", err)
		}
		agentConfig.Kubernetes = k8sConfig
	}

	// Expand mTLS configuration
	if mtls, ok := conf["mtls"].([]interface{}); ok && len(mtls) > 0 {
		mtlsConfig, err := ExpandMtlsConfig(mtls)
		if err != nil {
			return nil, fmt.Errorf("mtls: %w", err)
		}
		agentConfig.Mtls = mtlsConfig
	}

	// Expand proxy configuration
	if proxy, ok := conf["proxy"].([]interface{}); ok && len(proxy) > 0 {
		proxyConfig, err := ExpandProxyConfig(proxy)
		if err != nil {
			return nil, fmt.Errorf("proxy: %w", err)
		}
		agentConfig.Proxy = proxyConfig
	}

	// Expand data collection configuration
	if data, ok := conf["data"].([]interface{}); ok && len(data) > 0 {
		dataConfig, err := ExpandDataCollectionConfig(data)
		if err != nil {
			return nil, fmt.Errorf("data: %w", err)
		}
		agentConfig.Data = dataConfig
	}

	return agentConfig, nil
}

// FlattenAgentConfig converts the API model to Terraform configuration
func FlattenAgentConfig(config *svcdiscovery.DatabaseAgentConfiguration) (map[string]interface{}, error) {
	if config == nil {
		return nil, nil
	}

	result := make(map[string]interface{})

	// Set basic configuration
	if config.CollectorImage != "" {
		result["collector_image"] = config.CollectorImage
	}

	if config.LogWatcherImage != "" {
		result["log_watcher_image"] = config.LogWatcherImage
	}

	if config.SkipSecureVerify {
		result["skip_secure_verify"] = config.SkipSecureVerify
	}

	// Set image pull secrets
	if len(config.ImagePullSecrets) > 0 {
		result["image_pull_secrets"] = config.ImagePullSecrets
	}

	// Flatten Kubernetes configuration
	if config.Kubernetes != nil {
		k8s, err := FlattenKubernetesConfig(config.Kubernetes)
		if err != nil {
			return nil, fmt.Errorf("flattening kubernetes config: %w", err)
		}
		if k8s != nil {
			result["kubernetes"] = []interface{}{k8s}
		}
	}

	// Flatten mTLS configuration
	if config.Mtls != nil {
		mtls, err := FlattenMtlsConfig(config.Mtls)
		if err != nil {
			return nil, fmt.Errorf("flattening mtls config: %w", err)
		}
		if mtls != nil {
			result["mtls"] = []interface{}{mtls}
		}
	}

	// Flatten proxy configuration
	if config.Proxy != nil {
		proxy, err := FlattenProxyConfig(config.Proxy)
		if err != nil {
			return nil, fmt.Errorf("flattening proxy config: %w", err)
		}
		if proxy != nil {
			result["proxy"] = []interface{}{proxy}
		}
	}

	// Flatten data collection configuration
	if config.Data != nil {
		data, err := FlattenDataCollectionConfig(config.Data)
		if err != nil {
			return nil, fmt.Errorf("flattening data config: %w", err)
		}
		if data != nil {
			result["data"] = []interface{}{data}
		}
	}

	return result, nil
}
