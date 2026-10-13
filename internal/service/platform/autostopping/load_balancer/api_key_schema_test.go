package load_balancer

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestApiKeyRequiredOnALBAndAzureGateway(t *testing.T) {
	resources := map[string]*schema.Resource{
		"aws_alb":       ResourceAwsALB(),
		"azure_gateway": ResourceAzureGateway(),
	}

	for name, resource := range resources {
		t.Run(name, func(t *testing.T) {
			if err := resource.InternalValidate(resource.Schema, true); err != nil {
				t.Fatalf("schema failed InternalValidate: %v", err)
			}
			apiKey := resource.Schema["api_key"]
			if apiKey == nil {
				t.Fatal("api_key is missing from schema")
			}
			if !apiKey.Required {
				t.Fatal("api_key must be required")
			}
			if !apiKey.Sensitive {
				t.Fatal("api_key must be sensitive")
			}
			if apiKey.ValidateFunc == nil {
				t.Fatal("api_key must reject an empty string")
			}
			if _, errs := apiKey.ValidateFunc("", "api_key"); len(errs) == 0 {
				t.Fatal("empty api_key must be rejected")
			}
			if _, errs := apiKey.ValidateFunc("key", "api_key"); len(errs) != 0 {
				t.Fatalf("non-empty api_key rejected: %v", errs)
			}
		})
	}
}

func TestBuildLoadBalancerSendsAPIKeyAsAccessToken(t *testing.T) {
	cases := []struct {
		name  string
		res   *schema.Resource
		raw   map[string]interface{}
		cloud string
		kind  string
	}{
		{
			name:  "aws_alb",
			res:   ResourceAwsALB(),
			cloud: "aws",
			kind:  "",
			raw: map[string]interface{}{
				"name":                              "alb",
				"cloud_connector_id":                "connector",
				"region":                            "us-east-1",
				"vpc":                               "vpc-1",
				"api_key":                           "alb-key",
				"delete_cloud_resources_on_destroy": true,
			},
		},
		{
			name:  "azure_gateway",
			res:   ResourceAzureGateway(),
			cloud: "azure",
			kind:  KindAppGateway,
			raw: map[string]interface{}{
				"name":                              "gateway",
				"cloud_connector_id":                "connector",
				"region":                            "westus2",
				"resource_group":                    "rg",
				"vpc":                               "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet",
				"subnet_id":                         "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet/subnets/snet",
				"azure_func_region":                 "westus2",
				"frontend_ip":                       "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Network/publicIPAddresses/ip",
				"sku_size":                          "sku2",
				"api_key":                           "azure-key",
				"delete_cloud_resources_on_destroy": false,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, tc.res.Schema, tc.raw)
			lb, err := buildLoadBalancer(d, "account", tc.cloud, tc.kind)
			if err != nil {
				t.Fatalf("buildLoadBalancer: %v", err)
			}
			if lb.AccessToken != tc.raw["api_key"] {
				t.Fatalf("AccessToken = %q, want %q", lb.AccessToken, tc.raw["api_key"])
			}
			if lb.Metadata != nil && lb.Metadata.ApiKey != "" {
				t.Fatalf("Metadata.ApiKey = %q, want empty for non-proxy load balancers", lb.Metadata.ApiKey)
			}
		})
	}
}
