package contractv1

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestName is a typed constant identifying one acceptance scenario.
type TestName string

const (
	/*
	   Verify that Terraform can provision a new remote resource.
	   Must check:
	     - Terraform apply successfully creates the resource.
	     - The Harness API contains the newly created resource.
	     - Terraform state contains the expected ID and attributes.
	     - The provider correctly handles the API creation response.
	*/
	TestCreateResource TestName = "CreateResourceTest"

	/*
	   Verify that the provider can reconstruct and refresh Terraform
	   state from an existing remote resource.
	   Must check:
	     - Terraform first creates the resource.
	     - A subsequent refresh invokes the provider's Read operation.
	     - The resource ID and attributes remain correct after refresh.
	     - Remote values are accurately reflected in Terraform state.
	*/
	TestReadResource TestName = "ReadResourceTest"

	/*
	   Verify that changes to mutable Terraform attributes are
	   correctly applied to an existing remote resource.
	   Must check:
	     - Terraform creates the initial resource.
	     - Updated HCL changes one or more mutable attributes.
	     - Terraform applies the changes successfully.
	     - The Harness resource reflects the updated values.
	     - The resource retains its ID for an in-place update.
	     - Terraform state reflects the updated configuration.
	*/
	TestUpdateResource TestName = "UpdateResourceTest"

	/*
	   Verify that Terraform can delete a managed remote resource
	   without leaving orphaned resources.
	   Must check:
	     - Terraform successfully creates the resource.
	     - Terraform destroy invokes the provider's Delete operation.
	     - The Harness API confirms the resource no longer exists.
	     - Terraform removes the resource from state.
	     - Cleanup completes without unexpected errors.
	*/
	TestDeleteResource TestName = "DeleteResourceTest"

	/*
	   Verify the complete lifecycle of a resource in one
	   continuous acceptance-test scenario.
	   Must check:
	     - Create provisions the resource successfully.
	     - Update modifies the existing resource correctly.
	     - Read returns the expected values after each operation.
	     - Delete removes the remote resource.
	     - No stale remote resource remains after destruction.
	*/
	TestCreateUpdateDeleteResource TestName = "CreateUpdateDeleteResourceTest"

	/*
	   Verify that an existing Harness resource can be adopted
	   into Terraform management without being recreated.
	   Must check:
	     - Terraform first creates a real remote resource.
	     - The provider accepts its valid import ID.
	     - Import reconstructs Terraform state from the remote API.
	     - Imported state matches the original Terraform state.
	     - Only explicitly ignored attributes are excluded from
	       import-state verification.
	     - Import does not create a duplicate remote resource.
	*/
	TestImportResource TestName = "ImportResourceTest"

	/*
	   Verify that applying the same Terraform configuration
	   repeatedly does not produce unnecessary changes.
	   Must check:
	     - The initial Terraform apply succeeds.
	     - A subsequent plan with unchanged HCL has no changes.
	     - A repeated apply does not modify the remote resource.
	     - No unexpected replacement or attribute drift occurs.
	*/
	TestIdempotencyResource TestName = "IdempotencyResourceTest"

	/*
	   Verify that Terraform detects resources deleted outside
	   Terraform, such as through the Harness UI or API.
	   Must check:
	     - Terraform creates and manages the resource.
	     - An external API operation deletes the remote resource.
	     - Terraform refresh detects that the resource is missing.
	     - The provider clears the missing resource from state.
	     - Terraform plans to recreate the resource when its
	       configuration is still present.
	*/
	TestExternalDeleteResource TestName = "ExternalDeleteResourceTest"

	/*
	   Verify that Terraform detects and handles remote changes
	   made outside Terraform.
	   Must check:
	     - Terraform creates the resource.
	     - An external API operation modifies a managed attribute.
	     - Terraform refresh detects the changed remote value.
	     - Terraform plans corrective changes when the remote value
	       differs from the desired configuration.
	     - Applying the plan restores the configured value,
	       unless the attribute is intentionally ignored.
	*/
	TestExternalUpdateResource TestName = "ExternalUpdateResourceTest"

	/*
	   Verify that changes to immutable resource attributes
	   cause Terraform to replace the resource rather than
	   attempting an unsupported in-place update.
	   Must check:
	     - Terraform creates the initial resource.
	     - A ForceNew attribute is changed in the HCL.
	     - Terraform plans resource replacement.
	     - The old resource is destroyed and a new one is created,
	       respecting lifecycle settings.
	     - The replacement has the expected remote identity
	       and attributes.
	     - No orphaned original resource remains.
	*/
	TestForceNewResource TestName = "ForceNewResourceTest"

	/*
	   Verify that invalid Terraform configurations are rejected
	   with useful validation errors.
	   Must check:
	     - Required attributes are validated.
	     - Invalid attribute values or combinations are rejected.
	     - Terraform returns the expected validation error.
	     - Invalid configuration does not create a remote resource.
	     - Unrelated errors do not cause the test to pass.
	*/
	TestInvalidConfigurationResource TestName = "InvalidConfigurationResourceTest"

	/*
	   Verify that the provider handles Harness API failures
	   safely and reports meaningful diagnostics.
	   Must check:
	     - A controlled scenario produces an expected API failure.
	     - The provider returns an appropriate diagnostic.
	     - Terraform does not report a false success.
	     - Provider state remains consistent with the remote result.
	     - Errors are propagated rather than silently discarded.
	     - The expected error is matched explicitly.
	*/
	TestAPIErrorResource TestName = "APIErrorResourceTest"

	/*
	   Verify correct behavior for remote operations that do not
	   complete immediately, including eventual consistency.
	   Must check:
	     - Terraform initiates the remote operation.
	     - The provider handles intermediate or pending states.
	     - Polling or retry logic follows the expected behavior.
	     - Terraform succeeds only after the required remote
	       completion condition is reached.
	     - The final state reflects the completed operation.
	     - Permanent failures are not mistaken for pending work.
	*/
	TestAsyncResource TestName = "AsyncResourceTest"

	/*
	   Verify that the provider respects operation timeouts
	   and handles operations that never complete.
	   Must check:
	     - A controlled remote operation exceeds its timeout.
	     - The provider stops waiting within the configured limit.
	     - Terraform returns an appropriate timeout diagnostic.
	     - The provider does not report a false success.
	     - State and cleanup behavior are correct for the actual
	       remote outcome, including partial completion.
	*/
	TestTimeoutResource TestName = "TimeoutResourceTest"

	/*
	   Verify that importing a correctly formatted ID for a
	   nonexistent Harness resource does not create phantom state.
	   Must check:
	     - The import ID is syntactically valid.
	     - The referenced remote resource does not exist.
	     - The provider handles the missing-resource response.
	     - Import does not report a successfully managed resource
	       that is absent remotely.
	     - No phantom resource remains in Terraform state.
	     - The outcome follows the provider's documented behavior,
	       whether that is an explicit error or a missing-resource
	       diagnostic from the import verification process.
	*/
	TestImportNonExistentResource TestName = "ImportNonExistentResourceTest"
)

// TestOrder defines the execution order for all acceptance scenarios.
var TestOrder = []TestName{
	TestCreateResource, TestReadResource, TestUpdateResource, TestDeleteResource,
	TestCreateUpdateDeleteResource, TestImportResource, TestIdempotencyResource,
	TestExternalDeleteResource, TestExternalUpdateResource, TestForceNewResource,
	TestInvalidConfigurationResource, TestAPIErrorResource, TestAsyncResource,
	TestTimeoutResource, TestImportNonExistentResource,
}

// VerifierDefinition holds Terraform-state check functions.
// Callbacks may also query the Harness API directly for independent verification.
// VerifyDeleted runs as TestCase.CheckDestroy after Terraform cleanup.
type VerifierDefinition struct {
	VerifyCreated resource.TestCheckFunc
	VerifyRead    resource.TestCheckFunc
	VerifyUpdated resource.TestCheckFunc
	VerifyDeleted resource.TestCheckFunc
	VerifyAsync   resource.TestCheckFunc
}

// ContractDefinition declares exactly one acceptance scenario.
// Supported=false requires a non-empty UnsupportedReason. Supported tests must supply
// the callbacks relevant to their scenario.
type ContractDefinition struct {
	TestName          TestName
	Supported         bool
	UnsupportedReason string
	Description       string
	IsRunParallel     bool
	GetTerraformConfig        func() string
	GetUpdatedTerraformConfig func() string
	Verifiers                 VerifierDefinition
	PreTest                   func() error
	PostTest                  func() error
	// ExternalAction changes or deletes the remote object outside Terraform.
	ExternalAction func() error
	// ResourceName is the Terraform resource address (e.g. "harness_platform_foo.test").
	ResourceName            string
	GetImportID             func() string
	ImportStateVerifyIgnore []string
	ExpectError             *regexp.Regexp
}

// DebugSettings controls test execution during local development and debugging only.
// These fields must never be committed with non-zero values — they exist solely to
// speed up the inner loop when iterating on a single scenario. CI ignores them because
// an empty DebugSettings is a no-op.
type DebugSettings struct {
	// RunOnlyTests lists the specific test names to execute; all others are skipped.
	// Leave nil or empty to run all declared tests normally.
	RunOnlyTests []TestName
}

// ResourceAccTestContractV1 is the top-level contract a resource test file fills in.
type ResourceAccTestContractV1 struct {
	Definitions      []ContractDefinition
	GetSerializerKey func() string
	GetRequiredEnv   func() []string
	TestCase         func(*testing.T) resource.TestCase
	DebugSettings    DebugSettings
}

// Definition looks up a scenario by name.
func (c ResourceAccTestContractV1) Definition(name TestName) (ContractDefinition, bool) {
	for _, d := range c.Definitions {
		if d.TestName == name {
			return d, true
		}
	}
	return ContractDefinition{}, false
}

// Run validates and executes all declared acceptance scenarios.
func (c ResourceAccTestContractV1) Run(t *testing.T) {
	t.Helper()
	if err := c.Validate(); err != nil {
		t.Fatalf("invalid acceptance contract: %v", err)
	}
	c.run(t)
}
