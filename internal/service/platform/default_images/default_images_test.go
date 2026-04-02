package default_images_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

/*

DI_CI_TEST_PROVIDER_BLOCK="true" to be set only for dev testing

Run CI:

  DI_CI_TEST_DRY_RUN="false" \
  DI_CI_TEST_PROVIDER_BLOCK="true" \
  DI_CI_TRACE_ATTR="true" \
  go test -v ./internal/service/platform/default_images/... \
      -run TestAccDefaultImagesCILifecycle

Run IACM:

  DI_IACM_TEST_DRY_RUN="false" \
  DI_CI_TEST_PROVIDER_BLOCK="true" \
  DI_IACM_TRACE_ATTR="true" \
  go test -v ./internal/service/platform/default_images/... \
      -run TestAccDefaultImagesIACMLifecycle

Run IDP:

  DI_IDP_TEST_DRY_RUN="false" \
  DI_CI_TEST_PROVIDER_BLOCK="true" \
  DI_IDP_TRACE_ATTR="true" \
  go test -v ./internal/service/platform/default_images/... \
      -run TestAccDefaultImagesIDPLifecycle
*/

const ciLiteEngineField = "liteEngineTag"
const ciResourceAddr = "harness_platform_default_images.ci_lite_engine"
const ciCustomerDataAddr = "data.harness_platform_default_images.ci_customer"
const ciDefaultDataAddr = "data.harness_platform_default_images.ci_defaults"

func TestAccDefaultImagesCILifecycle(t *testing.T) {
	isDryRun := os.Getenv("DI_CI_TEST_DRY_RUN") == "true"

	configs := map[string]string{
		"1-read-defaults":      configCIReadDefaults(),
		"2-create-TeethyTiger": configCIWithOverride("harness/ci-lite-engine:TeethyTiger"),
		"3-update-RunningFox":  configCIWithOverride("harness/ci-lite-engine:RunningFox"),
		"4-reset":              configCIReset(),
	}

	if isDryRun {
		for _, key := range []string{
			"1-read-defaults",
			"2-create-TeethyTiger",
			"3-update-RunningFox",
			"4-reset",
		} {
			t.Logf("\n=== %s ===\n%s", key, configs[key])
		}
		return
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configs["1-read-defaults"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_CI_TRACE_ATTR", ciDefaultDataAddr, fmt.Sprintf("images.%s", ciLiteEngineField)),
					resource.TestMatchResourceAttr(
						ciDefaultDataAddr,
						fmt.Sprintf("images.%s", ciLiteEngineField),
						regexp.MustCompile(`^harness/ci-lite-engine:.+`),
					),
				),
			},
			{
				Config: configs["2-create-TeethyTiger"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_CI_TRACE_ATTR", ciResourceAddr, "value"),
					traceAttr(t, "DI_CI_TRACE_ATTR", ciCustomerDataAddr, fmt.Sprintf("images.%s", ciLiteEngineField)),
					resource.TestCheckResourceAttr(
						ciResourceAddr,
						"value",
						"harness/ci-lite-engine:TeethyTiger",
					),
					resource.TestCheckResourceAttr(
						ciCustomerDataAddr,
						fmt.Sprintf("images.%s", ciLiteEngineField),
						"harness/ci-lite-engine:TeethyTiger",
					),
				),
			},
			{
				Config: configs["3-update-RunningFox"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_CI_TRACE_ATTR", ciResourceAddr, "value"),
					traceAttr(t, "DI_CI_TRACE_ATTR", ciCustomerDataAddr, fmt.Sprintf("images.%s", ciLiteEngineField)),
					resource.TestCheckResourceAttr(
						ciResourceAddr,
						"value",
						"harness/ci-lite-engine:RunningFox",
					),
					resource.TestCheckResourceAttr(
						ciCustomerDataAddr,
						fmt.Sprintf("images.%s", ciLiteEngineField),
						"harness/ci-lite-engine:RunningFox",
					),
				),
			},
			{
				Config: configs["4-reset"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_CI_TRACE_ATTR", ciDefaultDataAddr, fmt.Sprintf("images.%s", ciLiteEngineField)),
					resource.TestCheckResourceAttr(
						ciResourceAddr,
						"value",
						"",
					),
					resource.TestMatchResourceAttr(
						ciDefaultDataAddr,
						fmt.Sprintf("images.%s", ciLiteEngineField),
						regexp.MustCompile(`^harness/ci-lite-engine:.+`),
					),
				),
			},
		},
	})
}

// traceAttr returns a check func that logs the attribute value when the given envVar is "true".
// It never fails – it is purely observational.
func traceAttr(t *testing.T, envVar, addr, attr string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrWith(addr, attr, func(val string) error {
		if os.Getenv(envVar) == "true" {
			t.Logf("info: %s . %s = %q", addr, attr, val)
		}
		return nil
	})
}

func diProviderBlock() string {
	if os.Getenv("DI_CI_TEST_PROVIDER_BLOCK") != "true" &&
		os.Getenv("DI_IACM_TEST_PROVIDER_BLOCK") != "true" &&
		os.Getenv("DI_IDP_TEST_PROVIDER_BLOCK") != "true" {
		return ""
	}
	return `terraform {
  required_providers {
    harness = {
      source  = "harness/harness"
      version = "0.4000.2"
    }
  }
}
`
}

func configCIReadDefaults() string {
	return diProviderBlock() + `
data "harness_platform_default_images" "ci_defaults" {
  kind = "ci"
}
`
}

func configCIWithOverride(value string) string {
	return diProviderBlock() + fmt.Sprintf(`
resource "harness_platform_default_images" "ci_lite_engine" {
  kind  = "ci"
  field = %q
  value = %q
}

data "harness_platform_default_images" "ci_customer" {
  kind       = "ci"
  type       = "customer"
  depends_on = [harness_platform_default_images.ci_lite_engine]
}
`, ciLiteEngineField, value)
}

func configCIReset() string {
	return diProviderBlock() + fmt.Sprintf(`
resource "harness_platform_default_images" "ci_lite_engine" {
  kind  = "ci"
  field = %q
}

data "harness_platform_default_images" "ci_defaults" {
  kind       = "ci"
  depends_on = [harness_platform_default_images.ci_lite_engine]
}
`, ciLiteEngineField)
}

const iacmField = "iacmAwsCdk"
const iacmResourceAddr = "harness_platform_default_images.iacm_field"
const iacmCustomerDataAddr = "data.harness_platform_default_images.iacm_customer"
const iacmDefaultDataAddr = "data.harness_platform_default_images.iacm_defaults"

func TestAccDefaultImagesIACMLifecycle(t *testing.T) {
	isDryRun := os.Getenv("DI_IACM_TEST_DRY_RUN") == "true"

	configs := map[string]string{
		"1-read-defaults":      configIACMReadDefaults(),
		"2-create-TeethyTiger": configIACMWithOverride(iacmField, "plugins/harness_aws_cdk:TeethyTiger"),
		"3-update-RunningFox":  configIACMWithOverride(iacmField, "plugins/harness_aws_cdk:RunningFox"),
		"4-reset":              configIACMReset(iacmField),
	}

	if isDryRun {
		for _, key := range []string{
			"1-read-defaults",
			"2-create-TeethyTiger",
			"3-update-RunningFox",
			"4-reset",
		} {
			t.Logf("\n=== %s ===\n%s", key, configs[key])
		}
		return
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configs["1-read-defaults"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_IACM_TRACE_ATTR", iacmDefaultDataAddr, fmt.Sprintf("images.%s", iacmField)),
					resource.TestMatchResourceAttr(
						iacmDefaultDataAddr,
						fmt.Sprintf("images.%s", iacmField),
						regexp.MustCompile(`^plugins/harness_aws_cdk:.+`),
					),
				),
			},
			{
				Config: configs["2-create-TeethyTiger"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_IACM_TRACE_ATTR", iacmResourceAddr, "value"),
					resource.TestCheckResourceAttr(
						iacmResourceAddr,
						"value",
						"plugins/harness_aws_cdk:TeethyTiger",
					),
				),
			},
			{
				Config: configs["3-update-RunningFox"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_IACM_TRACE_ATTR", iacmResourceAddr, "value"),
					resource.TestCheckResourceAttr(
						iacmResourceAddr,
						"value",
						"plugins/harness_aws_cdk:RunningFox",
					),
				),
			},
			{
				Config: configs["4-reset"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_IACM_TRACE_ATTR", iacmDefaultDataAddr, fmt.Sprintf("images.%s", iacmField)),
					resource.TestCheckResourceAttr(
						iacmResourceAddr,
						"value",
						"",
					),
					resource.TestMatchResourceAttr(
						iacmDefaultDataAddr,
						fmt.Sprintf("images.%s", iacmField),
						regexp.MustCompile(`^plugins/harness_aws_cdk:.+`),
					),
				),
			},
		},
	})
}

func configIACMReadDefaults() string {
	return diProviderBlock() + `
data "harness_platform_default_images" "iacm_defaults" {
  kind = "iacm"
}
`
}

func configIACMWithOverride(field, value string) string {
	return diProviderBlock() + fmt.Sprintf(`
resource "harness_platform_default_images" "iacm_field" {
  kind  = "iacm"
  field = %q
  value = %q
}
`, field, value)
}

func configIACMReset(field string) string {
	return diProviderBlock() + fmt.Sprintf(`
resource "harness_platform_default_images" "iacm_field" {
  kind  = "iacm"
  field = %q
}

data "harness_platform_default_images" "iacm_defaults" {
  kind       = "iacm"
  depends_on = [harness_platform_default_images.iacm_field]
}
`, field)
}

const idpField = "registerCatalog"
const idpResourceAddr = "harness_platform_default_images.idp_field"
const idpDefaultDataAddr = "data.harness_platform_default_images.idp_defaults"

func TestAccDefaultImagesIDPLifecycle(t *testing.T) {
	isDryRun := os.Getenv("DI_IDP_TEST_DRY_RUN") == "true"

	configs := map[string]string{
		"1-read-defaults":      configIDPReadDefaults(),
		"2-create-TeethyTiger": configIDPWithOverride(idpField, "harness/registercatalog:TeethyTiger"),
		"3-update-RunningFox":  configIDPWithOverride(idpField, "harness/registercatalog:RunningFox"),
		"4-reset":              configIDPReset(idpField),
	}

	if isDryRun {
		for _, key := range []string{
			"1-read-defaults",
			"2-create-TeethyTiger",
			"3-update-RunningFox",
			"4-reset",
		} {
			t.Logf("\n=== %s ===\n%s", key, configs[key])
		}
		return
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configs["1-read-defaults"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_IDP_TRACE_ATTR", idpDefaultDataAddr, fmt.Sprintf("images.%s", idpField)),
					resource.TestMatchResourceAttr(
						idpDefaultDataAddr,
						fmt.Sprintf("images.%s", idpField),
						regexp.MustCompile(`^harness/registercatalog:.+`),
					),
				),
			},
			{
				Config: configs["2-create-TeethyTiger"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_IDP_TRACE_ATTR", idpResourceAddr, "value"),
					resource.TestCheckResourceAttr(
						idpResourceAddr,
						"value",
						"harness/registercatalog:TeethyTiger",
					),
				),
			},
			{
				Config: configs["3-update-RunningFox"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_IDP_TRACE_ATTR", idpResourceAddr, "value"),
					resource.TestCheckResourceAttr(
						idpResourceAddr,
						"value",
						"harness/registercatalog:RunningFox",
					),
				),
			},
			{
				Config: configs["4-reset"],
				Check: resource.ComposeTestCheckFunc(
					traceAttr(t, "DI_IDP_TRACE_ATTR", idpDefaultDataAddr, fmt.Sprintf("images.%s", idpField)),
					resource.TestCheckResourceAttr(
						idpResourceAddr,
						"value",
						"",
					),
					resource.TestMatchResourceAttr(
						idpDefaultDataAddr,
						fmt.Sprintf("images.%s", idpField),
						regexp.MustCompile(`^harness/registercatalog:.+`),
					),
				),
			},
		},
	})
}

func configIDPReadDefaults() string {
	return diProviderBlock() + `
data "harness_platform_default_images" "idp_defaults" {
  kind = "idp"
}
`
}

func configIDPWithOverride(field, value string) string {
	return diProviderBlock() + fmt.Sprintf(`
resource "harness_platform_default_images" "idp_field" {
  kind  = "idp"
  field = %q
  value = %q
}
`, field, value)
}

func configIDPReset(field string) string {
	return diProviderBlock() + fmt.Sprintf(`
resource "harness_platform_default_images" "idp_field" {
  kind  = "idp"
  field = %q
}

data "harness_platform_default_images" "idp_defaults" {
  kind       = "idp"
  depends_on = [harness_platform_default_images.idp_field]
}
`, field)
}
