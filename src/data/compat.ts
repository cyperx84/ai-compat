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

export const compatData = rawData as CompatData;

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

export const getRecommendedCombos = (usecase?: string, limit = 5) =>
  compatData.combos
    .filter((combo) => (usecase ? combo.usecase === usecase : true))
    .sort((left, right) => right.score - left.score)
    .slice(0, limit)
    .map((combo) => ({
      ...combo,
      modelDetails: getModel(combo.model),
      harnessDetails: getHarness(combo.harness),
      usecaseDetails: combo.usecase ? getUsecase(combo.usecase) : undefined,
    }));
