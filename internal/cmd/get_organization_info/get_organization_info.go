package get_organization_info

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	"github.com/wizedkyle/cvecli/internal/logging"
	"github.com/wizedkyle/cveservices-go-sdk"
	"os"
)

func NewCmdGetOrganizationInfo(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var output string
	cmd := &cobra.Command{
		Use:   "get-organization-info",
		Short: "Retrieves information about the organization the user authenticating is apart of.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getOrganizationInfo(client, output)
		},
	}
	cmd.Flags().StringVarP(&output, "output", "o", "", "Specify a specific value to output. Accepted values are: "+
		"active-roles, id-quota, name, shortname, uuid")
	return cmd
}

func getOrganizationInfo(client *cveservices_go_sdk.APIClient, output string) {
	if output != "" && outputValidation(output) == false {
		logging.ConsoleLogger().Error().Msg("Please select a valid output.")
		os.Exit(1)
	}
	data, response, err := client.GetOrganizationRecord()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if output == "active-roles" {
			for _, role := range data.Authority.ActiveRoles {
				fmt.Println(role)
			}
		} else if output == "id-quota" {
			fmt.Println(data.Policies.IdQuota)
		} else if output == "name" {
			fmt.Println(data.Name)
		} else if output == "shortname" {
			fmt.Println(data.ShortName)
		} else if output == "uuid" {
			fmt.Println(data.UUID)
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
}

func outputValidation(output string) bool {
	switch output {
	case
		"active-roles",
		"id-quota",
		"name",
		"shortname",
		"uuid":
		return true
	}
	return false
}
