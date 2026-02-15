# gh-deps-kit

A tool to manage GitHub Dependency graph.

## Commands

### list

```sh
gh deps-kit list [flags]
```

List dependency packages related to GitHub Actions in the repository's SBOM.

#### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--ecosystem` | `-e` | `""` | The ecosystem of the dependencies |
| `--format` | | | Output format: {json} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--name-only` | | `false` | Output only team names |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' (optional, defaults to current repository) |
| `--template` | `-t` | | Format JSON output using a Go template; see "gh help formatting" |

### submodule list

```sh
gh deps-kit submodule list [flags]
```

List submodules of the specified repository. Use --recursive to include nested submodules.

#### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | | | Output format: {json} |
| `--jq` | `-q` | | Filter JSON output using a jq expression |
| `--name-only` | | `false` | Output only submodule names |
| `--recursive` | `-r` | `false` | Recursively list nested submodules |
| `--repo` | `-R` | `""` | The repository in the format 'owner/repo' (optional, defaults to current repository) |
| `--template` | `-t` | | Format JSON output using a Go template; see "gh help formatting" |
