---
name: gh-deps-kit
description: gh-deps-kit is a GitHub CLI extension for inspecting dependency graphs. Use it to list SBOM packages, analyze GitHub Actions workflow dependencies (with graph/lint/filter options), list Git submodules, and inspect Unity package manifests — all directly from the command line via `gh deps-kit`.
---

# gh-deps-kit

A GitHub CLI extension (`gh deps-kit`) to manage and inspect GitHub dependency graphs.
It supports SBOM-based package listing, GitHub Actions dependency analysis (graph/lint/filter), Git submodule listing, and Unity package inspection.

## Prerequisites

### Installation

```sh
gh extension install srz-zumix/gh-deps-kit
```

### Authentication

`gh deps-kit` uses the `gh` CLI's authentication. Ensure you are authenticated before using the extension:

```sh
gh auth login
gh auth status
```

## CLI Structure

```
gh deps-kit                        # Root command
├── list                           # List dependency packages (SBOM)
├── actions                        # GitHub Actions subcommands
│   ├── graph                      # Graph Actions dependencies
│   ├── lint                       # Lint workflow/action YAML files
│   ├── list                       # List Actions packages from SBOM
│   └── workflow                   # List action deps from workflow YAML
├── submodule                      # Git submodule subcommands
│   └── list                       # List repository submodules
└── unity                          # Unity project subcommands
    └── list                       # List Unity package dependencies
```

## List dependency packages (gh deps-kit list)

```sh
gh deps-kit list [flags]
```

List dependency packages in the repository's SBOM.

```sh
# List all packages in the current repository
gh deps-kit list

# List packages for a specific repository
gh deps-kit list --repo owner/repo

# Filter by ecosystem
gh deps-kit list --include npm
gh deps-kit list --include npm --include pip

# Exclude an ecosystem
gh deps-kit list --exclude rubygems

# Output package names only
gh deps-kit list --name-only

# JSON output
gh deps-kit list --format json

# Filter JSON with jq
gh deps-kit list --format json --jq '.[].name'
```

**Flags:**

| Flag | Short | Default | Description |
| ------ | ------- | --------- | ------------- |
| `--exclude` | `-e` | | Exclude packages by ecosystem (repeatable) |
| `--format` | | | Output format: {json} |
| `--include` | `-i` | | Filter by ecosystem (repeatable) |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--name-only` | | `false` | Output only package names |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--template` | `-t` | | Format JSON output using a Go template |

## Actions

### Graph Actions dependencies (gh deps-kit actions graph)

```sh
gh deps-kit actions graph [flags]
```

Output dependency relationships of GitHub Actions as a graph. Use `--recursive` to traverse referenced action repositories.

```sh
# Output Mermaid flowchart (default)
gh deps-kit actions graph

# Output as DOT format
gh deps-kit actions graph --format dot

# Output as draw.io XML
gh deps-kit actions graph --format drawio --output deps.drawio

# Output as Markdown
gh deps-kit actions graph --format markdown

# Recursively include referenced action repositories
gh deps-kit actions graph --recursive

# For a specific repository
gh deps-kit actions graph --repo owner/repo --recursive
```

**Flags:**

| Flag | Short | Default | Description |
| ------ | ------- | --------- | ------------- |
| `--format` | | `"mermaid"` | Output format: {json\|dot\|drawio\|mermaid\|markdown} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--output` | `-o` | | Output file path (default: stdout) |
| `--recursive` | `-r` | `false` | Recursively traverse referenced action repositories |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--template` | `-t` | | Format JSON output using a Go template |

### Lint workflow and action YAML files (gh deps-kit actions lint)

```sh
gh deps-kit actions lint [<workflow-id> | <workflow-name> | <filename>] [flags] [-- <tool-args>...]
```

Run an external lint tool (actionlint or zizmor) against workflow YAML and action.yml files fetched via the GitHub API.
Optionally specify a workflow by its ID, name, or filename to lint only that workflow's dependencies.
Extra arguments after `--` are passed directly to the lint tool.

```sh
# Lint all workflows in the current repository (default tool: zizmor)
gh deps-kit actions lint

# Use actionlint instead
gh deps-kit actions lint --tool actionlint

# Lint only a specific workflow
gh deps-kit actions lint ci.yml

# Recursively lint referenced action repositories
gh deps-kit actions lint --recursive

# Pass extra args to the lint tool
gh deps-kit actions lint -- --no-color

# Lint from a specific branch
gh deps-kit actions lint --ref main

# Keep downloaded files in a specific directory
gh deps-kit actions lint --tmpdir /tmp/lint-work
```

**Flags:**

| Flag | Short | Default | Description |
| ------ | ------- | --------- | ------------- |
| `--recursive` | `-r` | `false` | Recursively traverse referenced action repositories |
| `--ref` | | `""` | Git reference (branch, tag, or commit SHA) to read workflow files from |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--tmpdir` | | `""` | Directory to store downloaded files (default: auto-created temp dir, removed after lint) |
| `--tool` | | `"zizmor"` | Lint tool to use (supported: actionlint, zizmor) |

### List action dependencies from workflow YAML files (gh deps-kit actions workflow)

```sh
gh deps-kit actions workflow [<workflow-id> | <workflow-name> | <filename>] [flags]
```

Parse workflow YAML (`.github/workflows/*.yml`) and `action.yml` files to list GitHub Actions dependencies.
Unlike `gh deps-kit list`, this command parses YAML files directly without the Dependency Graph API.
`--min-node-version` and `--filter-using` automatically enable `--recursive` to populate `runs.using` fields.

```sh
# List all action dependencies in the current repository
gh deps-kit actions workflow

# List dependencies for a specific workflow
gh deps-kit actions workflow ci.yml

# Recursively traverse referenced action repositories
gh deps-kit actions workflow --recursive

# Output as Mermaid graph (which annotates nodes with runs.using)
gh deps-kit actions workflow --recursive --format mermaid

# Output as DOT graph with tooltip annotations
gh deps-kit actions workflow --recursive --format dot

# Output as ASCII tree (grouped by job)
gh deps-kit actions workflow --recursive --format tree

# Output as draw.io XML with tooltip annotations
gh deps-kit actions workflow --recursive --format drawio --output deps.drawio

# Show only action names
gh deps-kit actions workflow --name-only

# Show action names with version ref
gh deps-kit actions workflow --name-with-ref

# Show specific fields in table output
gh deps-kit actions workflow --field Name,Version,Using,Job

# Filter: show only workflows/actions that use a Node action older than node24
gh deps-kit actions workflow --min-node-version 24

# Filter: show only actions using composite runtime
gh deps-kit actions workflow --filter-using composite

# Filter: show all node actions (prefix match: node matches node16, node20, etc.)
gh deps-kit actions workflow --filter-using node

# Filter: multiple using types
gh deps-kit actions workflow --filter-using composite --filter-using docker

# Read from a specific branch/tag/SHA
gh deps-kit actions workflow --ref v1.2.3
```

**Flags:**

| Flag | Short | Default | Description |
| ------ | ------- | --------- | ------------- |
| `--field` | | | Comma-separated list of fields. Available: Name, Version, Owner, Repo, Path, Raw, Using, Node_Version, Job |
| `--filter-using` | | | Filter by `runs.using` type (e.g. `node16`, `composite`, `docker`); prefix match; repeatable; auto-enables `--recursive` |
| `--format` | | | Output format: {json\|dot\|drawio\|mermaid\|markdown\|tree} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--min-node-version` | | `0` | Filter to show only actions using a Node version older than specified (e.g. `24` → node20, node16); auto-enables `--recursive` |
| `--name-only` | | `false` | Output only action names |
| `--name-with-ref` | | `false` | Output action names with version ref (e.g. `actions/checkout@v4`) |
| `--recursive` | `-r` | `false` | Recursively traverse referenced action repositories |
| `--ref` | | `""` | Git reference (branch, tag, or commit SHA) to read workflow files from |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--template` | `-t` | | Format JSON output using a Go template |

### List actions dependency packages (gh deps-kit actions list)

```sh
gh deps-kit actions list [flags]
```

List dependency packages related to GitHub Actions in the repository's SBOM. Use `--recursive` to traverse referenced action repositories.

```sh
# List Actions packages in the current repository
gh deps-kit actions list

# Recursively include referenced action repositories
gh deps-kit actions list --recursive

# JSON output
gh deps-kit actions list --format json

# Output names only
gh deps-kit actions list --name-only
```

**Flags:**

| Flag | Short | Default | Description |
| ------ | ------- | --------- | ------------- |
| `--format` | | | Output format: {json} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--name-only` | | `false` | Output only package names |
| `--recursive` | `-r` | `false` | Recursively traverse referenced action repositories |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--template` | `-t` | | Format JSON output using a Go template |

## Submodule

### List repository submodules (gh deps-kit submodule list)

```sh
gh deps-kit submodule list [flags]
```

List submodules of the specified repository. Use `--recursive` to include nested submodules.

```sh
# List submodules in the current repository
gh deps-kit submodule list

# Include nested submodules
gh deps-kit submodule list --recursive

# JSON output
gh deps-kit submodule list --format json

# Output submodule names only
gh deps-kit submodule list --name-only
```

**Flags:**

| Flag | Short | Default | Description |
| ------ | ------- | --------- | ------------- |
| `--format` | | | Output format: {json} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--name-only` | | `false` | Output only submodule names |
| `--recursive` | `-r` | `false` | Recursively list nested submodules |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--template` | `-t` | | Format JSON output using a Go template |

## Unity

### List Unity package dependencies (gh deps-kit unity list)

```sh
gh deps-kit unity list [flags]
```

List dependency packages defined in a Unity project's `Packages/manifest.json`. The manifest path defaults to `Packages/manifest.json` and can be overridden with `--path`.

```sh
# List Unity packages in the current repository
gh deps-kit unity list

# Use a custom manifest path
gh deps-kit unity list --path Assets/Packages/manifest.json

# Read from a specific branch or tag
gh deps-kit unity list --ref release/1.0

# Show specific fields
gh deps-kit unity list --field Name,Version,URL

# Output package names only
gh deps-kit unity list --name-only

# JSON output
gh deps-kit unity list --format json
```

**Flags:**

| Flag | Short | Default | Description |
| ------ | ------- | --------- | ------------- |
| `--field` | | `"Name,Version,SHA,Path,URL"` | Comma-separated list of fields. Available: Name, Version, SHA, Path, URL |
| `--format` | | | Output format: {json} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--name-only` | | `false` | Output only package names |
| `--path` | | `"Packages/manifest.json"` | Path to manifest.json within the repository |
| `--ref` | | `""` | Branch, tag, or commit SHA to read from (default: repository default branch) |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--template` | `-t` | | Format JSON output using a Go template |

## Common Workflows

### Find workflows using outdated Node actions

```sh
# Find all workflows/actions still using node16 or node20 (older than node24)
gh deps-kit actions workflow --min-node-version 24

# Output as ASCII tree showing which job uses each old action
gh deps-kit actions workflow --min-node-version 24 --format tree

# Get only the action names for batch update reference
gh deps-kit actions workflow --min-node-version 24 --name-with-ref
```

### Visualize the full Actions dependency graph

```sh
# Generate a Mermaid diagram showing all action dependencies
gh deps-kit actions workflow --recursive --format mermaid

# Save draw.io diagram for sharing
gh deps-kit actions workflow --recursive --format drawio --output actions-deps.drawio

# DOT format for Graphviz rendering
gh deps-kit actions workflow --recursive --format dot | dot -Tpng -o deps.png
```

### Filter actions by runtime type

```sh
# Show only composite actions and their callers
gh deps-kit actions workflow --filter-using composite

# Show only Docker-based actions
gh deps-kit actions workflow --filter-using docker

# Show all node-based actions (prefix match)
gh deps-kit actions workflow --filter-using node

# Combine multiple filters
gh deps-kit actions workflow --filter-using composite --filter-using docker
```

### Audit dependencies for a specific branch or release

```sh
# Inspect dependencies on a release tag
gh deps-kit actions workflow --ref v2.3.0 --recursive

# Inspect SBOM packages at a specific commit
gh deps-kit list --repo owner/repo

# Lint workflows from a feature branch
gh deps-kit actions lint --ref feature/new-workflow
```

## Getting Help

```sh
# General help
gh deps-kit --help

# Subcommand help
gh deps-kit actions --help
gh deps-kit actions workflow --help
gh deps-kit actions lint --help

# Shell completion setup
gh deps-kit completion --help
```

## References

- Repository: https://github.com/srz-zumix/gh-deps-kit
- Shell Completion Guide: https://github.com/srz-zumix/go-gh-extension/blob/main/docs/shell-completion.md
