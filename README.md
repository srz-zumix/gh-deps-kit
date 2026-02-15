# gh-deps-kit

A tool to manage GitHub Dependency graph.

## Commands

### List dependency packages

```sh
gh deps-kit list [flags]
```

List dependency packages in the repository's SBOM.

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--ecosystem` | `-e` | `""` | The ecosystem of the dependencies |
| `--format` | | | Output format: {json} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--name-only` | | `false` | Output only team names |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--template` | `-t` | | Format JSON output using a Go template; see "gh help formatting" |

### Actions

### Graph actions dependency

```sh
gh deps-kit actions graph [flags]
```

Output dependency relationships of GitHub Actions as a Mermaid flowchart. Use --recursive to traverse referenced action repositories.

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | | `"mermaid"` | Output format: {json\|mermaid\|markdown} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--output` | `-o` | | Output file path (default: stdout) |
| `--recursive` | `-r` | `false` | Recursively traverse referenced action repositories |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--template` | `-t` | | Format JSON output using a Go template; see "gh help formatting" |

### List actions dependency packages

```sh
gh deps-kit actions list [flags]
```

List dependency packages related to GitHub Actions in the repository's SBOM. Use --recursive to traverse referenced action repositories.

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | | | Output format: {json} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--name-only` | | `false` | Output only team names |
| `--recursive` | `-r` | `false` | Recursively traverse referenced action repositories |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--template` | `-t` | | Format JSON output using a Go template; see "gh help formatting" |

### Submodule

### List repository submodules

```sh
gh deps-kit submodule list [flags]
```

List submodules of the specified repository. Use --recursive to include nested submodules.

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | | | Output format: {json} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--name-only` | | `false` | Output only submodule names |
| `--recursive` | `-r` | `false` | Recursively list nested submodules |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' |
| `--template` | `-t` | | Format JSON output using a Go template; see "gh help formatting" |
