package data

import "testing"

func TestLoadData(t *testing.T) {
	compat, err := LoadData()
	if err != nil {
		t.Fatalf("LoadData failed: %v", err)
	}

	if len(compat.Models) != 14 {
		t.Fatalf("expected 14 models, got %d", len(compat.Models))
	}

	if len(compat.Harnesses) != 11 {
		t.Fatalf("expected 11 harnesses, got %d", len(compat.Harnesses))
	}

	if len(compat.Combos) < 100 {
		t.Fatalf("expected seeded combos, got %d", len(compat.Combos))
	}

	if len(compat.Usecases) == 0 {
		t.Fatal("expected usecases to be loaded")
	}
}

func TestFinders(t *testing.T) {
	compat, err := LoadData()
	if err != nil {
		t.Fatalf("LoadData failed: %v", err)
	}

	if compat.FindModel("claude-opus-4") == nil {
		t.Fatal("expected claude-opus-4 model")
	}

	if compat.FindHarness("openclaw") == nil {
		t.Fatal("expected openclaw harness")
	}

	combo := compat.FindCombo("claude-opus-4-claude-code")
	if combo == nil || combo.Score < 9.5 {
		t.Fatalf("unexpected combo payload: %#v", combo)
	}

	if compat.FindComboByParts("gpt-4.1", "codex-cli") == nil {
		t.Fatal("expected combo by parts")
	}

	if compat.FindUsecase("autonomous") == nil {
		t.Fatal("expected autonomous usecase")
	}
}
