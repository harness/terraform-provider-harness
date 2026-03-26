package split

import (
	"testing"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
)

func TestSplitAPIKeyResourceID(t *testing.T) {
	t.Parallel()
	if got := splitAPIKeyResourceID(nil); got != "" {
		t.Fatalf("nil: %q", got)
	}
	if got := splitAPIKeyResourceID(&splitsdk.KeyResponse{}); got != "" {
		t.Fatalf("empty: %q", got)
	}
	if got := splitAPIKeyResourceID(&splitsdk.KeyResponse{Id: "id-1"}); got != "id-1" {
		t.Fatalf("id: %q", got)
	}
	if got := splitAPIKeyResourceID(&splitsdk.KeyResponse{Key: "secretkey"}); got != "secretkey" {
		t.Fatalf("key fallback: %q", got)
	}
	if got := splitAPIKeyResourceID(&splitsdk.KeyResponse{Id: "id-1", Key: "k"}); got != "id-1" {
		t.Fatalf("prefer id: %q", got)
	}
}
