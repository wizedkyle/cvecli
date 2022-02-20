package list_cve_ids

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
)

func NewCmdListCveIds(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	var (
		cveIdYear int32
		state     string
	)
	cmd := &cobra.Command{
		Use:   "list-cve-ids",
		Short: "Lists all CVE IDs associated to an organization",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			listCveIds(client, cveIdYear, state, jsonOutput)
		},
	}
	cmd.Flags().StringVarP(&state, "state", "s", "", "Specify the state of the CVE ID. "+
		"Valid values are: RESERVED, PUBLISHED, REJECTED.")
	cmd.Flags().Int32VarP(&cveIdYear, "cve-id-year", "c", 0, "Specify the year of the CVE ID.")
	return cmd
}

func listCveIds(client *cveservices_go_sdk.APIClient, cveIdYear int32, state string, jsonOutput *bool) {
	var (
		options         types.ListCveIdsOpts
		paginationToken int32
	)
	if cveIdYear != 0 {
		options.CveIdYear = optional.NewInt32(cveIdYear)
	}
	if state != "" {
		if stateValidation(state) == true {
			options.State = optional.NewString(state)
		} else {
			fmt.Println("Please enter a valid CVE ID state. Valid states are: RESERVED, PUBLISHED, REJECTED.")
			os.Exit(1)
		}
	}
	data, response, err := client.ListCveIds(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if !*jsonOutput {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "CVE ID\tCVE YEAR\tSTATE\tOWNING CNA\tRESERVED DATE")
			for i := 0; i < len(*data.CveIds); i++ {
				fmt.Fprintln(writer, (*data.CveIds)[i].CveId+"\t"+(*data.CveIds)[i].CveYear+"\t"+(*data.CveIds)[i].State+
					"\t"+(*data.CveIds)[i].OwningCNA+"\t"+(*data.CveIds)[i].Reserved.String())
			}
			writer.Flush()
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
	paginationToken = data.NextPage
	for paginationToken != 0 {
		data = listCveIdsPagination(client, options, paginationToken, jsonOutput)
		paginationToken = data.NextPage
	}
}

func listCveIdsPagination(client *cveservices_go_sdk.APIClient, options types.ListCveIdsOpts, token int32, jsonOutput *bool) types.ListCveIdsResponse {
	options.Page = optional.NewInt32(token)
	data, response, err := client.ListCveIds(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if !*jsonOutput {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			for i := 0; i < len(*data.CveIds); i++ {
				fmt.Fprintln(writer, (*data.CveIds)[i].CveId+"\t"+(*data.CveIds)[i].CveYear+"\t"+(*data.CveIds)[i].State+"\t"+(*data.CveIds)[i].OwningCNA+"\t"+(*data.CveIds)[i].Reserved.String())
			}
			writer.Flush()
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
