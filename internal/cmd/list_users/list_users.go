package list_users

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"os"
	"strconv"
	"text/tabwriter"
)

func NewCmdListUsers(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-users",
		Short: "Retrieves all users from the organization",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			listUsers(client, jsonOutput)
		},
	}
	return cmd
}

func listUsers(client *cveservices_go_sdk.APIClient, jsonOutput *bool) {
	var (
		options         types.ListUsersOpts
		paginationToken int32
	)
	data, response, err := client.ListUsers(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if *jsonOutput == false {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "FIRST NAME\tLAST NAME\tUSERNAME\tUUID\tACTIVE")
			for i := 0; i < len(*data.Users); i++ {
				fmt.Fprintln(writer, (*data.Users)[i].Name.First+"\t"+(*data.Users)[i].Name.Last+"\t"+(*data.Users)[i].Username+
					"\t"+(*data.Users)[i].UUID+"\t"+strconv.FormatBool((*data.Users)[i].Active))
			}
			writer.Flush()
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
	paginationToken = data.NextPage
	for paginationToken != 0 {
		data = listUsersPagination(client, options, paginationToken, jsonOutput)
		paginationToken = data.NextPage
	}
}

func listUsersPagination(client *cveservices_go_sdk.APIClient, options types.ListUsersOpts, token int32, jsonOutput *bool) types.ListUsersResponse {
	options.Page = optional.NewInt32(token)
	data, response, err := client.ListUsers(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if *jsonOutput == false {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			for i := 0; i < len(*data.Users); i++ {
				fmt.Fprintln(writer, (*data.Users)[i].Name.First+"\t"+(*data.Users)[i].Name.Last+"\t"+(*data.Users)[i].Username+
					"\t"+(*data.Users)[i].UUID+"\t"+strconv.FormatBool((*data.Users)[i].Active))
			}
			writer.Flush()
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
	return data
}
