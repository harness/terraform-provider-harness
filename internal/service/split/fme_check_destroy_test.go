package split_test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	splitapi "github.com/harness/harness-go-sdk/harness/split"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	splitpkg "github.com/harness/terraform-provider-harness/internal/service/split"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// fmeSplitAPIErrIsNotFound returns true when err indicates the Split object no longer exists.
func fmeSplitAPIErrIsNotFound(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "404") ||
		strings.Contains(s, "Not Found") ||
		strings.Contains(s, "not found") ||
		strings.Contains(s, http.StatusText(http.StatusNotFound))
}

// fmeSplitErrLooksTransient is true for likely-retryable Split / gateway errors during acc test polling.
func fmeSplitErrLooksTransient(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "500") ||
		strings.Contains(s, "502") ||
		strings.Contains(s, "503") ||
		strings.Contains(s, "429")
}

func fmeStateGet(resourceName string, state *terraform.State) (*terraform.ResourceState, error) {
	rs := state.RootModule().Resources[resourceName]
	if rs == nil {
		return nil, fmt.Errorf("resource not found in state: %s", resourceName)
	}
	if rs.Primary == nil {
		return nil, fmt.Errorf("resource %s has no primary instance", resourceName)
	}
	return rs, nil
}

func fmeSplitClient(ctx context.Context) (*splitapi.APIClient, context.Context) {
	sess := acctest.TestAccGetApiClientFromProvider()
	return sess.GetSplitClientWithContext(ctx)
}

// testAccFMECaptureAttr copies an attribute from the resource in state into dest (for tail-step Split checks).
func testAccFMECaptureAttr(resourceName, attr string, dest *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, err := fmeStateGet(resourceName, s)
		if err != nil {
			return err
		}
		v := rs.Primary.Attributes[attr]
		if v == "" {
			return fmt.Errorf("capture %s.%s: empty", resourceName, attr)
		}
		*dest = v
		return nil
	}
}

func fmeAssertEnvironmentGone(orgID, projectID, envID string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	client, cctx := fmeSplitClient(ctx)
	wsID, err := splitpkg.AccWorkspaceID(cctx, meta, orgID, projectID)
	if err != nil {
		return err
	}
	env, err := client.Environments.FindByID(wsID, envID)
	if err != nil {
		return err
	}
	if env != nil {
		return fmt.Errorf("Split environment still exists: id=%s name=%s", env.ID, env.Name)
	}
	return nil
}

func testAccFMEVerifyEnvironmentGone(orgID, projectID, envID string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertEnvironmentGone(orgID, projectID, envID)
	}
}

// fmeAssertFlagSetGone checks the flag set ID is absent from FlagSets.List(workspace) with polling.
// List avoids Split FindByID 500s seen after delete in acceptance runs.
func fmeAssertFlagSetGone(orgID, projectID, flagSetID string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	client, cctx := fmeSplitClient(ctx)
	wsID, err := splitpkg.AccWorkspaceID(cctx, meta, orgID, projectID)
	if err != nil {
		return err
	}
	const maxAttempts = 24
	const poll = 500 * time.Millisecond
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if attempt > 0 {
			time.Sleep(poll)
		}
		list, err := client.FlagSets.List(wsID)
		if err != nil {
			if fmeSplitAPIErrIsNotFound(err) {
				return nil
			}
			if fmeSplitErrLooksTransient(err) {
				continue
			}
			return fmt.Errorf("flag sets list: %w", err)
		}
		stillThere := false
		for i := range list {
			if list[i].ID == flagSetID {
				stillThere = true
				break
			}
		}
		if !stillThere {
			return nil
		}
	}
	return fmt.Errorf("Split flag set %q still listed in workspace after %d polls", flagSetID, maxAttempts)
}

func testAccFMEVerifyFlagSetGone(orgID, projectID, flagSetID string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertFlagSetGone(orgID, projectID, flagSetID)
	}
}

func fmeAssertTrafficTypeGone(orgID, projectID, ttID string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	client, cctx := fmeSplitClient(ctx)
	wsID, err := splitpkg.AccWorkspaceID(cctx, meta, orgID, projectID)
	if err != nil {
		return err
	}
	tt, err := client.TrafficTypes.FindByID(wsID, ttID)
	if err != nil {
		return err
	}
	if tt != nil {
		return fmt.Errorf("Split traffic type still exists: id=%s name=%s", tt.ID, tt.Name)
	}
	return nil
}

func testAccFMEVerifyTrafficTypeGone(orgID, projectID, ttID string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertTrafficTypeGone(orgID, projectID, ttID)
	}
}

func fmeAssertTrafficTypeAttributeGone(orgID, projectID, ttID, attrID string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	client, cctx := fmeSplitClient(ctx)
	wsID, err := splitpkg.AccWorkspaceID(cctx, meta, orgID, projectID)
	if err != nil {
		return err
	}
	attr, err := client.Attributes.FindByID(wsID, ttID, attrID, nil)
	if err != nil {
		if fmeSplitAPIErrIsNotFound(err) {
			return nil
		}
		return err
	}
	if attr != nil {
		return fmt.Errorf("Split traffic type attribute still exists: id=%s", attr.ID)
	}
	return nil
}

func testAccFMEVerifyTrafficTypeAttributeGone(orgID, projectID, ttID, attrID string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertTrafficTypeAttributeGone(orgID, projectID, ttID, attrID)
	}
}

func fmeAssertSegmentGone(orgID, projectID, name string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	client, cctx := fmeSplitClient(ctx)
	wsID, err := splitpkg.AccWorkspaceID(cctx, meta, orgID, projectID)
	if err != nil {
		return err
	}
	seg, err := client.Segments.Get(wsID, name)
	if err != nil {
		if fmeSplitAPIErrIsNotFound(err) {
			return nil
		}
		return err
	}
	if seg != nil {
		return fmt.Errorf("Split segment still exists: name=%s", name)
	}
	return nil
}

func testAccFMEVerifySegmentGone(orgID, projectID, name string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertSegmentGone(orgID, projectID, name)
	}
}

func fmeAssertSegmentEnvAssociationInactive(orgID, projectID, envID, segName string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	active, err := splitpkg.AccSegmentEnvAssociationActive(ctx, meta, orgID, projectID, envID, segName)
	if err != nil {
		if fmeSplitAPIErrIsNotFound(err) {
			return nil
		}
		return err
	}
	if active {
		return fmt.Errorf("classic segment %q still active in environment %s", segName, envID)
	}
	return nil
}

func testAccFMEVerifySegmentEnvAssociationInactive(orgID, projectID, envID, segName string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertSegmentEnvAssociationInactive(orgID, projectID, envID, segName)
	}
}

func fmeAssertRuleBasedSegmentGone(orgID, projectID, name string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	client, cctx := fmeSplitClient(ctx)
	wsID, err := splitpkg.AccWorkspaceID(cctx, meta, orgID, projectID)
	if err != nil {
		return err
	}
	def, err := client.RuleBasedSegments.Get(wsID, name)
	if err != nil {
		if fmeSplitAPIErrIsNotFound(err) {
			return nil
		}
		return err
	}
	if def != nil {
		return fmt.Errorf("rule-based segment still exists: name=%s", name)
	}
	return nil
}

func testAccFMEVerifyRuleBasedSegmentGone(orgID, projectID, name string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertRuleBasedSegmentGone(orgID, projectID, name)
	}
}

func fmeAssertRuleBasedSegmentEnvAssocGone(orgID, projectID, envID, segName string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	present, err := splitpkg.AccRuleBasedSegmentInEnvironment(ctx, meta, orgID, projectID, envID, segName)
	if err != nil {
		if fmeSplitAPIErrIsNotFound(err) {
			return nil
		}
		return err
	}
	if present {
		return fmt.Errorf("rule-based segment %q still present in environment %s", segName, envID)
	}
	return nil
}

func testAccFMEVerifyRuleBasedSegmentEnvAssocGone(orgID, projectID, envID, segName string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertRuleBasedSegmentEnvAssocGone(orgID, projectID, envID, segName)
	}
}

func fmeAssertLargeSegmentGone(orgID, projectID, name string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	client, cctx := fmeSplitClient(ctx)
	wsID, err := splitpkg.AccWorkspaceID(cctx, meta, orgID, projectID)
	if err != nil {
		return err
	}
	ls, err := client.LargeSegments.Get(wsID, name)
	if err != nil {
		if fmeSplitAPIErrIsNotFound(err) {
			return nil
		}
		return err
	}
	if ls != nil {
		return fmt.Errorf("large segment still exists: name=%s id=%s", name, ls.ID)
	}
	return nil
}

func testAccFMEVerifyLargeSegmentGone(orgID, projectID, name string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertLargeSegmentGone(orgID, projectID, name)
	}
}

func fmeAssertLargeSegmentEnvAssocGone(orgID, projectID, envID, segName string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	present, err := splitpkg.AccLargeSegmentDefinedInEnvironment(ctx, meta, orgID, projectID, envID, segName)
	if err != nil {
		if fmeSplitAPIErrIsNotFound(err) {
			return nil
		}
		return err
	}
	if present {
		return fmt.Errorf("large segment %q still defined in environment %s", segName, envID)
	}
	return nil
}

func testAccFMEVerifyLargeSegmentEnvAssocGone(orgID, projectID, envID, segName string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertLargeSegmentEnvAssocGone(orgID, projectID, envID, segName)
	}
}

func fmeAssertFeatureFlagGone(orgID, projectID, name string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	client, cctx := fmeSplitClient(ctx)
	wsID, err := splitpkg.AccWorkspaceID(cctx, meta, orgID, projectID)
	if err != nil {
		return err
	}
	sp, err := client.Splits.Get(wsID, name)
	if err != nil {
		if fmeSplitAPIErrIsNotFound(err) {
			return nil
		}
		return err
	}
	if sp != nil {
		return fmt.Errorf("split still exists: name=%s id=%s", sp.Name, sp.ID)
	}
	return nil
}

func testAccFMEVerifyFeatureFlagGone(orgID, projectID, name string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertFeatureFlagGone(orgID, projectID, name)
	}
}

func fmeAssertFeatureFlagDefinitionGone(orgID, projectID, envID, flagName string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	client, cctx := fmeSplitClient(ctx)
	wsID, err := splitpkg.AccWorkspaceID(cctx, meta, orgID, projectID)
	if err != nil {
		return err
	}
	def, err := client.Splits.GetDefinition(wsID, flagName, envID)
	if err != nil {
		if fmeSplitAPIErrIsNotFound(err) {
			return nil
		}
		return err
	}
	if def != nil {
		return fmt.Errorf("split definition still exists for flag %q in env %s", flagName, envID)
	}
	return nil
}

func testAccFMEVerifyFeatureFlagDefinitionGone(orgID, projectID, envID, flagName string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertFeatureFlagDefinitionGone(orgID, projectID, envID, flagName)
	}
}

func fmeAssertEnvironmentSegmentKeysGone(orgID, projectID, envID, segName string) error {
	ctx := context.Background()
	meta := acctest.TestAccGetApiClientFromProvider()
	client, cctx := fmeSplitClient(ctx)
	wsID, err := splitpkg.AccWorkspaceID(cctx, meta, orgID, projectID)
	if err != nil {
		return err
	}
	keys, err := client.Environments.GetSegmentKeysAll(wsID, envID, segName)
	if err != nil {
		if fmeSplitAPIErrIsNotFound(err) {
			return nil
		}
		return err
	}
	if len(keys) > 0 {
		return fmt.Errorf("segment %q still has %d keys in environment %s after destroy", segName, len(keys), envID)
	}
	return nil
}

func testAccFMEVerifyEnvironmentSegmentKeysGone(orgID, projectID, envID, segName string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		return fmeAssertEnvironmentSegmentKeysGone(orgID, projectID, envID, segName)
	}
}
