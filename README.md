# ai-compat.info

`ai-compat.info` is a small compatibility matrix for AI models and coding harnesses.
This repo ships one shared dataset that powers:

- An Astro site
- JSON API endpoints under `/api`
- A Go CLI named `aicomp`
- An OpenClaw skill scaffold at `skills/ai-compat/SKILL.md`

## MVP surface

Site routes:

- `/`
- `/matrix`
- `/models/[slug]`
- `/harnesses/[slug]`
- `/combos/[slug]`

JSON endpoints:

- `/api/models.json`
- `/api/harnesses.json`
- `/api/combos.json`
- `/api/recommend.json`

## Shared data

The source of truth is [`src/data/compat.json`](./src/data/compat.json).
Astro imports it through [`src/data/compat.ts`](./src/data/compat.ts), and the Go CLI loads the same JSON file from disk.

Seed data includes:

- 4 models
- 5 harnesses
- 7 combos
- 6 use cases

## Local usage

Install dependencies and run the site:

```bash
npm install
npm run dev
```

Build the Astro site:

```bash
npm run build
```

Build the CLI:

```bash
go build -o aicomp ./cmd/aicomp
```

Examples:

```bash
./aicomp search claude
./aicomp search agent --json
./aicomp compare claude-opus-4 claude-sonnet-4
./aicomp compare codex-cli openclaw --json
./aicomp combo --model gpt-4.1 --harness codex-cli
./aicomp best --for agent-development
./aicomp best --for multimodal-analysis --json
```

Run tests:

```bash
go test ./...
```

## CLI commands

`aicomp search <query>`

- Searches models, harnesses, and combos.
- Supports `--json`.

`aicomp compare <slugA> <slugB>`

- Compares two models or two harnesses.
- Supports `--json`.

`aicomp combo --model <slug> --harness <slug>`

- Returns one specific pairing.
- Supports `--json`.

`aicomp best --for <usecase>`

- Returns top combos overall or for a specific use case.
- Supports `--json`.

## OpenClaw skill

The scaffold at [`skills/ai-compat/SKILL.md`](./skills/ai-compat/SKILL.md) describes how an agent can use the CLI or JSON endpoints to search, compare, and recommend stacks.
