package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Model struct {
	ID            string   `json:"id"`
	Slug          string   `json:"slug"`
	Name          string   `json:"name"`
	Provider      string   `json:"provider"`
	Description   string   `json:"description"`
	ContextWindow string   `json:"contextWindow,omitempty"`
	Released      string   `json:"released,omitempty"`
	Capabilities  []string `json:"capabilities"`
	Website       string   `json:"website,omitempty"`
}

type Harness struct {
	ID          string   `json:"id"`
	Slug        string   `json:"slug"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Provider    string   `json:"provider"`
	Status      string   `json:"status"`
	Website     string   `json:"website,omitempty"`
	Features    []string `json:"features"`
}

type Combo struct {
	ID          string   `json:"id"`
	Slug        string   `json:"slug"`
	Model       string   `json:"model"`
	Harness     string   `json:"harness"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Score       float64  `json:"score"`
	Status      string   `json:"status"`
	Notes       string   `json:"notes,omitempty"`
	Usecase     string   `json:"usecase,omitempty"`
	Usecases    []string `json:"usecases,omitempty"`
	Pros        []string `json:"pros,omitempty"`
	Cons        []string `json:"cons,omitempty"`
}

type Usecase struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CompatData struct {
	Models    []Model   `json:"models"`
	Harnesses []Harness `json:"harnesses"`
	Combos    []Combo   `json:"combos"`
	Usecases  []Usecase `json:"usecases"`
}

func LoadData() (*CompatData, error) {
	path, err := findDataFile()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var compat CompatData
	if err := json.Unmarshal(content, &compat); err != nil {
		return nil, err
	}

	return &compat, nil
}

func findDataFile() (string, error) {
	candidates := []string{
		filepath.Join("src", "data", "compat.json"),
		filepath.Join("..", "src", "data", "compat.json"),
		filepath.Join("..", "..", "src", "data", "compat.json"),
	}

	if _, file, _, ok := runtime.Caller(0); ok {
		base := filepath.Dir(file)
		candidates = append(candidates, filepath.Join(base, "..", "..", "src", "data", "compat.json"))
	}

	if exePath, err := os.Executable(); err == nil {
		base := filepath.Dir(exePath)
		candidates = append(candidates,
			filepath.Join(base, "src", "data", "compat.json"),
			filepath.Join(base, "..", "src", "data", "compat.json"),
			filepath.Join(base, "..", "..", "src", "data", "compat.json"),
			filepath.Join(base, "..", "..", "..", "src", "data", "compat.json"),
		)
	}

	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("compat dataset not found")
}

func (c *CompatData) FindModel(slug string) *Model {
	for i := range c.Models {
		if c.Models[i].Slug == slug {
			return &c.Models[i]
		}
	}
	return nil
}

func (c *CompatData) FindHarness(slug string) *Harness {
	for i := range c.Harnesses {
		if c.Harnesses[i].Slug == slug {
			return &c.Harnesses[i]
		}
	}
	return nil
}

func (c *CompatData) FindCombo(slug string) *Combo {
	for i := range c.Combos {
		if c.Combos[i].Slug == slug {
			return &c.Combos[i]
		}
	}
	return nil
}

func (c *CompatData) FindUsecase(id string) *Usecase {
	for i := range c.Usecases {
		if c.Usecases[i].ID == id {
			return &c.Usecases[i]
		}
	}
	return nil
}

func (c *CompatData) FindComboByParts(modelSlug, harnessSlug string) *Combo {
	for i := range c.Combos {
		if c.Combos[i].Model == modelSlug && c.Combos[i].Harness == harnessSlug {
			return &c.Combos[i]
		}
	}
	return nil
}
