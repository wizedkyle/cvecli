package list_users

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	"github.com/wizedkyle/cvecli/internal/logging"
	"github.com/wizedkyle/cvecli/internal/validation"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"os"
	"strings"
)

func NewCmdListUsers(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var output string
	cmd := &cobra.Command{
		Use:   "list-users",
		Short: "Retrieves all users from the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			listUsers(client, output)
		},
	}
	cmd.Flags().StringVarP(&output, "output", "o", "", "Specify a specific value to output. Accepted values are: "+
		"active, active-roles, name, uuid")
	return cmd
}

func listUsers(client *cveservices_go_sdk.APIClient, output string) {
	var (
		options         types.ListUsersOpts
		paginationToken int32
	)
	if output != "" && validation.ListUserOutputValidation(output) == false {
		logging.ConsoleLogger().Error().Msg("Please select a valid output.")
		os.Exit(1)
	}
	outputLower := strings.ToLower(output)
	data, response, err := client.ListUsers(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		for _, user := range *data.Users {
			if outputLower == "active" {
				fmt.Println(user.Username, user.Active)
			} else if outputLower == "active-roles" {
				fmt.Println(user.Username, user.Authority.ActiveRoles)
			} else if outputLower == "name" {
				var name string
				if user.Name.Suffix != "" {
					name = name + user.Name.Suffix + " "
				}
				if user.Name.First != "" {
					name = name + user.Name.First + " "
				}
				if user.Name.Middle != "" {
					name = name + user.Name.Middle + " "
				}
				if user.Name.Last != "" {
					name = name + user.Name.Last
				}
				fmt.Println(user.Username, name)
			} else if outputLower == "uuid" {
				fmt.Println(user.Username, user.UUID)
			} else {
				fmt.Println(string(cmdutils.OutputJson(data)))
			}
		}
	}
	paginationToken = data.NextPage
	for paginationToken != 0 {
		data = listUsersPagination(client, options, paginationToken, outputLower)
		paginationToken = data.NextPage
	}
}

func listUsersPagination(client *cveservices_go_sdk.APIClient, options types.ListUsersOpts, token int32, outputLower string) types.ListUsersResponse {
	options.Page = optional.NewInt32(token)
	data, response, err := client.ListUsers(&options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		for _, user := range *data.Users {
			if outputLower == "active" {
				fmt.Println(user.Username, user.Active)
			} else if outputLower == "active-roles" {
				fmt.Println(user.Username, user.Authority.ActiveRoles)
			} else if outputLower == "name" {
				var name string
				if user.Name.Suffix != "" {
					name = name + user.Name.Suffix + " "
				}
				if user.Name.First != "" {
					name = name + user.Name.First + " "
				}
				if user.Name.Middle != "" {
					name = name + user.Name.Middle + " "
				}
				if user.Name.Last != "" {
					name = name + user.Name.Last
				}
				fmt.Println(user.Username, name)
			} else if outputLower == "uuid" {
				fmt.Println(user.Username, user.UUID)
			} else {
				fmt.Println(string(cmdutils.OutputJson(data)))
			}
		}
	}
	return data
}
