package reset_secret

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	"github.com/wizedkyle/cveservices-go-sdk"
)

func NewCmdResetSecret(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var username string
	cmd := &cobra.Command{
		Use:   "reset-secret",
		Short: "Resets the secret for a user in the organization.",
		Long:  "The user account that is being used to authenticate needs to have the ADMIN role to perform this action.",
		Run: func(cmd *cobra.Command, args []string) {
			resetSecret(client, username)
		},
	}
	cmd.Flags().StringVar(&username, "username", "", "Specify the username which needs the secret reset.")
	cmd.MarkFlagRequired("username")
	return cmd
}

func resetSecret(client *cveservices_go_sdk.APIClient, username string) {
	data, response, err := client.ResetSecret(username)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		fmt.Println(data.APISecret)
	}
}
