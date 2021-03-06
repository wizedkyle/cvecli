package reset_secret

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
)

func NewCmdResetSecret(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	var username string
	cmd := &cobra.Command{
		Use:   "reset-secret",
		Short: "Resets the secret for a user in the organization",
		Long:  "The user account that is being used to authenticate needs to have the ADMIN role to perform this action.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			resetSecret(client, username, jsonOutput)
		},
	}
	cmd.Flags().StringVarP(&username, "username", "u", "", "Specify the username which needs the secret reset.")
	cmd.MarkFlagRequired("username")
	return cmd
}

func resetSecret(client *cveservices_go_sdk.APIClient, username string, jsonOutput *bool) {
	data, response, err := client.ResetSecret(username)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if !*jsonOutput {
			fmt.Println(data.APISecret)
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
}
