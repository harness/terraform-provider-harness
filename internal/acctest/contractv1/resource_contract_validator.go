package contractv1

import (
	"fmt"
	"strings"
)

// Validate applies v1 validation rules to the contract.
func (c ResourceAccTestContractV1) Validate() error {
	if c.GetSerializerKey == nil || c.GetRequiredEnv == nil {
		return fmt.Errorf("GetSerializerKey and GetRequiredEnv are required")
	}
	if c.TestCase == nil {
		return fmt.Errorf("TestCase is required")
	}
	known := map[TestName]bool{}
	for _, n := range TestOrder {
		known[n] = true
	}
	seen := map[TestName]bool{}
	for _, d := range c.Definitions {
		if !known[d.TestName] {
			return fmt.Errorf("unknown test %q", d.TestName)
		}
		if seen[d.TestName] {
			return fmt.Errorf("duplicate test declaration %q", d.TestName)
		}
		seen[d.TestName] = true
		if !d.Supported {
			if strings.TrimSpace(d.UnsupportedReason) == "" {
				return fmt.Errorf("%s: unsupported test requires UnsupportedReason", d.TestName)
			}
			continue
		}
		if err := d.validateSupported(); err != nil {
			return err
		}
	}
	for _, n := range TestOrder {
		if !seen[n] {
			return fmt.Errorf("missing declaration for %q", n)
		}
	}
	return nil
}

func (d ContractDefinition) validateSupported() error {
	require := func(ok bool, field string) error {
		if !ok {
			return fmt.Errorf("%s: %s required", d.TestName, field)
		}
		return nil
	}
	if err := require(d.GetTerraformConfig != nil, "GetTerraformConfig"); err != nil {
		return err
	}
	switch d.TestName {
	case TestCreateResource:
		return validateCreateResourceTest(d, require)
	case TestReadResource:
		return validateReadResourceTest(d, require)
	case TestUpdateResource:
		return validateUpdateResourceTest(d, require)
	case TestDeleteResource:
		return validateDeleteResourceTest(d, require)
	case TestCreateUpdateDeleteResource:
		return validateCreateUpdateDeleteResourceTest(d, require)
	case TestImportResource:
		return validateImportResourceTest(d, require)
	case TestExternalDeleteResource, TestExternalUpdateResource:
		return validateExternalActionResourceTest(d, require)
	case TestForceNewResource:
		return validateForceNewResourceTest(d, require)
	case TestInvalidConfigurationResource, TestTimeoutResource:
		return validateExpectErrorResourceTest(d, require)
	case TestAPIErrorResource:
		return validateAPIErrorResourceTest(d, require)
	case TestAsyncResource:
		return validateAsyncResourceTest(d, require)
	case TestImportNonExistentResource:
		return validateImportNonExistentResourceTest(d, require)
	}
	if len(d.ImportStateVerifyIgnore) > 0 && d.TestName != TestImportResource {
		return fmt.Errorf("%s: ImportStateVerifyIgnore only applies to ImportResource", d.TestName)
	}
	return nil
}

type requireFn func(ok bool, field string) error

func validateCreateResourceTest(d ContractDefinition, require requireFn) error {
	return require(d.Verifiers.VerifyCreated != nil, "Verifiers.VerifyCreated")
}

func validateReadResourceTest(d ContractDefinition, require requireFn) error {
	return require(d.Verifiers.VerifyRead != nil, "Verifiers.VerifyRead")
}

func validateUpdateResourceTest(d ContractDefinition, require requireFn) error {
	if err := require(d.GetUpdatedTerraformConfig != nil, "GetUpdatedTerraformConfig"); err != nil {
		return err
	}
	return require(d.Verifiers.VerifyUpdated != nil, "Verifiers.VerifyUpdated")
}

func validateDeleteResourceTest(d ContractDefinition, require requireFn) error {
	return require(d.Verifiers.VerifyDeleted != nil, "Verifiers.VerifyDeleted")
}

func validateCreateUpdateDeleteResourceTest(d ContractDefinition, require requireFn) error {
	for _, item := range []struct {
		ok   bool
		name string
	}{
		{d.GetUpdatedTerraformConfig != nil, "GetUpdatedTerraformConfig"},
		{d.Verifiers.VerifyCreated != nil, "Verifiers.VerifyCreated"},
		{d.Verifiers.VerifyUpdated != nil, "Verifiers.VerifyUpdated"},
		{d.Verifiers.VerifyDeleted != nil, "Verifiers.VerifyDeleted"},
	} {
		if err := require(item.ok, item.name); err != nil {
			return err
		}
	}
	return nil
}

func validateImportResourceTest(d ContractDefinition, require requireFn) error {
	return require(d.ResourceName != "", "ResourceName")
}

func validateExternalActionResourceTest(d ContractDefinition, require requireFn) error {
	return require(d.ExternalAction != nil, "ExternalAction")
}

func validateForceNewResourceTest(d ContractDefinition, require requireFn) error {
	if err := require(d.GetUpdatedTerraformConfig != nil, "GetUpdatedTerraformConfig"); err != nil {
		return err
	}
	return require(d.Verifiers.VerifyUpdated != nil, "Verifiers.VerifyUpdated")
}

func validateExpectErrorResourceTest(d ContractDefinition, require requireFn) error {
	return require(d.ExpectError != nil, "ExpectError")
}

func validateAPIErrorResourceTest(d ContractDefinition, require requireFn) error {
	if err := require(d.ExpectError != nil, "ExpectError"); err != nil {
		return err
	}
	return require(d.PreTest != nil, "PreTest (arrange API error)")
}

func validateAsyncResourceTest(d ContractDefinition, require requireFn) error {
	return require(d.Verifiers.VerifyAsync != nil, "Verifiers.VerifyAsync")
}

func validateImportNonExistentResourceTest(d ContractDefinition, require requireFn) error {
	for _, item := range []struct {
		ok   bool
		name string
	}{
		{d.ResourceName != "", "ResourceName"},
		{d.GetImportID != nil, "GetImportID"},
		{d.ExpectError != nil, "ExpectError"},
	} {
		if err := require(item.ok, item.name); err != nil {
			return err
		}
	}
	return nil
}
