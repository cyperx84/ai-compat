package cli

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func executeRoot(t *testing.T, args ...string) (string, error) {
	t.Helper()

	cmd := NewRootCommand()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs(args)

	err := cmd.Execute()
	return out.String(), err
}

func TestSearchJSON(t *testing.T) {
	out, err := executeRoot(t, "search", "claude", "--json")
	if err != nil {
		t.Fatalf("search failed: %v", err)
	}

	var results []map[string]any
	if err := json.Unmarshal([]byte(out), &results); err != nil {
		t.Fatalf("invalid json output: %v", err)
	}

	if len(results) == 0 {
		t.Fatal("expected non-empty search results")
	}
}

func TestCompareModelsText(t *testing.T) {
	out, err := executeRoot(t, "compare", "claude-opus-4", "claude-sonnet-4")
	if err != nil {
		t.Fatalf("compare failed: %v", err)
	}

	if !strings.Contains(out, "Model comparison") || !strings.Contains(out, "Claude Opus 4") {
		t.Fatalf("unexpected compare output: %s", out)
	}
}

func TestComboJSON(t *testing.T) {
	out, err := executeRoot(t, "combo", "--model", "gpt-4.1", "--harness", "codex-cli", "--json")
	if err != nil {
		t.Fatalf("combo failed: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal([]byte(out), &payload); err != nil {
		t.Fatalf("invalid json output: %v", err)
	}

	if payload["combo"] == nil {
		t.Fatalf("expected combo payload, got: %v", payload)
	}
}

func TestBestForUsecase(t *testing.T) {
	out, err := executeRoot(t, "best", "--for", "openclaw")
	if err != nil {
		t.Fatalf("best failed: %v", err)
	}

	if !strings.Contains(out, "Top combos for openclaw") {
		t.Fatalf("unexpected best output: %s", out)
	}
}

func TestTiersModelsJSON(t *testing.T) {
	out, err := executeRoot(t, "tiers", "models", "--json")
	if err != nil {
		t.Fatalf("tiers models failed: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal([]byte(out), &payload); err != nil {
		t.Fatalf("invalid json output: %v", err)
	}

	if payload["kind"] != "models" {
		t.Fatalf("unexpected payload: %v", payload)
	}
}
