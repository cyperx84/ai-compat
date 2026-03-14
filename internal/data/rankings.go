package data

import "sort"

type Tier string

const (
	TierS Tier = "S"
	TierA Tier = "A"
	TierB Tier = "B"
	TierC Tier = "C"
)

type ModelRanking struct {
	Model          *Model  `json:"model"`
	AggregateScore float64 `json:"aggregateScore"`
	ComboCount     int     `json:"comboCount"`
	Tier           Tier    `json:"tier"`
	BestCombo      *Combo  `json:"bestCombo,omitempty"`
}

type HarnessRanking struct {
	Harness        *Harness `json:"harness"`
	AggregateScore float64  `json:"aggregateScore"`
	ComboCount     int      `json:"comboCount"`
	Tier           Tier     `json:"tier"`
	BestCombo      *Combo   `json:"bestCombo,omitempty"`
}

func GetTierForScore(score float64) Tier {
	switch {
	case score >= 9.2:
		return TierS
	case score >= 8.7:
		return TierA
	case score >= 8.1:
		return TierB
	default:
		return TierC
	}
}

func (c *CompatData) CombosForModel(slug string) []Combo {
	var combos []Combo
	for _, combo := range c.Combos {
		if combo.Model == slug {
			combos = append(combos, combo)
		}
	}
	sort.Slice(combos, func(i, j int) bool {
		return combos[i].Score > combos[j].Score
	})
	return combos
}

func (c *CompatData) CombosForHarness(slug string) []Combo {
	var combos []Combo
	for _, combo := range c.Combos {
		if combo.Harness == slug {
			combos = append(combos, combo)
		}
	}
	sort.Slice(combos, func(i, j int) bool {
		return combos[i].Score > combos[j].Score
	})
	return combos
}

func (c *CompatData) ComboMatchesUsecase(combo Combo, usecase string) bool {
	if usecase == "" || combo.Usecase == usecase {
		return true
	}
	for _, candidate := range combo.Usecases {
		if candidate == usecase {
			return true
		}
	}
	return false
}

func (c *CompatData) BestCombos(usecase string, limit int) []Combo {
	var matches []Combo
	for _, combo := range c.Combos {
		if c.ComboMatchesUsecase(combo, usecase) {
			matches = append(matches, combo)
		}
	}

	sort.Slice(matches, func(i, j int) bool {
		leftPrimary := 0
		rightPrimary := 0
		if matches[i].Usecase == usecase {
			leftPrimary = 1
		}
		if matches[j].Usecase == usecase {
			rightPrimary = 1
		}
		if leftPrimary != rightPrimary {
			return leftPrimary > rightPrimary
		}
		return matches[i].Score > matches[j].Score
	})

	if limit > 0 && len(matches) > limit {
		matches = matches[:limit]
	}

	return matches
}

func averageScore(combos []Combo) float64 {
	if len(combos) == 0 {
		return 0
	}
	total := 0.0
	for _, combo := range combos {
		total += combo.Score
	}
	return total / float64(len(combos))
}

func (c *CompatData) ModelRankings() []ModelRanking {
	rankings := make([]ModelRanking, 0, len(c.Models))
	for i := range c.Models {
		model := &c.Models[i]
		combos := c.CombosForModel(model.Slug)
		ranking := ModelRanking{
			Model:          model,
			AggregateScore: averageScore(combos),
			ComboCount:     len(combos),
			Tier:           GetTierForScore(averageScore(combos)),
		}
		if len(combos) > 0 {
			ranking.BestCombo = &combos[0]
		}
		rankings = append(rankings, ranking)
	}

	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].AggregateScore > rankings[j].AggregateScore
	})

	return rankings
}

func (c *CompatData) HarnessRankings() []HarnessRanking {
	rankings := make([]HarnessRanking, 0, len(c.Harnesses))
	for i := range c.Harnesses {
		harness := &c.Harnesses[i]
		combos := c.CombosForHarness(harness.Slug)
		ranking := HarnessRanking{
			Harness:        harness,
			AggregateScore: averageScore(combos),
			ComboCount:     len(combos),
			Tier:           GetTierForScore(averageScore(combos)),
		}
		if len(combos) > 0 {
			ranking.BestCombo = &combos[0]
		}
		rankings = append(rankings, ranking)
	}

	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].AggregateScore > rankings[j].AggregateScore
	})

	return rankings
}
