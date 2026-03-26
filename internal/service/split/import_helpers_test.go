package split_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/service/split"
)

func TestParseImportID3(t *testing.T) {
	t.Parallel()
	org, proj, third, err := split.ParseImportID3("org1/proj1/seg_a")
	if err != nil {
		t.Fatal(err)
	}
	if org != "org1" || proj != "proj1" || third != "seg_a" {
		t.Fatalf("got %q %q %q", org, proj, third)
	}
	_, _, _, err = split.ParseImportID3("bad")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseImportID4(t *testing.T) {
	t.Parallel()
	a, b, c, d, err := split.ParseImportID4("o/p/e/flag")
	if err != nil {
		t.Fatal(err)
	}
	if a != "o" || b != "p" || c != "e" || d != "flag" {
		t.Fatalf("got %q %q %q %q", a, b, c, d)
	}
	_, _, _, _, err = split.ParseImportID4("a/b/c")
	if err == nil {
		t.Fatal("expected error")
	}
	_, _, _, _, err = split.ParseImportID4("a/b/c/d/e")
	if err == nil {
		t.Fatal("expected error for 5 segments")
	}
}

func TestParseImportID3_edgeCases(t *testing.T) {
	t.Parallel()
	_, _, _, err := split.ParseImportID3("onlyone")
	if err == nil {
		t.Fatal("expected error for 1 segment")
	}
	_, _, _, err = split.ParseImportID3("a/b")
	if err == nil {
		t.Fatal("expected error for 2 segments")
	}
	_, _, _, err = split.ParseImportID3("a/b/c/d")
	if err == nil {
		t.Fatal("expected error for 4 segments")
	}
}
