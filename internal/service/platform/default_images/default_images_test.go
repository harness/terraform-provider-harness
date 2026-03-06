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
Environment variables:

  DI_CI_TEST_DRY_RUN="true"        print generated configs and exit without running
  DI_CI_TEST_PROVIDER_BLOCK="true" prepend a terraform{} provider block (dev overrides only)
  DI_CI_TRACE_ATTR="true"          print actual attribute values observed during each check step

Run:

  DI_CI_TEST_DRY_RUN="false" \
  DI_CI_TEST_PROVIDER_BLOCK="true" \
  DI_CI_TRACE_ATTR="true" \
  go test -v ./internal/service/platform/default_images/... \
      -run TestAccDefaultImagesCILifecycle
*/

const ciLiteEngineField = "liteEngineTag"
const ciResourceAddr = "harness_platform_default_images.ci_lite_engine"
const ciCustomerDataAddr = "data.harness_platform_default_images.ci_customer"
const ciDefaultDataAddr = "data.harness_platform_default_images.ci_defaults"

// TestAccDefaultImagesCILifecycle runs four sequential steps against the real API:
//
//  1. Read defaults   – liteEngineTag must match harness/ci-lite-engine:<tag>
//  2. Create override – set liteEngineTag = harness/ci-lite-engine:TeethyTiger
//  3. Update override – change to harness/ci-lite-engine:RunningFox
//  4. Reset           – omit value; liteEngineTag must revert to a Harness default
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
					traceAttr(t, ciDefaultDataAddr, fmt.Sprintf("images.%s", ciLiteEngineField)),
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
					traceAttr(t, ciResourceAddr, "value"),
					traceAttr(t, ciCustomerDataAddr, fmt.Sprintf("images.%s", ciLiteEngineField)),
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
					traceAttr(t, ciResourceAddr, "value"),
					traceAttr(t, ciCustomerDataAddr, fmt.Sprintf("images.%s", ciLiteEngineField)),
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
					traceAttr(t, ciDefaultDataAddr, fmt.Sprintf("images.%s", ciLiteEngineField)),
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

// traceAttr returns a check func that logs the attribute value when DI_CI_TRACE_ATTR=true.
// It never fails – it is purely observational.
func traceAttr(t *testing.T, addr, attr string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrWith(addr, attr, func(val string) error {
		if os.Getenv("DI_CI_TRACE_ATTR") == "true" {
			t.Logf("info: %s . %s = %q", addr, attr, val)
		}
		return nil
	})
}

func diProviderBlock() string {
	if os.Getenv("DI_CI_TEST_PROVIDER_BLOCK") != "true" {
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
