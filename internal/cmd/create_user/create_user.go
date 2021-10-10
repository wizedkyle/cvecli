package create_user

import (
	"fmt"

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
		orgUUID    string
		roles      []string
		suffix     string
		username   string
	)
	cmd := &cobra.Command{
		Use:   "create-user",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			createUser(client, firstName, lastName, middleName, orgUUID, roles, suffix, username)
		},
	}
	cmd.Flags().StringVar(&firstName, "firstName", "", "Specify the first name of the user.")
	cmd.Flags().StringVar(&lastName, "lastName", "", "Specify the last name of the user.")
	cmd.Flags().StringVar(&middleName, "middleName", "", "Specify the middle name of the user (if applicable).")
	cmd.Flags().StringVar(&orgUUID, "orgUUID", "", "Specify the organization UUID.")
	cmd.Flags().StringSliceVar(&roles, "roles", []string{}, "Specify the roles for the user comma separated. "+
		"Valid roles are: ")
	cmd.MarkFlagRequired("firstName")
	cmd.MarkFlagRequired("lastName")
	cmd.MarkFlagRequired("orgUUID")
	return cmd
}

func createUser(client *cveservices_go_sdk.APIClient, firstName string, lastName string, middleName string, orgUUID string,
	roles []string, suffix string, username string) {
	data, response, err := client.CreateUser(types.CreateUserRequest{
		Username: username,
		OrgUUID:  orgUUID,
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
		fmt.Println(data)
	}
}
