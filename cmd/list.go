package cmd

import (
	"context"
	"fmt"

	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/srz-zumix/go-gh-extension/pkg/gh"
	"github.com/srz-zumix/go-gh-extension/pkg/parser"
	"github.com/srz-zumix/go-gh-extension/pkg/render"
)

type ListOptions struct {
	Exporter cmdutil.Exporter
}

// NewListCmd returns the actions list command
func NewListCmd() *cobra.Command {
	var repo string
	var ecosystem string
	var nameOnly bool
	opts := &ListOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List dependency packages",
		Long:  "List dependency packages in the repository's SBOM.",
		RunE: func(cmd *cobra.Command, args []string) error {
			repository, err := parser.Repository(parser.RepositoryInput(repo))
			if err != nil {
				return fmt.Errorf("failed to parse repository: %w", err)
			}

			client, err := gh.NewGitHubClientWithRepo(repository)
			if err != nil {
				return fmt.Errorf("failed to create GitHub client: %w", err)
			}

			sbom, err := gh.GetRepositoryDependencyGraphSBOM(context.Background(), client, repository)
			if err != nil {
				return fmt.Errorf("failed to get SBOM: %w", err)
			}

			if ecosystem != "" {
				sbom = gh.FilterSBOMPackage(sbom, ecosystem)
			}

			renderer := render.NewRenderer(opts.Exporter)
			if nameOnly {
				renderer.RenderNames(sbom.SBOM.Packages)
			} else {
				if ecosystem != "" {
					renderer.RenderSBOMPackages(sbom, []string{"Name", "Version"})
				} else {
					renderer.RenderSBOMPackagesDefault(sbom)
				}
			}
			return nil
		},
	}
	f := cmd.Flags()
	f.BoolVar(&nameOnly, "name-only", false, "Output only team names")
	f.StringVarP(&ecosystem, "ecosystem", "e", "", "The ecosystem of the dependencies")
	f.StringVarP(&repo, "repo", "R", "", "The repository in the format 'owner/repo'")
	cmdutil.AddFormatFlags(cmd, &opts.Exporter)
	return cmd
}
