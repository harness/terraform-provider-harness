package convert

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
)

// ExpandProxyConfig expands the proxy configuration from Terraform schema
func ExpandProxyConfig(input []interface{}) (*svcdiscovery.DatabaseProxyConfiguration, error) {
	if len(input) == 0 || input[0] == nil {
		return nil, nil
	}

	cfg := input[0].(map[string]interface{})
	proxy := &svcdiscovery.DatabaseProxyConfiguration{}

	var err error

	proxy.HttpProxy, err = getString(cfg, "http_proxy", false)
	if err != nil {
		return nil, fmt.Errorf("proxy.http_proxy: %w", err)
	}

	proxy.HttpsProxy, err = getString(cfg, "https_proxy", false)
	if err != nil {
		return nil, fmt.Errorf("proxy.https_proxy: %w", err)
	}

	proxy.NoProxy, err = getString(cfg, "no_proxy", false)
	if err != nil {
		return nil, fmt.Errorf("proxy.no_proxy: %w", err)
	}

	proxy.Url, err = getString(cfg, "url", false)
	if err != nil {
		return nil, fmt.Errorf("proxy.url: %w", err)
	}

	return proxy, nil
}

// FlattenProxyConfig flattens the proxy configuration to Terraform schema
func FlattenProxyConfig(proxy *svcdiscovery.DatabaseProxyConfiguration) (map[string]interface{}, error) {
	if proxy == nil {
		return nil, nil
	}

	result := map[string]interface{}{
		"http_proxy":  proxy.HttpProxy,
		"https_proxy": proxy.HttpsProxy,
		"no_proxy":    proxy.NoProxy,
		"url":         proxy.Url,
	}

	return result, nil
}
