package create_user

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
)

func NewCmdCreateUser(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	var (
		firstName  string
		lastName   string
		middleName string
		roles      []string
		suffix     string
		username   string
	)
	cmd := &cobra.Command{
		Use:   "create-user",
		Short: "Creates a new user in the organization",
		Long:  "The user account that is being used to authenticate needs to have the ADMIN role to perform this action.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			createUser(client, firstName, lastName, middleName, roles, suffix, username, jsonOutput)
		},
	}
	cmd.Flags().StringVarP(&firstName, "first-name", "f", "", "Specify the first name of the user")
	cmd.Flags().StringVarP(&lastName, "last-name", "l", "", "Specify the last name of the user")
	cmd.Flags().StringVarP(&middleName, "middle-name", "m", "", "Specify the middle name of the user (if applicable)")
	cmd.Flags().StringSliceVarP(&roles, "roles", "r", []string{}, "Specify the roles for the user comma separated. "+
		"Valid roles are: 'ADMIN'. Only add the user as an ADMIN if you want them to have control over the organization.")
	cmd.Flags().StringVarP(&suffix, "suffix", "s", "", "Specify the suffix of the user (if applicable)")
	cmd.Flags().StringVarP(&username, "username", "u", "", "Specify the email address of the user")
	cmd.MarkFlagRequired("first-name")
	cmd.MarkFlagRequired("last-name")
	cmd.MarkFlagRequired("username")
	return cmd
}

func createUser(client *cveservices_go_sdk.APIClient, firstName string, lastName string, middleName string,
	roles []string, suffix string, username string, jsonOutput *bool) {
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
		if !*jsonOutput {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "FIRST NAME\tLAST NAME\tUSERNAME\tUUID\tACTIVE\tSECRET")
			fmt.Fprintln(writer, data.Created.Name.First+"\t"+data.Created.Name.Last+
				"\t"+data.Created.Username+"\t"+data.Created.UUID+"\t"+strconv.FormatBool(data.Created.Active)+"\t"+data.Created.Secret)
			writer.Flush()
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
}
