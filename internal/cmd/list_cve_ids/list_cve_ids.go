package list_cve_ids

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"github.com/wizedkyle/cvesub/internal/cmdutils"
	"github.com/wizedkyle/cvesub/internal/logging"
	"os"
	"strings"
)

func NewCmdListCveIds(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var (
		cveIdYear int32
		output    string
		state     string
	)
	cmd := &cobra.Command{
		Use:   "list-cve-ids",
		Short: "Lists all CVE Ids associated to an organization.",
		Run: func(cmd *cobra.Command, args []string) {
			listCveIds(client, cveIdYear, output, state)
		},
	}
	cmd.Flags().StringVar(&state, "state", "", "Specify the state of the CVE ID. "+
		"Valid values are: RESERVED, PUBLISHED, REJECTED.")
	cmd.Flags().Int32Var(&cveIdYear, "cveIdYear", 0, "Specify the year of the CVE ID.")
	cmd.Flags().StringVar(&output, "output", "", "Specify a specific value to output. Accepted values are: cveid")
	return cmd
}

func listCveIds(client *cveservices_go_sdk.APIClient, cveIdYear int32, output string, state string) {
	var (
		options         types.ListCveIdsOpts
		paginationToken int32
	)
	if output != "" && outputValidation(output) == false {
		logging.ConsoleLogger().Error().Msg("Please select a valid output.")
		os.Exit(1)
	}
	if cveIdYear != 0 {
		options.CveIdYear = optional.NewInt32(cveIdYear)
	}
	if state != "" {
		if stateValidation(state) == true {
			options.State = optional.NewString(state)
		} else {
			fmt.Println("Please enter a valid CVE ID state. Valid states are: RESERVED, PUBLISHED, REJECTED.")
		}
	}
	data, response, err := client.ListCveIds(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if output == "cveid" {
			for _, cveid := range *data.CveIds {
				fmt.Println(cveid.CveId)
			}
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
	paginationToken = data.NextPage
	for paginationToken != 0 {
		data = listCveIdsPagination(client, options, output, paginationToken)
		paginationToken = data.NextPage
	}
}

func listCveIdsPagination(client *cveservices_go_sdk.APIClient, options types.ListCveIdsOpts, output string, token int32) types.ListCveIdsResponse {
	options.Page = optional.NewInt32(token)
	data, response, err := client.ListCveIds(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if output == "cveid" {
			for _, cveid := range *data.CveIds {
				fmt.Println(cveid.CveId)
			}
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
	return data
}

func stateValidation(state string) bool {
	state = strings.ToUpper(state)
	switch state {
	case
		"RESERVED",
		"PUBLISHED",
		"REJECTED":
		return true
	}
	return false
}

func outputValidation(output string) bool {
	switch output {
	case
		"cveid":
		return true
	}
	return false
}
