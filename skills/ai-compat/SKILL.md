# ai-compat

OpenClaw skill scaffold for selecting an AI model and harness pairing from the local `ai-compat` dataset.

## Purpose

Use this skill when an agent needs to:

- search available models, harnesses, or combos
- compare two models or two harnesses
- inspect one exact combo
- recommend the best stack for a use case

## Data source

Read from the repository dataset at `src/data/compat.json`.

The same data also powers:

- the Astro site
- the JSON API under `/api`
- the `aicomp` CLI

## Preferred interface

Use the local CLI first.

### Search

```bash
aicomp search claude
aicomp search multimodal --json
```

### Compare

```bash
aicomp compare claude-opus-4 claude-sonnet-4
aicomp compare codex-cli openclaw --json
```

### Combo lookup

```bash
aicomp combo --model gpt-4.1 --harness codex-cli
aicomp combo --model claude-opus-4 --harness openclaw --json
```

### Recommendations

```bash
aicomp best
aicomp best --for agent-development
aicomp best --for multimodal-analysis --json
```

## API fallback

If the site is running, these endpoints expose the same data:

- `GET /api/models.json`
- `GET /api/harnesses.json`
- `GET /api/combos.json`
- `GET /api/recommend.json?usecase=agent-development`

## Agent guidance

Recommendation heuristics:

- prefer the highest score unless the user asks for a specific provider or workflow
- use `agent-development` for autonomous tool-using systems
- use `complex-development` for hard multi-file engineering work
- use `multimodal-analysis` when screenshots, docs, or visual context matter
- use `local-development` when open tooling and portability matter

## Expected outputs

Good responses should include:

- the selected combo name
- the score
- the model and harness slugs
- the relevant use case
- one short reason grounded in the dataset notes
