package get_organization_info

import (
	"fmt"
	"github.com/spf13/cobra"
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
			getOrganizationInfo(client, output)
		},
	}
	cmd.Flags().StringVar(&output, "output", "", "Specify a specific value to output. Accepted values are: "+
		"activeroles, idquota, name, shortname, uuid")
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
		if output == "activeroles" {
			for _, role := range data.Authority.ActiveRoles {
				fmt.Println(role)
			}
		} else if output == "idquota" {
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
		"activeroles",
		"idquota",
		"name",
		"shortname",
		"uuid":
		return true
	}
	return false
}
