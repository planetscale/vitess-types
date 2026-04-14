# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`vitess-types` extracts Protocol Buffer definitions from [Vitess](https://github.com/vitessio/vitess) into a lightweight, standalone Go module. The goal is to allow third-party projects to use Vitess types (e.g., `QueryResult`, `Row`) via gRPC without pulling in the entire Vitess codebase as a dependency.

The module is published as a Buf module at `buf.build/planetscale/vitess`.

## Common Commands

```bash
# Download proto files from Vitess GitHub and regenerate all code
make download

# Regenerate Go code from existing proto files
make proto

# Format Go and YAML files
make fmt

# Install required tools into bin/
make tools

# Update Go dependencies and regenerate
make update

# Remove generated code (gen/) and binaries (bin/)
make clean
```

There are no unit tests — this is a pure code generation project. Validation happens through CI (license checks and Buf module publishing).

## Architecture

### Versioned Proto Structure

Proto source files live under `src/vitess/{package}/{version}/{package}.proto`, and generated Go code mirrors that structure under `gen/vitess/{package}/{version}/`. Each package (e.g., `query`, `binlogdata`, `topodata`) supports multiple Vitess versions simultaneously:

- `dev` — Vitess main branch
- `v21` — Vitess v21.0.5
- `v22` — Vitess v22.0.4
- `v23` — Vitess v23.0.3
- `v24` — Vitess v24.0.0-rc1

Version pinning is controlled by `manifest.json`:
```json
{
  "versions": {
    "dev": "refs/heads/main",
    "v21": "refs/tags/v21.0.5",
    "v22": "refs/tags/v22.0.4",
    "v23": "refs/tags/v23.0.3",
    "v24": "refs/tags/v24.0.0-rc1"
  }
}
```

### Code Generation Pipeline

`make download` runs `scripts/download.go`, which:
1. Fetches proto files from Vitess GitHub for each configured version
2. Rewrites proto `import` paths to match the versioned directory structure
3. Regenerates `buf.gen.yaml` from `buf.gen.yaml.tmpl` with version-specific object pooling

`make proto` then runs `buf generate` using three generators:
1. `protoc-gen-go` — standard protobuf (`.pb.go`)
2. `protoc-gen-go-vtproto` — Vitess-optimized protobuf with object pooling for `Row`, `BoundQuery`, `VStreamRowsResponse` (`_vtproto.pb.go`)
3. `protoc-gen-connect-go` — Connect RPC service stubs (`.connect.go`)

After `buf generate`, `scripts/fix-service-names.go` post-processes generated service code to fix naming conventions.

### Key Files

| File | Purpose |
|------|---------|
| `manifest.json` | Vitess version → git ref mapping |
| `buf.gen.yaml.tmpl` | Template for code generator config (rendered per `make download`) |
| `src/buf.yaml` | Buf module definition (`buf.build/planetscale/vitess`) |
| `scripts/download.go` | Downloads protos from Vitess GitHub, rewrites imports |
| `scripts/fix-service-names.go` | Post-processes generated Connect service names |

### CI

- **`publish.yml`** — Publishes to Buf Schema Registry on push to main or PR. Uses `BUF_TOKEN` and `BUF_USER` secrets.
- **`licensing.yml`** — Runs `license_finder` to verify all dependencies have approved licenses.

## Adding a New Vitess Version

1. Add the version and its git ref to `manifest.json`
2. Run `make download` to fetch protos and regenerate code
3. Verify generated files look correct under `gen/vitess/`
