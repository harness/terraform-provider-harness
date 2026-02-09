package registry

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// getParentRef returns the parent reference, using the provided parentRef or constructing it from account/org/project IDs
func getParentRef(accountID, orgID, projectID string, parentRef string) string {
	if parentRef != "" {
		return parentRef
	}
	return getRef(accountID, orgID, projectID)
}

// getRef constructs a hierarchical reference path by joining non-empty parameters with forward slashes
func getRef(params ...string) string {
	var result []string
	for _, param := range params {
		if param == "" {
			break
		}
		result = append(result, param)
	}
	return strings.Join(result, "/")
}

// expandStringSet converts a Terraform schema.Set into a slice of strings
func expandStringSet(s *schema.Set) []string {
	if s == nil {
		return nil
	}
	out := make([]string, 0, s.Len())
	for _, v := range s.List() {
		out = append(out, v.(string))
	}
	return out
}

// TestAccRegistryImportStateIdFunc returns a function that constructs the import ID (parent_ref/identifier) for registry resources in acceptance tests
func TestAccRegistryImportStateIdFunc(resourceName string) func(*terraform.State) (string, error) {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		parentRef := rs.Primary.Attributes["parent_ref"]
		identifier := rs.Primary.Attributes["identifier"]

		return fmt.Sprintf("%s/%s", parentRef, identifier), nil
	}
}

// Generates Terraform config for a virtual Docker registry without upstream proxies (local storage)
func TestAccResourceVirtualDockerRegistryNoUpstream(id string, accId string) string {
	return fmt.Sprintf(`
resource "harness_platform_har_registry" "test" {
  identifier   = "%[1]s"
  space_ref    = "%[2]s"
  package_type = "DOCKER"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "%[2]s"
}
`, id, accId)
}

// Generates Terraform config for a virtual Docker registry with one upstream proxy
func TestAccResourceVirtualDockerRegistryWithOneUpstream(id string, upstreamId string, accId string) string {
	return fmt.Sprintf(`
resource "harness_platform_har_registry" "upstream1" {
  identifier   = "%[2]s"
  space_ref    = "%[3]s"
  package_type = "DOCKER"

  config {
    type = "UPSTREAM"
    auth_type = "Anonymous"
    source = "Dockerhub"
  }
  parent_ref = "%[3]s"
}

resource "harness_platform_har_registry" "test" {
  identifier   = "%[1]s"
  space_ref    = "%[3]s"
  package_type = "DOCKER"

  config {
    type = "VIRTUAL"
    upstream_proxies = [harness_platform_har_registry.upstream1.identifier]
  }
  parent_ref = "%[3]s"
  
  depends_on = [harness_platform_har_registry.upstream1]
}
`, id, upstreamId, accId)
}

// Generates Terraform config for a virtual Docker registry with two upstream proxies
func TestAccResourceVirtualDockerRegistryWithTwoUpstreams(id string, upstream1 string, upstream2 string, accId string) string {
	return fmt.Sprintf(`
resource "harness_platform_har_registry" "upstream1" {
  identifier   = "%[2]s"
  space_ref    = "%[4]s"
  package_type = "DOCKER"

  config {
    type = "UPSTREAM"
    auth_type = "Anonymous"
    source = "Dockerhub"
  }
  parent_ref = "%[4]s"
}

resource "harness_platform_har_registry" "upstream2" {
  identifier   = "%[3]s"
  space_ref    = "%[4]s"
  package_type = "DOCKER"

  config {
    type = "UPSTREAM"
    auth_type = "Anonymous"
    source = "Dockerhub"
  }
  parent_ref = "%[4]s"
}

resource "harness_platform_har_registry" "test" {
  identifier   = "%[1]s"
  space_ref    = "%[4]s"
  package_type = "DOCKER"

  config {
    type = "VIRTUAL"
    upstream_proxies = [
      harness_platform_har_registry.upstream1.identifier,
      harness_platform_har_registry.upstream2.identifier
    ]
  }
  parent_ref = "%[4]s"
  
  depends_on = [
    harness_platform_har_registry.upstream1,
    harness_platform_har_registry.upstream2
  ]
}
`, id, upstream1, upstream2, accId)
}

// Generates Terraform config for a virtual Docker registry without patterns
func TestAccResourceVirtualDockerRegistryNoPatterns(id string, accId string) string {
	return fmt.Sprintf(`
resource "harness_platform_har_registry" "test" {
  identifier   = "%[1]s"
  space_ref    = "%[2]s"
  package_type = "DOCKER"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "%[2]s"
}
`, id, accId)
}

// Generates Terraform config with upstream registries created but not used by virtual registry
func TestAccResourceVirtualDockerRegistryWithUpstreamsButNotUsed(id string, upstream1 string, upstream2 string, accId string) string {
	return fmt.Sprintf(`
resource "harness_platform_har_registry" "upstream1" {
  identifier   = "%[2]s"
  space_ref    = "%[4]s"
  package_type = "DOCKER"

  config {
    type = "UPSTREAM"
    auth_type = "Anonymous"
    source = "Dockerhub"
  }
  parent_ref = "%[4]s"
}

resource "harness_platform_har_registry" "upstream2" {
  identifier   = "%[3]s"
  space_ref    = "%[4]s"
  package_type = "DOCKER"

  config {
    type = "UPSTREAM"
    auth_type = "Anonymous"
    source = "Dockerhub"
  }
  parent_ref = "%[4]s"
}

resource "harness_platform_har_registry" "test" {
  identifier   = "%[1]s"
  space_ref    = "%[4]s"
  package_type = "DOCKER"

  config {
    type = "VIRTUAL"
  }
  parent_ref = "%[4]s"
  
  depends_on = [
    harness_platform_har_registry.upstream1,
    harness_platform_har_registry.upstream2
  ]
}
`, id, upstream1, upstream2, accId)
}
