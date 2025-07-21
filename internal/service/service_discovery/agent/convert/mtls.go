package convert

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
)

// ExpandMtlsConfig expands the mTLS configuration from Terraform schema
func ExpandMtlsConfig(input []interface{}) (*svcdiscovery.DatabaseMtlsConfiguration, error) {
	if len(input) == 0 || input[0] == nil {
		return nil, nil
	}

	cfg := input[0].(map[string]interface{})
	mtls := &svcdiscovery.DatabaseMtlsConfiguration{}

	var err error

	mtls.CertPath, err = getString(cfg, "cert_path", false)
	if err != nil {
		return nil, fmt.Errorf("mtls.cert_path: %w", err)
	}

	mtls.KeyPath, err = getString(cfg, "key_path", false)
	if err != nil {
		return nil, fmt.Errorf("mtls.key_path: %w", err)
	}

	mtls.SecretName, err = getString(cfg, "secret_name", false)
	if err != nil {
		return nil, fmt.Errorf("mtls.secret_name: %w", err)
	}

	mtls.Url, err = getString(cfg, "url", false)
	if err != nil {
		return nil, fmt.Errorf("mtls.url: %w", err)
	}

	return mtls, nil
}

// FlattenMtlsConfig flattens the mTLS configuration to Terraform schema
func FlattenMtlsConfig(mtls *svcdiscovery.DatabaseMtlsConfiguration) (map[string]interface{}, error) {
	if mtls == nil {
		return nil, nil
	}

	result := map[string]interface{}{
		"cert_path":   mtls.CertPath,
		"key_path":    mtls.KeyPath,
		"secret_name": mtls.SecretName,
		"url":         mtls.Url,
	}

	return result, nil
}
