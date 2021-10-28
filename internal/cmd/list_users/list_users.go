package list_users

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"github.com/wizedkyle/cvesub/internal/cmdutils"
)

func NewCmdListUsers(client *cveservices_go_sdk.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-users",
		Short: "Retrieves all users from the organization",
		Run: func(cmd *cobra.Command, args []string) {
			listUsers(client)
		},
	}
	return cmd
}

func listUsers(client *cveservices_go_sdk.APIClient) {
	var (
		options         types.ListUsersOpts
		paginationToken int32
	)
	data, response, err := client.ListUsers(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		fmt.Println(string(cmdutils.OutputJson(data)))
	}
	paginationToken = data.NextPage
	for paginationToken != 0 {
		data = listUsersPagination(client, options, paginationToken)
		paginationToken = data.NextPage
	}
}

func listUsersPagination(client *cveservices_go_sdk.APIClient, options types.ListUsersOpts, token int32) types.ListUsersResponse {
	options.Page = optional.NewInt32(token)
	data, response, err := client.ListUsers(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		fmt.Println(string(cmdutils.OutputJson(data)))
	}
	return data
}
