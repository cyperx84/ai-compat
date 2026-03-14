export type Status = 'supported' | 'partial' | 'unofficial' | 'broken' | 'deprecated';

export const STATUS_ICONS: Record<Status, string> = {
  supported: '✅',
  partial: '⚠️',
  unofficial: '🔧',
  broken: '❌',
  deprecated: '🚫',
};

export const STATUS_LABELS: Record<Status, string> = {
  supported: 'Supported',
  partial: 'Partial',
  unofficial: 'Unofficial',
  broken: 'Broken',
  deprecated: 'Deprecated',
};

export const STATUS_COLORS: Record<Status, string> = {
  supported: '#22c55e',
  partial: '#eab308',
  unofficial: '#3b82f6',
  broken: '#ef4444',
  deprecated: '#6b7280',
};

export interface Evidence {
  label: string;
  url: string;
}

export interface CompatibilityEntry {
  tool: string;       // tool slug
  model: string;      // model slug
  status: Status;
  evidence?: Evidence[];
  notes?: string;
  since?: string;     // version/date when support was added
}

export interface ToolModelCompat {
  model: string;
  status: Status;
  evidence?: Evidence[];
  notes?: string;
  since?: string;
}

export interface ModelFeatureCompat {
  model: string;
  status: Status;
  notes?: string;
  evidence?: Evidence[];
}

export interface ToolMcpCompat {
  tool: string;
  status: Status;
  notes?: string;
  since?: string;
  evidence?: Evidence[];
}

export interface RouteEntry {
  provider: string;
  router: string;
  model?: string;
  status: Status;
  notes?: string;
  evidence?: Evidence[];
}

export interface Model {
  id: string;
  name: string;
  provider: string;
  providerSlug: string;
  description: string;
  contextWindow?: string;
  released?: string;
  capabilities: string[];
  website?: string;
}

export interface Tool {
  id: string;
  name: string;
  type: 'editor' | 'cli' | 'ide' | 'extension' | 'platform' | 'agent-runtime';
  description: string;
  website?: string;
  mcpSupport: Status;
  mcpNotes?: string;
  models: ToolModelCompat[];
}

export interface Feature {
  id: string;
  name: string;
  description: string;
  category: string;
  models: ModelFeatureCompat[];
}

export interface ProviderRoute {
  id: string;
  provider: string;
  description: string;
  website?: string;
  routes: RouteEntry[];
}
