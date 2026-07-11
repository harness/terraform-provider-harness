package time_sleep_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

/*
TIME_SLEEP_TEST_PROVIDER_BLOCK="true" to be set only for dev testing

Run all time_sleep tests:

  TIME_SLEEP_TEST_DRY_RUN="false" \
  TIME_SLEEP_TEST_PROVIDER_BLOCK="true" \
  go test -v ./internal/service/platform/time_sleep/... \
      -run TestAccTimeSleepAll

Run a single case:

  TIME_SLEEP_TEST_DRY_RUN="false" \
  TIME_SLEEP_TEST_PROVIDER_BLOCK="true" \
  go test -v ./internal/service/platform/time_sleep/... \
      -run TestAccTimeSleepCreateOnly
*/

func TestAccTimeSleepAll(t *testing.T) {
	TestAccTimeSleepCreateOnly(t)
	if t.Failed() {
		t.Log("TestAccTimeSleepCreateOnly FAILED – stopping")
		return
	}

	TestAccTimeSleepDestroyOnly(t)
	if t.Failed() {
		t.Log("TestAccTimeSleepDestroyOnly FAILED – stopping")
		return
	}

	TestAccTimeSleepBothDurations(t)
	if t.Failed() {
		t.Log("TestAccTimeSleepBothDurations FAILED – stopping")
		return
	}

	TestAccTimeSleepTriggers(t)
	if t.Failed() {
		t.Log("TestAccTimeSleepTriggers FAILED – stopping")
	}
}

// TestAccTimeSleepCreateOnly verifies a resource with only create_duration set
// and that the actual wall-clock sleep of at least 2s was honoured.
func TestAccTimeSleepCreateOnly(t *testing.T) {
	isDryRun := os.Getenv("TIME_SLEEP_TEST_DRY_RUN") == "true"

	resourceAddr := "harness_time_sleep.test"
	cfg := timeSleepProviderBlock() + configTimeSleepCreateOnly("2s")

	if isDryRun {
		t.Logf("\n=== generated Terraform source ===\n%s", cfg)
		return
	}

	start := time.Now()
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceAddr, "create_duration", "2s"),
					resource.TestCheckResourceAttrSet(resourceAddr, "id"),
					logAttr(t, resourceAddr, "id"),
					checkElapsed(t, &start, 2*time.Second),
				),
			},
		},
	})
}

// TestAccTimeSleepDestroyOnly verifies a resource with only destroy_duration set
// and that the actual wall-clock sleep of at least 2s was honoured on destroy.
func TestAccTimeSleepDestroyOnly(t *testing.T) {
	isDryRun := os.Getenv("TIME_SLEEP_TEST_DRY_RUN") == "true"

	resourceAddr := "harness_time_sleep.test"
	cfg := timeSleepProviderBlock() + configTimeSleepDestroyOnly("2s")

	if isDryRun {
		t.Logf("\n=== generated Terraform source ===\n%s", cfg)
		return
	}

	var destroyStart time.Time
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy: func(_ *terraform.State) error {
			elapsed := time.Since(destroyStart)
			if destroyStart.IsZero() {
				return fmt.Errorf("destroyStart was never set — destroy timing cannot be verified")
			}
			if elapsed < 2*time.Second {
				return fmt.Errorf("destroy sleep too short: elapsed %s, expected at least 2s", elapsed)
			}
			t.Logf("destroy elapsed: %s", elapsed)
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceAddr, "destroy_duration", "2s"),
					resource.TestCheckResourceAttrSet(resourceAddr, "id"),
					func(_ *terraform.State) error {
						destroyStart = time.Now()
						return nil
					},
				),
			},
		},
	})
}

// TestAccTimeSleepBothDurations verifies create + destroy durations, and that
// updating durations alone does not replace the resource (id is preserved).
func TestAccTimeSleepBothDurations(t *testing.T) {
	isDryRun := os.Getenv("TIME_SLEEP_TEST_DRY_RUN") == "true"

	resourceAddr := "harness_time_sleep.test"

	cfgCreate := timeSleepProviderBlock() + configTimeSleepBoth("2s", "2s")
	cfgUpdate := timeSleepProviderBlock() + configTimeSleepBoth("3s", "2s")

	if isDryRun {
		t.Logf("\n=== step 1 ===\n%s", cfgCreate)
		t.Logf("\n=== step 2 ===\n%s", cfgUpdate)
		return
	}

	start := time.Now()
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfgCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceAddr, "create_duration", "2s"),
					resource.TestCheckResourceAttr(resourceAddr, "destroy_duration", "2s"),
					resource.TestCheckResourceAttrSet(resourceAddr, "id"),
					logAttr(t, resourceAddr, "id"),
					checkElapsed(t, &start, 2*time.Second),
				),
			},
			{
				// Updating durations only must NOT replace the resource — id must stay the same.
				Config: cfgUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceAddr, "create_duration", "3s"),
					resource.TestCheckResourceAttr(resourceAddr, "destroy_duration", "2s"),
					resource.TestCheckResourceAttrSet(resourceAddr, "id"),
				),
			},
		},
	})
}

// TestAccTimeSleepTriggers verifies that a change in triggers forces resource replacement
// (new id is assigned after replacement).
func TestAccTimeSleepTriggers(t *testing.T) {
	isDryRun := os.Getenv("TIME_SLEEP_TEST_DRY_RUN") == "true"

	resourceAddr := "harness_time_sleep.test"

	cfgV1 := timeSleepProviderBlock() + configTimeSleepWithTrigger("2s", "v1")
	cfgV2 := timeSleepProviderBlock() + configTimeSleepWithTrigger("2s", "v2")

	if isDryRun {
		t.Logf("\n=== step 1 ===\n%s", cfgV1)
		t.Logf("\n=== step 2 (trigger change) ===\n%s", cfgV2)
		return
	}

	var idBeforeReplace string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfgV1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceAddr, "triggers.key", "v1"),
					resource.TestCheckResourceAttrSet(resourceAddr, "id"),
					captureAttr(resourceAddr, "id", &idBeforeReplace),
				),
			},
			{
				// Trigger change must force replacement — id must differ from the first step.
				Config: cfgV2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceAddr, "triggers.key", "v2"),
					resource.TestCheckResourceAttrSet(resourceAddr, "id"),
					checkAttrChanged(resourceAddr, "id", &idBeforeReplace),
				),
			},
		},
	})
}

// TestAccTimeSleepMinutesValidation verifies "5m" passes schema validation — no actual sleep.
func TestAccTimeSleepMinutesValidation(t *testing.T) {
	cfg := timeSleepProviderBlock() + configTimeSleepCreateOnly("5m")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             cfg,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// TestAccTimeSleepInvalidDuration verifies that an invalid duration value is rejected at plan time.
func TestAccTimeSleepInvalidDuration(t *testing.T) {
	isDryRun := os.Getenv("TIME_SLEEP_TEST_DRY_RUN") == "true"

	cfg := timeSleepProviderBlock() + configTimeSleepCreateOnly("not-a-duration")

	if isDryRun {
		t.Logf("\n=== generated Terraform source ===\n%s", cfg)
		return
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      cfg,
				ExpectError: regexp.MustCompile(`must be a positive number followed by ms, s, m, or h`),
			},
		},
	})
}

// checkElapsed fails if less than minDuration has elapsed since *start.
// start is a pointer so it can be set just before resource.UnitTest is called,
// capturing only the time the provider spent (including the sleep).
func checkElapsed(t *testing.T, start *time.Time, minDuration time.Duration) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		elapsed := time.Since(*start)
		t.Logf("elapsed since test start: %s (expected at least %s)", elapsed, minDuration)
		if elapsed < minDuration {
			return fmt.Errorf("sleep too short: elapsed %s, expected at least %s", elapsed, minDuration)
		}
		return nil
	}
}

func logAttr(t *testing.T, resourceAddr, attr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceAddr]
		if !ok {
			return fmt.Errorf("resource %q not found in state", resourceAddr)
		}
		t.Logf("%s.%s = %q", resourceAddr, attr, rs.Primary.Attributes[attr])
		return nil
	}
}

func captureAttr(resourceAddr, attr string, dest *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceAddr]
		if !ok {
			return fmt.Errorf("resource %q not found in state", resourceAddr)
		}
		*dest = rs.Primary.Attributes[attr]
		return nil
	}
}

func checkAttrChanged(resourceAddr, attr string, before *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceAddr]
		if !ok {
			return fmt.Errorf("resource %q not found in state", resourceAddr)
		}
		current := rs.Primary.Attributes[attr]
		if current == *before {
			return fmt.Errorf("expected %s.%s to change after resource replacement, but it stayed %q", resourceAddr, attr, current)
		}
		return nil
	}
}

func timeSleepProviderBlock() string {
	if os.Getenv("TIME_SLEEP_TEST_PROVIDER_BLOCK") != "true" {
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

func configTimeSleepCreateOnly(createDuration string) string {
	return fmt.Sprintf(`
resource "harness_time_sleep" "test" {
  create_duration = %q
}
`, createDuration)
}

func configTimeSleepDestroyOnly(destroyDuration string) string {
	return fmt.Sprintf(`
resource "harness_time_sleep" "test" {
  destroy_duration = %q
}
`, destroyDuration)
}

func configTimeSleepBoth(createDuration, destroyDuration string) string {
	return fmt.Sprintf(`
resource "harness_time_sleep" "test" {
  create_duration  = %q
  destroy_duration = %q
}
`, createDuration, destroyDuration)
}

func configTimeSleepWithTrigger(createDuration, triggerValue string) string {
	return fmt.Sprintf(`
resource "harness_time_sleep" "test" {
  create_duration = %q
  triggers = {
    key = %q
  }
}
`, createDuration, triggerValue)
}
