package trigger_pipeline

import (
	"encoding/json"
	"testing"

	pipeline_go_sdk "github.com/harness/harness-go-sdk/harness/nextgen"
)

func TestFlattenOutcomesNilGraph(t *testing.T) {
	got, err := flattenOutcomes(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "{}" {
		t.Fatalf("expected empty object, got %q", got)
	}
}

func TestFlattenOutcomesSkipsNodesWithoutOutcomes(t *testing.T) {
	graph := &pipeline_go_sdk.ExecutionGraph{
		NodeMap: map[string]pipeline_go_sdk.ExecutionNode{
			"n1": {Identifier: "no_outcomes"},
			"n2": {
				Identifier: "ShellScript_1",
				Outcomes: map[string]map[string]interface{}{
					"output": {"greeting": "hello world"},
				},
			},
		},
	}

	got, err := flattenOutcomes(graph)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var decoded map[string][]map[string]map[string]interface{}
	if err := json.Unmarshal([]byte(got), &decoded); err != nil {
		t.Fatalf("outputs not valid JSON: %v", err)
	}

	if _, ok := decoded["no_outcomes"]; ok {
		t.Fatalf("expected node without outcomes to be skipped, got %v", decoded)
	}

	entries := decoded["ShellScript_1"]
	if len(entries) != 1 || entries[0]["output"]["greeting"] != "hello world" {
		t.Fatalf("expected ShellScript_1 outcome greeting=hello world, got %v", decoded)
	}
}

// NodeMap is keyed by node execution id, not by step identifier. Matrix/parallel looping
// strategies run the same identifier as multiple nodes, so both sets of outcomes must be
// preserved instead of one silently overwriting the other.
func TestFlattenOutcomesGroupsMatrixNodesByIdentifier(t *testing.T) {
	graph := &pipeline_go_sdk.ExecutionGraph{
		NodeMap: map[string]pipeline_go_sdk.ExecutionNode{
			"node-b": {
				Identifier: "ShellScript_1",
				Outcomes: map[string]map[string]interface{}{
					"output": {"greeting": "second"},
				},
			},
			"node-a": {
				Identifier: "ShellScript_1",
				Outcomes: map[string]map[string]interface{}{
					"output": {"greeting": "first"},
				},
			},
		},
	}

	got, err := flattenOutcomes(graph)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var decoded map[string][]map[string]map[string]interface{}
	if err := json.Unmarshal([]byte(got), &decoded); err != nil {
		t.Fatalf("outputs not valid JSON: %v", err)
	}

	entries := decoded["ShellScript_1"]
	if len(entries) != 2 {
		t.Fatalf("expected both matrix node outcomes to be preserved, got %v", decoded)
	}
	// Node map keys are sorted (node-a before node-b) for stable, reproducible output.
	if entries[0]["output"]["greeting"] != "first" || entries[1]["output"]["greeting"] != "second" {
		t.Fatalf("expected outcomes ordered by node id, got %v", decoded)
	}
}
