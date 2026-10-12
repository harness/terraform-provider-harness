package load_balancer_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// HARNESS_PLATFORM_API_KEY must be set for proxy tests (backend requires valid NG API key). If unset, tests are skipped.
const platformAPIKeyEnv = "HARNESS_PLATFORM_API_KEY"

// randAlnum returns a short random string with no underscore or space (for resource names).
func randAlnum(n int) string {
	s := strings.ToLower(utils.RandStringBytes(n))
	s = strings.ReplaceAll(s, "_", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}

// extractAttr captures a Terraform state attribute into dest for use in PreConfig (e.g. proxy id).
func extractAttr(resourceName, key string, dest *string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrWith(resourceName, key, func(value string) error {
		*dest = value
		return nil
	})
}

// waitForProxyReady polls until the autostopping proxy leaves provisioning or reaches a terminal state.
// Returns nil on "created"; error on "errored" or timeout (matches autostopping/rule tests).
// A transient NG outage is retried until the deadline. That includes the
// gateway body and the bodyless error left after CheckRetry exhausts.
func waitForProxyReady(proxyID string, timeout time.Duration) error {
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	deadline := time.Now().Add(timeout)

	var lastTransient error
	lastStatus := "unknown"
	for time.Now().Before(deadline) {
		resp, _, err := c.CloudCostAutoStoppingLoadBalancersApi.DescribeLoadBalancer(ctx, c.AccountId, proxyID, c.AccountId)
		if err != nil {
			if transientNGError(err) {
				lastTransient = err
				time.Sleep(5 * time.Second)
				continue
			}
			return fmt.Errorf("reading proxy %s: %w", proxyID, err)
		}
		if resp.Response == nil {
			return fmt.Errorf("proxy %s not found", proxyID)
		}
		lastStatus = resp.Response.Status
		switch lastStatus {
		case "created":
			return nil
		case "errored":
			return fmt.Errorf("proxy %s provisioning failed", proxyID)
		}
		time.Sleep(5 * time.Second)
	}
	if lastTransient != nil {
		return fmt.Errorf("timeout waiting %v for proxy %s to provision (last status %q): %w", timeout, proxyID, lastStatus, lastTransient)
	}
	return fmt.Errorf("timeout waiting %v for proxy %s to provision (last status %q)", timeout, proxyID, lastStatus)
}

func transientNGError(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	// The 401 body is still attached when the client returns it directly.
	if strings.Contains(msg, "Could not connect to NG") {
		return true
	}
	// CheckRetry consumes that 401. After RetryMax, go-retryablehttp closes the
	// body and returns "METHOD URL giving up after N attempt(s)" with no marker.
	return strings.Contains(msg, "giving up after") && strings.Contains(msg, "attempt(s)")
}

func TestTransientNGError(t *testing.T) {
	t.Parallel()

	if transientNGError(nil) {
		t.Fatal("nil error is not transient")
	}
	if !transientNGError(fmt.Errorf("Error validating API key: Could not connect to NG")) {
		t.Fatal("gateway body should be transient")
	}
	exhausted := fmt.Errorf("GET https://qa.harness.io/gateway/lw/api/accounts/acc/autostopping/loadbalancers/ap-1 giving up after 11 attempt(s)")
	if !transientNGError(exhausted) {
		t.Fatal("retry exhaustion should be transient")
	}
	if transientNGError(fmt.Errorf(`{"message":"Invalid credentials"}`)) {
		t.Fatal("a real unauthorized response is not transient")
	}
}

// cleanupStaleAWSProxies deletes any orphaned AWS autostopping proxies left by
// previous test runs. This releases cloud resources (including Elastic IPs) that
// would otherwise cause AddressLimitExceeded errors on subsequent runs.
func cleanupStaleAWSProxies(t *testing.T) {
	t.Helper()
	c, ctx := acctest.TestAccGetPlatformClientWithContext()

	resp, _, err := c.CloudCostAutoStoppingLoadBalancersApi.ListLoadBalancers(ctx, c.AccountId, awsProxyCloudConnectorID, c.AccountId, nil)
	if err != nil {
		t.Logf("cleanup: unable to list load balancers: %v", err)
		return
	}

	for _, ap := range resp.Response {
		if ap.Type_ != "aws" || ap.Kind != "autostopping_proxy" {
			continue
		}
		if !strings.HasPrefix(ap.Name, "terr-awsproxy-") {
			continue
		}
		t.Logf("cleanup: deleting stale AWS proxy %s (%s, status=%s)", ap.Id, ap.Name, ap.Status)
		_, _ = c.CloudCostAutoStoppingLoadBalancersApi.DeleteLoadBalancer(ctx, nextgen.DeleteAccessPointPayload{
			Ids:           []string{ap.Id},
			WithResources: true,
		}, c.AccountId, c.AccountId)
	}
}

func testGetLoadBalancer(resourceName string, state *terraform.State) (*nextgen.AccessPoint, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.CloudCostAutoStoppingLoadBalancersApi.DescribeLoadBalancer(ctx, c.AccountId, id, c.AccountId)

	if err != nil {
		return nil, err
	}

	if resp.Response == nil {
		return nil, nil
	}

	return resp.Response, nil
}
