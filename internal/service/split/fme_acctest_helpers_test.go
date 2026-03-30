package split_test

import (
	"crypto/rand"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// testAccFMEAlphanum returns n characters from [a-z0-9] for Split naming constraints in acc tests.
func testAccFMEAlphanum(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	out := make([]byte, n)
	rb := make([]byte, n)
	_, _ = rand.Read(rb)
	for i := range out {
		out[i] = letters[int(rb[i])%len(letters)]
	}
	return string(out)
}

func fmeImportStateIDOrgProjectThird(resourceName, thirdAttr string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		org := rs.Primary.Attributes["org_id"]
		proj := rs.Primary.Attributes["project_id"]
		third := rs.Primary.Attributes[thirdAttr]
		if org == "" || proj == "" || third == "" {
			return "", fmt.Errorf("missing import id parts: org=%q proj=%q %s=%q", org, proj, thirdAttr, third)
		}
		return fmt.Sprintf("%s/%s/%s", org, proj, third), nil
	}
}

// fmeImportStateIDApiKey builds org_id/project_id/environment_id/api_key_type/name/key_id for harness_fme_api_key import.
func fmeImportStateIDApiKey(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		a := rs.Primary.Attributes
		org, proj := a["org_id"], a["project_id"]
		env, typ, name, keyID := a["environment_id"], a["api_key_type"], a["name"], a["key_id"]
		if org == "" || proj == "" || env == "" || typ == "" || name == "" || keyID == "" {
			return "", fmt.Errorf("missing api key import id parts: org=%q proj=%q env=%q type=%q name=%q key_id=%q", org, proj, env, typ, name, keyID)
		}
		return fmt.Sprintf("%s/%s/%s/%s/%s/%s", org, proj, env, typ, name, keyID), nil
	}
}

// fmeImportStateIDOrgProjectTTFourth builds org/project/traffic_type_id/<resource primary id> for attribute import.
func fmeImportStateIDOrgProjectTTFourth(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		org := rs.Primary.Attributes["org_id"]
		proj := rs.Primary.Attributes["project_id"]
		tt := rs.Primary.Attributes["traffic_type_id"]
		if org == "" || proj == "" || tt == "" || rs.Primary.ID == "" {
			return "", fmt.Errorf("missing import id parts")
		}
		return fmt.Sprintf("%s/%s/%s/%s", org, proj, tt, rs.Primary.ID), nil
	}
}

func fmeImportStatePrimaryID(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		id := rs.Primary.ID
		if id == "" {
			return "", fmt.Errorf("empty primary id for %s", resourceName)
		}
		return id, nil
	}
}
