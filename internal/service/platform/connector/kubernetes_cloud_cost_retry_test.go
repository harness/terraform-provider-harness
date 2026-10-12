package connector

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func TestIsTransientConnectorValidationError(t *testing.T) {
	t.Parallel()

	if !isTransientConnectorValidationError(diag.Errorf("Invalid request: Internal error while validating connector")) {
		t.Fatal("internal connector validation error should be retried")
	}
	if isTransientConnectorValidationError(diag.Errorf("Invalid request: connector identifier already exists")) {
		t.Fatal("a real validation error should not be retried")
	}
	if isTransientConnectorValidationError(nil) {
		t.Fatal("nil diagnostics are not a transient validation error")
	}
}
