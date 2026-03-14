# ai-compat.info

`ai-compat.info` is a shared compatibility dataset and reference app for AI
models, coding harnesses, and real workflow recommendations.

One source of truth in [`src/data/compat.json`](./src/data/compat.json) powers:

- The Astro site
- JSON API endpoints under `/api`
- The Go CLI `aicomp`
- The agent skill scaffold at [`skills/ai-compat/SKILL.md`](./skills/ai-compat/SKILL.md)

Live site: `https://ai-compat.info`

## Features

- Model, harness, and combo detail pages
- Full compatibility matrix across all seeded pairings
- Tier lists for models and harnesses with S/A/B/C grouping
- Best-for pages for coding, research, autonomous, local-first, and budget workflows
- Shared JSON API for site and tool consumers
- Go CLI for search, compare, combo lookup, recommendations, and tiers
- One shared dataset used everywhere

## Routes

Site pages:

- `/`
- `/matrix`
- `/models`
- `/models/[slug]`
- `/harnesses`
- `/harnesses/[slug]`
- `/combos`
- `/combos/[slug]`
- `/tiers`
- `/tiers/models`
- `/tiers/harnesses`
- `/best`
- `/best/coding`
- `/best/research`
- `/best/autonomous`
- `/best/local`
- `/best/budget`
- `/docs`
- `/about`

API endpoints:

- `/api/models.json`
- `/api/harnesses.json`
- `/api/combos.json`
- `/api/recommend.json`
- `/api/tiers/models.json`
- `/api/tiers/harnesses.json`
- `/api/best/coding.json`
- `/api/best/research.json`
- `/api/best/autonomous.json`
- `/api/best/local.json`
- `/api/best/budget.json`

## Local Development

Install dependencies and run the Astro app:

```bash
npm install
npm run dev
```

Build the site:

```bash
npm run build
```

Build the CLI:

```bash
go build -o aicomp ./cmd/aicomp
```

Run tests:

```bash
go test ./...
```

## CLI

Search:

```bash
./aicomp search claude
./aicomp search local --json
```

Compare models or harnesses:

```bash
./aicomp compare claude-opus-4 gpt-4.1
./aicomp compare claude-code aider --json
```

Inspect a specific combo:

```bash
./aicomp combo --model gpt-4.1 --harness codex-cli
./aicomp combo --model deepseek-v3 --harness continue --json
```

Get best recommendations:

```bash
./aicomp best --for coding
./aicomp best --for research --json
./aicomp best --for autonomous
./aicomp best --for local
./aicomp best --for budget
```

View tier lists:

```bash
./aicomp tiers models
./aicomp tiers models --json
./aicomp tiers harnesses
./aicomp tiers harnesses --json
```

Command summary:

- `aicomp search <query> [--json]`
- `aicomp compare <slugA> <slugB> [--json]`
- `aicomp combo --model <slug> --harness <slug> [--json]`
- `aicomp best [--for coding|research|autonomous|local|budget] [--json]`
- `aicomp tiers models [--json]`
- `aicomp tiers harnesses [--json]`

## API

Fetch models:

```bash
curl http://localhost:4321/api/models.json
```

Fetch harnesses:

```bash
curl http://localhost:4321/api/harnesses.json
```

Fetch all combos:

```bash
curl http://localhost:4321/api/combos.json
```

Fetch recommendations:

```bash
curl "http://localhost:4321/api/recommend.json?usecase=coding&limit=5"
```

Fetch model tiers:

```bash
curl http://localhost:4321/api/tiers/models.json
```

Fetch harness tiers:

```bash
curl http://localhost:4321/api/tiers/harnesses.json
```

Fetch best-for pages as JSON:

```bash
curl http://localhost:4321/api/best/coding.json
curl http://localhost:4321/api/best/research.json
curl http://localhost:4321/api/best/autonomous.json
curl http://localhost:4321/api/best/local.json
curl http://localhost:4321/api/best/budget.json
```

## Contributing Data

The source of truth is [`src/data/compat.json`](./src/data/compat.json).

When contributing:

1. Keep slugs stable and lowercase.
2. Update models, harnesses, combos, and use cases in the same file.
3. Keep scores realistic relative to the rest of the matrix.
4. Include practical descriptions, notes, and pros/cons.
5. Run `npm run build` and `go test ./...` before opening a PR.

## Deploy

The site is a standard Astro static build.

1. Install dependencies with `npm install`.
2. Run `npm run build`.
3. Deploy the `dist/` directory to any static host.
4. Build the CLI separately with `go build -o aicomp ./cmd/aicomp` if you want binaries.

## Skill Usage

The skill scaffold in [`skills/ai-compat/SKILL.md`](./skills/ai-compat/SKILL.md)
describes how an agent can use the CLI and API to:

- Search the dataset
- Compare models or harnesses
- Inspect a specific pairing
- Fetch recommendations by workflow
- Use the same compatibility knowledge from outside the website
