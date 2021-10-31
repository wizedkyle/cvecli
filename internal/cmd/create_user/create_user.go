package create_user

import (
	"fmt"
	"github.com/wizedkyle/cvesub/internal/logging"
	"os"

	"github.com/spf13/cobra"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"github.com/wizedkyle/cvesub/internal/cmdutils"
)

func NewCmdCreateUser(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var (
		firstName  string
		lastName   string
		middleName string
		output     string
		roles      []string
		suffix     string
		username   string
	)
	cmd := &cobra.Command{
		Use:   "create-user",
		Short: "Create a new user in the organization.",
		Long:  "The user account that is being used to authenticate needs to have the ADMIN role to perform this action.",
		Run: func(cmd *cobra.Command, args []string) {
			createUser(client, firstName, lastName, middleName, roles, suffix, username, output)
		},
	}
	cmd.Flags().StringVar(&firstName, "firstname", "", "Specify the first name of the user.")
	cmd.Flags().StringVar(&lastName, "lastname", "", "Specify the last name of the user.")
	cmd.Flags().StringVar(&middleName, "middlename", "", "Specify the middle name of the user (if applicable).")
	cmd.Flags().StringSliceVar(&roles, "roles", []string{}, "Specify the roles for the user comma separated. "+
		"Valid roles are: 'ADMIN'. Only add the user as an ADMIN if you want them to have control over the organization.")
	cmd.Flags().StringVar(&suffix, "suffix", "", "Specify the suffix of the user (if applicable).")
	cmd.Flags().StringVar(&username, "username", "", "Specify the email address of the user.")
	cmd.Flags().StringVar(&output, "output", "", "Specify a specific value to output. Accepted values are: "+
		"uuid, secret")
	cmd.MarkFlagRequired("firstName")
	cmd.MarkFlagRequired("lastName")
	cmd.MarkFlagRequired("username")
	return cmd
}

func createUser(client *cveservices_go_sdk.APIClient, firstName string, lastName string, middleName string,
	roles []string, suffix string, username string, output string) {
	if output != "" && outputValidation(output) == false {
		logging.ConsoleLogger().Error().Msg("Please select a valid output.")
		os.Exit(1)
	}

	data, response, err := client.CreateUser(types.CreateUserRequest{
		Username: username,
		Name: &types.OrgShortNameUserName{
			First:  firstName,
			Last:   lastName,
			Middle: middleName,
			Suffix: suffix,
		},
		Authority: &types.OrgShortNameUserAuthority{
			ActiveRoles: roles,
		},
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if output == "uuid" {
			fmt.Println(data.Created.UUID)
		} else if output == "secret" {
			fmt.Println(data.Created.Secret)
		} else {
			fmt.Println(data.Message)
		}
	}
}
func outputValidation(output string) bool {
	switch output {
	case
		"uuid",
		"secret":
		return true
	}
	return false
}
