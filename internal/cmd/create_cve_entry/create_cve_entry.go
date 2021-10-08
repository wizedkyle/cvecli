package create_cve_entry

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"github.com/wizedkyle/cvesub/config"
	"github.com/wizedkyle/cvesub/internal/authentication"
	"github.com/wizedkyle/cvesub/internal/cmdutils"
)

func NewCmdCreateCveEntry(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var (
		cveYear    int32
		githubRepo string
	)
	cmd := &cobra.Command{
		Use:   "create-cve-entry",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			createCveEntry(client, cveYear, githubRepo)
		},
	}
	cmd.Flags().Int32Var(&cveYear, "cveYear", 0, "The year the CVE IDs will be reserved for.")
	cmd.Flags().StringVar(&githubRepo, "githubRepo", "", "Specify your fork of the CVEProject/cvelist repo.")
	cmd.MarkFlagRequired("cveYear")
	return cmd
}

func createCveEntry(client *cveservices_go_sdk.APIClient, cveYear int32, githubRepo string) {
	//var cveEntry types.CveJson4
	data, response, err := client.ReserveCveId(1, cveYear, &types.ReserveCveIdOpts{})
	if err != nil {
		cmdutils.OutputError(response, err)
	}
	fmt.Println(data.CveIds)
	_, err = git.PlainClone(config.Path(false, true), false, &git.CloneOptions{
		URL: githubRepo,
		Auth: &http.BasicAuth{
			Username: authentication.ReadGitHubCredentials(true, false),
			Password: authentication.ReadGitHubCredentials(false, true),
		},
		RemoteName: config.GetCveListRemote(),
	})
	// ask for information regarding the CVE entry and update the cveEntry array
	// clone repo that is a fork of cvelist to home dir .cvesub/repos/
	// set upstream, update etc
	// Find cve id json file
	// update cve id json file with details from cveEntry
	// commit changes
	// push changes to repo
	// create PR

}
