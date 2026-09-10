package monitored_service

import "testing"

// Regression test for CDS-131295: gcpProjectId set in the Stackdriver health source
// spec must be carried through into the nextgen.StackdriverMetricHealthSource struct
// that gets sent to the backend, instead of being silently dropped.
func TestGetStackDriverHealthSource_GcpProjectId(t *testing.T) {
	hs := map[string]interface{}{
		"connectorRef":      "account.TestGCP",
		"gcpProjectId":      "prj-ch-sharedservices-test",
		"metricDefinitions": []interface{}{},
	}

	result := getStackDriverHealthSource(hs)

	if result.GcpProjectId != "prj-ch-sharedservices-test" {
		t.Errorf("expected GcpProjectId to be %q, got %q", "prj-ch-sharedservices-test", result.GcpProjectId)
	}
}

func TestGetStackDriverHealthSource_GcpProjectIdOmitted(t *testing.T) {
	hs := map[string]interface{}{
		"connectorRef":      "account.TestGCP",
		"metricDefinitions": []interface{}{},
	}

	result := getStackDriverHealthSource(hs)

	if result.GcpProjectId != "" {
		t.Errorf("expected GcpProjectId to be empty when omitted, got %q", result.GcpProjectId)
	}
}
