package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// The four tests below collectively form the regression suite for CCM-32336
// across all CCM cloud-cost connectors. The bug class is: a NextGen GET
// endpoint returning HTTP 500 for a deleted entity caused terraform plan to
// fail with "giving up after 11 attempt(s)". The connector Read() routes
// through helpers.HandleReadApiError which clears state on
// 404 + ENTITY_NOT_FOUND, so once the GET returns 404 the provider re-plans
// a create instead of erroring.
//
// Each test:
//  1. Creates the connector via terraform.
//  2. PreConfig: deletes the connector out-of-band via the SDK
//     (simulating a UI delete).
//  3. PlanOnly + ExpectNonEmptyPlan: refresh+plan must not error and must
//     show drift (the connector is gone -> plan to recreate).
//  4. Re-applies; the connector is created with a new uuid, identifier stays
//     the same since callers control the identifier.

// TestAccResourceConnectorAwsCC_CCM32336_OutOfBandDeleteRecreates is the
// CCM-32336 regression for the AWS Cloud Cost connector.
func TestAccResourceConnectorAwsCC_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_connector_awscc.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAwsCC(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig:          func() { deleteConnectorOOB(t, id) },
				Config:             testAccResourceConnectorAwsCC(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccResourceConnectorAwsCC(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

// TestAccResourceConnectorKubernetsCloudCost_CCM32336_OutOfBandDeleteRecreates
// is the CCM-32336 regression for the Kubernetes Cloud Cost connector.
func TestAccResourceConnectorKubernetsCloudCost_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	id := fmt.Sprintf("%s_%s", "KubernetsCloudCost_CCM32336", utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_connector_kubernetes_cloud_cost.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorKubernetsCloudCost(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig:          func() { deleteConnectorOOB(t, id) },
				Config:             testAccResourceConnectorKubernetsCloudCost(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccResourceConnectorKubernetsCloudCost(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

// TestAccResourceConnectorGCPCloudCost_CCM32336_OutOfBandDeleteRecreates
// is the CCM-32336 regression for the GCP Cloud Cost connector.
func TestAccResourceConnectorGCPCloudCost_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_connector_gcp_cloud_cost.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorGCPCloudCost(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig:          func() { deleteConnectorOOB(t, id) },
				Config:             testAccResourceConnectorGCPCloudCost(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccResourceConnectorGCPCloudCost(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

// TestAccResourceConnectorAzureCloudCost_CCM32336_OutOfBandDeleteRecreates
// is the CCM-32336 regression for the Azure Cloud Cost connector.
func TestAccResourceConnectorAzureCloudCost_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_connector_azure_cloud_cost.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAzureCloudCost(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig:          func() { deleteConnectorOOB(t, id) },
				Config:             testAccResourceConnectorAzureCloudCost(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccResourceConnectorAzureCloudCost(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

// deleteConnectorOOB calls the platform DeleteConnector API directly with the
// account-level scope used by the cloud-cost connector test fixtures, mimicking
// a UI deletion. Test fails if the API call returns an error (a clean 200 is
// what we need so that subsequent terraform refresh exercises the
// GET-after-delete path that CCM-32336 fixes).
func deleteConnectorOOB(t *testing.T, identifier string) {
	t.Helper()
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	if _, _, err := c.ConnectorsApi.DeleteConnector(
		ctx, c.AccountId, identifier,
		&nextgen.ConnectorsApiDeleteConnectorOpts{},
	); err != nil {
		t.Fatalf("CCM-32336: out-of-band delete of connector %q failed: %v", identifier, err)
	}
}
