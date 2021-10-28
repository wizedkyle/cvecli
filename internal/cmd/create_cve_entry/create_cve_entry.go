package create_cve_entry

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"github.com/wizedkyle/cvesub/config"
	"github.com/wizedkyle/cvesub/internal/authentication"
	"github.com/wizedkyle/cvesub/internal/cmdutils"
	"github.com/wizedkyle/cvesub/internal/logging"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func NewCmdCreateCveEntry(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var (
		assigner     string
		cveYear      int32
		description  string
		githubRepo   string
		language     string
		problemType  string
		productName  string
		referenceUrl string
		vendorName   string
	)
	cmd := &cobra.Command{
		Use:   "create-cve-entry",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			createCveEntry(client, assigner, cveYear, githubRepo)
		},
	}
	cmd.Flags().StringVar(&assigner, "assigner", "", "Specify the assigner email address.")
	cmd.Flags().Int32Var(&cveYear, "cveYear", 0, "The year the CVE IDs will be reserved for.")
	cmd.Flags().StringVar(&description, "description", "", "")
	cmd.Flags().StringVar(&githubRepo, "githubRepo", "", "Specify your fork of the CVEProject/cvelist repo.")
	cmd.Flags().StringVar(&language, "language", "", "Specify the language in ISO 639-3 format.")
	cmd.Flags().StringVar(&problemType, "problemType", "", "Specify the description of the problem type.")
	cmd.Flags().StringVar(&productName, "productName", "", "Specify the product name.")
	cmd.Flags().StringVar(&referenceUrl, "referenceUrl", "", "Specify")
	cmd.Flags().StringVar(&vendorName, "vendorName", "", "Specify the vendor name.")
	cmd.MarkFlagRequired("cveYear")
	return cmd
}

func createCveEntry(client *cveservices_go_sdk.APIClient, assigner string, cveYear int32, githubRepo string) {
	var (
		cveEntry types.CveJson4
		repoName string
	)
	data, response, err := client.ReserveCveId(1, cveYear, &types.ReserveCveIdOpts{})
	if err != nil {
		cmdutils.OutputError(response, err)
	}
	// TODO: remove print
	fmt.Println(data.CveIds)

	urlData, err := url.Parse(githubRepo)
	if err != nil {
		// TODO: update this error
		fmt.Println(err)
	}
	paths := strings.Split(urlData.Path, "/")
	if strings.Contains(paths[2], ".git") == true {
		repoNameSplit := strings.Split(paths[2], ".")
		repoName = repoNameSplit[0]
	} else {
		repoName = paths[2]
	}
	repoFolderPath := config.Path(false, true)
	repoPath := repoFolderPath + "/" + repoName + "/"
	repo := filepath.Dir(repoPath)
	if _, err := os.Stat(repo); os.IsNotExist(err) {
		err := os.MkdirAll(repo, 0755)
		if err != nil {
			logging.ConsoleLogger().Error().Err(err).Msg("failed to create folder structure for GitHub repository")
		}
	}
	pterm.Info.Println("cloning", githubRepo)
	_, err = git.PlainClone(repo, false, &git.CloneOptions{
		URL: githubRepo,
		Auth: &http.BasicAuth{
			Username: authentication.ReadGitHubCredentials(true, false),
			Password: authentication.ReadGitHubCredentials(false, true),
		},
		// TODO: add remote origin
	})
	if err != nil {
		if err.Error() == "repository already exists" {
			pterm.Info.Println("repository already exists, attempting to update local repository")
			cveRepo, err := git.PlainOpen(repo)
			if err != nil {
				// TODO: fix error
				fmt.Println("failed to open repository")
			}
			tree, err := cveRepo.Worktree()
			if err != nil {
				// TODO: fix error
				fmt.Println("failed to fine tree")
			}
			err = tree.Pull(&git.PullOptions{
				RemoteName: "origin",
			})
			if err != nil {
				// TODO: fix error
				fmt.Println("failed to pull changes from origin")
				fmt.Println(err)
			} else {
				gitRef, err := cveRepo.Head()
				if err != nil {
					// TODO: fix error
					fmt.Println("failed to get git ref")
				}
				commit, err := cveRepo.CommitObject(gitRef.Hash())
				if err != nil {
					// TODO: fix error
					fmt.Println("failed to get commit hash")
				}
				fmt.Println(commit)
			}
		} else {
			pterm.Error.WithFatal().Println()
		}
	}
	cveEntry.CVEDataMeta.ID = data.CveIds[0].CveId
	// ask for information regarding the CVE entry and update the cveEntry array
	// set upstream, update etc
	// Find cve id json file
	// update cve id json file with details from cveEntry
	// commit changes
	// push changes to repo
	// create PR

}

func referenceSourceValidation(referenceSource string) bool {

}
