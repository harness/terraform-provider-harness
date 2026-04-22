package split

import (
	"testing"

	splitsdk "github.com/harness/harness-go-sdk/harness/split"
)

func TestParseSplitAPIKeyMetadataJSON(t *testing.T) {
	t.Parallel()
	body := []byte(`{"name":"k1","apiKeyType":"client_side","environments":[{"id":"env-1"}]}`)
	n, typ, env, ok, err := parseSplitAPIKeyMetadataJSON(body)
	if err != nil || !ok || n != "k1" || typ != "client_side" || env != "env-1" {
		t.Fatalf("got ok=%v n=%q typ=%q env=%q err=%v", ok, n, typ, env, err)
	}
	n, typ, env, ok, err = parseSplitAPIKeyMetadataJSON([]byte(`{"name":"x","type":"server_side","environmentIds":["e2"]}`))
	if err != nil || !ok || n != "x" || typ != "server_side" || env != "e2" {
		t.Fatalf("environmentIds: ok=%v n=%q typ=%q env=%q err=%v", ok, n, typ, env, err)
	}
	_, _, _, ok, err = parseSplitAPIKeyMetadataJSON([]byte(`{"name":"x","apiKeyType":"server_side"}`))
	if err != nil || ok {
		t.Fatalf("missing env: ok=%v err=%v", ok, err)
	}
}

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
