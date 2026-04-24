package internal

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/split"
)

func TestSessionFMEWorkspaceCache_roundTrip(t *testing.T) {
	t.Parallel()
	var s Session
	w := split.Workspace{ID: "ws-1", Name: "workspace-one"}

	s.SetFMEWorkspace("orgA", "projB", w)
	got, ok := s.GetFMEWorkspace("orgA", "projB")
	if !ok || got.ID != "ws-1" || got.Name != "workspace-one" {
		t.Fatalf("GetFMEWorkspace(orgA, projB) = (%+v, %v), want workspace ws-1", got, ok)
	}

	_, ok = s.GetFMEWorkspace("orgX", "projB")
	if ok {
		t.Fatal("expected cache miss for different org")
	}
}

func TestSessionFMEWorkspaceCache_nilSession(t *testing.T) {
	t.Parallel()
	var s *Session
	s.SetFMEWorkspace("o", "p", split.Workspace{ID: "x"}) // must not panic
	if _, ok := s.GetFMEWorkspace("o", "p"); ok {
		t.Fatal("nil session should not report cache hit")
	}
}
