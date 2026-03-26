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
