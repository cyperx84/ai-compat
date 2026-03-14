import rawData from "./compat.json";

export interface Model {
  id: string;
  slug: string;
  name: string;
  provider: string;
  description: string;
  contextWindow?: string;
  released?: string;
  capabilities: string[];
  website?: string;
}

export interface Harness {
  id: string;
  slug: string;
  name: string;
  type: string;
  description: string;
  provider: string;
  status: string;
  website?: string;
  features: string[];
}

export interface Combo {
  id: string;
  slug: string;
  model: string;
  harness: string;
  name: string;
  description: string;
  score: number;
  status: string;
  notes?: string;
  usecase?: string;
  usecases?: string[];
  pros?: string[];
  cons?: string[];
}

export interface Usecase {
  id: string;
  name: string;
  description: string;
}

export interface CompatData {
  models: Model[];
  harnesses: Harness[];
  combos: Combo[];
  usecases: Usecase[];
}

export interface DecoratedCombo extends Combo {
  modelDetails?: Model;
  harnessDetails?: Harness;
  usecaseDetails?: Usecase;
  usecaseDetailsList: Usecase[];
}

export interface RankedEntry<T> {
  item: T;
  aggregateScore: number;
  comboCount: number;
  tier: TierKey;
  bestCombo?: Combo;
}

export type TierKey = "S" | "A" | "B" | "C";

export interface TierGroup<T> {
  tier: TierKey;
  label: string;
  description: string;
  items: RankedEntry<T>[];
}

export const compatData = rawData as CompatData;

const tierMeta: Record<TierKey, { label: string; description: string }> = {
  S: { label: "Exceptional", description: "Top-end options with standout aggregate results." },
  A: { label: "Great", description: "Strong choices that hold up across a wide range of combos." },
  B: { label: "Solid", description: "Dependable options with a few tradeoffs." },
  C: { label: "Usable", description: "Serviceable picks, usually for narrower or cost-driven needs." },
};

export const getModel = (slug: string) =>
  compatData.models.find((model) => model.slug === slug);

export const getHarness = (slug: string) =>
  compatData.harnesses.find((harness) => harness.slug === slug);

export const getCombo = (slug: string) =>
  compatData.combos.find((combo) => combo.slug === slug);

export const getUsecase = (id: string) =>
  compatData.usecases.find((usecase) => usecase.id === id);

export const getCombosForModel = (slug: string) =>
  compatData.combos.filter((combo) => combo.model === slug);

export const getCombosForHarness = (slug: string) =>
  compatData.combos.filter((combo) => combo.harness === slug);

export const getUsecasesForCombo = (combo: Combo) =>
  (combo.usecases?.length ? combo.usecases : combo.usecase ? [combo.usecase] : [])
    .map((id) => getUsecase(id))
    .filter(Boolean) as Usecase[];

export const decorateCombo = (combo: Combo): DecoratedCombo => ({
  ...combo,
  modelDetails: getModel(combo.model),
  harnessDetails: getHarness(combo.harness),
  usecaseDetails: combo.usecase ? getUsecase(combo.usecase) : undefined,
  usecaseDetailsList: getUsecasesForCombo(combo),
});

export const sortCombosByScore = (combos: Combo[]) =>
  [...combos].sort((left, right) => right.score - left.score);

export const getRecommendedCombos = (usecase?: string, limit = 5) =>
  getBestCombos(usecase, limit)
    .slice(0, limit)
    .map(decorateCombo);

export const getBestCombos = (usecase?: string, limit = 5) =>
  compatData.combos
    .filter((combo) => {
      if (!usecase) return true;
      return combo.usecase === usecase || combo.usecases?.includes(usecase);
    })
    .sort((left, right) => {
      const leftPrimary = left.usecase === usecase ? 1 : 0;
      const rightPrimary = right.usecase === usecase ? 1 : 0;
      if (leftPrimary !== rightPrimary) {
        return rightPrimary - leftPrimary;
      }
      return right.score - left.score;
    })
    .slice(0, limit)
    .map(decorateCombo);

const average = (values: number[]) =>
  values.reduce((sum, value) => sum + value, 0) / values.length;

export const getTierForScore = (score: number): TierKey => {
  if (score >= 9.2) return "S";
  if (score >= 8.7) return "A";
  if (score >= 8.1) return "B";
  return "C";
};

const buildRankings = <T extends Model | Harness>(
  items: T[],
  combosForItem: (slug: string) => Combo[],
): RankedEntry<T>[] =>
  items
    .map((item) => {
      const combos = sortCombosByScore(combosForItem(item.slug));
      const aggregateScore = combos.length ? average(combos.map((combo) => combo.score)) : 0;
      return {
        item,
        aggregateScore,
        comboCount: combos.length,
        tier: getTierForScore(aggregateScore),
        bestCombo: combos[0],
      };
    })
    .sort((left, right) => right.aggregateScore - left.aggregateScore);

export const getModelRankings = () =>
  buildRankings(compatData.models, getCombosForModel);

export const getHarnessRankings = () =>
  buildRankings(compatData.harnesses, getCombosForHarness);

export const getTierGroups = <T extends Model | Harness>(entries: RankedEntry<T>[]): TierGroup<T>[] =>
  (["S", "A", "B", "C"] as TierKey[]).map((tier) => ({
    tier,
    label: tierMeta[tier].label,
    description: tierMeta[tier].description,
    items: entries.filter((entry) => entry.tier === tier),
  }));
