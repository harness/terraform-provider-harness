package split_test

import (
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// testAccRandomFlagSetSuffix returns n lowercase alphanumeric characters (no underscores).
// Split names allow underscores but RandStringBytes can produce a trailing "_", which may not match
// how Split normalizes or lists names; keep the random segment strictly [a-z0-9] for stable acc tests.
func testAccRandomFlagSetSuffix(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	out := make([]byte, n)
	rb := make([]byte, n)
	_, _ = rand.Read(rb)
	for i := range out {
		out[i] = letters[int(rb[i])%len(letters)]
	}
	return string(out)
}

// TestAccFMEFlagSet_dataSourceMatchesResource creates a flag set and reads it with the data source in one apply.
// Split's list-by-name API can lag after create; findFlagSetByNameWithRetry in the provider allows time for consistency.
func TestAccFMEFlagSet_dataSourceMatchesResource(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping FME flag set acceptance test in short mode")
	}
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	fsName := fmt.Sprintf("tf_fs_%s", testAccRandomFlagSetSuffix(8))
	res := "harness_fme_flag_set.test"
	ds := "data.harness_fme_flag_set.lookup"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFMEFlagSetWithDataSource(id, fsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(res, "flag_set_id"),
					resource.TestCheckResourceAttrPair(ds, "flag_set_id", res, "flag_set_id"),
					resource.TestCheckResourceAttrPair(ds, "id", res, "id"),
					resource.TestCheckResourceAttr(ds, "name", fsName),
				),
			},
		},
	})
}

func testAccFMEFlagSetWithDataSource(id, flagSetName string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		org_id     = harness_platform_organization.test.id
		name       = "%[1]s"
	}

	resource "harness_fme_flag_set" "test" {
		org_id      = harness_platform_organization.test.id
		project_id  = harness_platform_project.test.id
		name        = "%[2]s"
		description = "acceptance test flag set"
	}

	data "harness_fme_flag_set" "lookup" {
		depends_on = [harness_fme_flag_set.test]
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = harness_fme_flag_set.test.name
	}
	`, id, flagSetName)
}
