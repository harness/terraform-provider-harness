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

	var decoded map[string]map[string]map[string]interface{}
	if err := json.Unmarshal([]byte(got), &decoded); err != nil {
		t.Fatalf("outputs not valid JSON: %v", err)
	}

	if _, ok := decoded["no_outcomes"]; ok {
		t.Fatalf("expected node without outcomes to be skipped, got %v", decoded)
	}

	greeting, ok := decoded["ShellScript_1"]["output"]["greeting"]
	if !ok || greeting != "hello world" {
		t.Fatalf("expected ShellScript_1 outcome greeting=hello world, got %v", decoded)
	}
}
