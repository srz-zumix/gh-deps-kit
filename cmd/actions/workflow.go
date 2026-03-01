package actions

import (
	"context"
	"fmt"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/srz-zumix/go-gh-extension/pkg/cmdflags"
	"github.com/srz-zumix/go-gh-extension/pkg/gh"
	"github.com/srz-zumix/go-gh-extension/pkg/parser"
	"github.com/srz-zumix/go-gh-extension/pkg/render"
)

type WorkflowOptions struct {
	Exporter cmdutil.Exporter
}

// NewWorkflowCmd returns the actions workflow command
func NewWorkflowCmd() *cobra.Command {
	var repo string
	var ref string
	var nameOnly bool
	var recursive bool
	var format string
	var fields []string
	opts := &WorkflowOptions{}

	cmd := &cobra.Command{
		Use:   "workflow [<workflow-id> | <workflow-name> | <filename>]",
		Short: "List action dependencies from workflow YAML files",
		Long:  "Parse workflow YAML (.github/workflows/*.yml) and action.yml files in the repository to list GitHub Actions dependencies. Unlike the 'list' command which uses the Dependency Graph API, this command directly parses YAML files. Optionally specify a workflow by its ID, name, or filename to parse only that workflow.",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			repository, err := parser.Repository(parser.RepositoryInput(repo))
			if err != nil {
				return fmt.Errorf("failed to parse repository: %w", err)
			}

			client, err := gh.NewGitHubClientWithRepo(repository)
			if err != nil {
				return fmt.Errorf("failed to create GitHub client: %w", err)
			}

			// Create fallback client for github.com when the primary host is GitHub Enterprise
			var fallbackClient *gh.GitHubClient
			if repository.Host != "" && repository.Host != "github.com" {
				fc, fcErr := gh.NewGitHubClientForDefaultHost()
				if fcErr == nil {
					fallbackClient = fc
				}
			}

			var refPtr *string
			if ref != "" {
				refPtr = &ref
			}

			ctx := context.Background()
			var deps []parser.WorkflowDependency
			if len(args) > 0 {
				// Resolve selector to a specific workflow file path, then parse only that file
				filePath, resolveErr := gh.ResolveWorkflowFilePath(ctx, client, repository, args[0])
				if resolveErr != nil {
					return fmt.Errorf("failed to resolve workflow selector %q: %w", args[0], resolveErr)
				}
				deps, err = gh.GetWorkflowFileDependency(ctx, client, repository, filePath, refPtr, recursive, fallbackClient)
			} else {
				// No selector - fetch all workflow files
				deps, err = gh.GetRepositoryWorkflowDependencies(ctx, client, repository, refPtr, recursive, fallbackClient)
			}
			if err != nil {
				return fmt.Errorf("failed to get workflow dependencies: %w", err)
			}

			renderer := render.NewRenderer(opts.Exporter)
			if nameOnly {
				refs := gh.FlattenWorkflowDependencies(deps)
				renderer.RenderNames(refs)
			} else {
				if len(fields) == 0 {
					fields = []string{"Name", "Version"}
				}
				renderer.RenderWorkflowDependenciesWithFormat(format, deps, fields)
			}
			return nil
		},
	}
	f := cmd.Flags()
	f.BoolVar(&nameOnly, "name-only", false, "Output only action names")
	f.BoolVarP(&recursive, "recursive", "r", false, "Recursively traverse referenced action repositories")
	f.StringVarP(&repo, "repo", "R", "", "The repository in the format 'owner/repo'")
	f.StringVar(&ref, "ref", "", "Git reference (branch, tag, or commit SHA) to read workflow files from")
	f.StringSliceVar(&fields, "fields", nil, `Comma-separated list of fields to display in table output (default: Name,Version). Available fields: Name, Version, Owner, Repo, Path, Raw, Using, Node_Version`)

	// Use AddFormatFlags to set up --format, --jq, --template with PreRunE
	cmdutil.AddFormatFlags(cmd, &opts.Exporter)
	// Setup format flag to also accept "mermaid" and handle non-JSON format validation
	cmdflags.SetupFormatFlagWithNonJSONFormats(cmd, &opts.Exporter, &format, "", []string{"mermaid", "markdown"})
	return cmd
}
