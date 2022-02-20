package update_user

import (
	"fmt"
	"os"
	"strings"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
)

func NewCmdUpdateUser(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	var (
		active       bool
		firstName    string
		lastName     string
		middleName   string
		newUsername  string
		username     string
		suffix       string
		roleToAdd    string
		roleToRemove string
	)
	cmd := &cobra.Command{
		Use:   "update-user",
		Short: "Updates a user from the organization",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			updateUser(client, active, firstName, lastName, middleName, newUsername, username, suffix, roleToAdd, roleToRemove,
				jsonOutput)
		},
	}
	cmd.Flags().StringVarP(&roleToAdd, "role-to-add", "a", "", "Specify the role for the user. "+
		"Valid roles are: 'ADMIN'. Only add the user as an ADMIN if you want them to have control over the organization")
	cmd.Flags().BoolVarP(&active, "enabled", "e", true, "Set to false if you want to disable the user")
	cmd.Flags().StringVarP(&firstName, "first-name", "f", "", "Specify the first name of the user")
	cmd.Flags().StringVarP(&lastName, "last-name", "l", "", "Specify the last name of the user")
	cmd.Flags().StringVarP(&middleName, "middle-name", "m", "", "Specify the middle name of the user (if applicable)")
	cmd.Flags().StringVarP(&newUsername, "new-username", "n", "", "Specify the new email address of the user")
	cmd.Flags().StringVarP(&username, "username", "u", "", "Specify the current email address of the user")
	cmd.Flags().StringVarP(&roleToRemove, "role-to-remove", "r", "", "Specify the role to remove from the user. "+
		"Valid roles are: 'ADMIN'")
	cmd.Flags().StringVarP(&suffix, "suffix", "s", "", "Specify the suffix of the user (if applicable)")
	return cmd
}

func updateUser(client *cveservices_go_sdk.APIClient, active bool, firstName string, lastName string, middleName string,
	newUsername string, username string, suffix string, roleToAdd string, roleToRemove string, jsonOutput *bool) {
	var (
		options types.UpdateUserOpts
	)
	data, response, err := client.GetUser(username)
	if err != nil {
		cmdutils.OutputError(response, err)
	}
	if data.Active == active {
		options.Active = optional.NewBool(data.Active)
	} else {
		options.Active = optional.NewBool(active)
	}
	if firstName != "" {
		options.NameFirst = optional.NewString(firstName)
	} else {
		options.NameFirst = optional.NewString(data.Name.First)
	}
	if lastName != "" {
		options.NameLast = optional.NewString(lastName)
	} else {
		options.NameLast = optional.NewString(data.Name.Last)
	}
	if middleName != "" {
		options.NameMiddle = optional.NewString(middleName)
	} else {
		options.NameMiddle = optional.NewString(data.Name.Middle)
	}
	if suffix != "" {
		options.NameSuffix = optional.NewString(suffix)
	} else {
		options.NameSuffix = optional.NewString(data.Name.Suffix)
	}
	if newUsername != "" {
		options.NewUsername = optional.NewString(newUsername)
	}
	if roleToAdd != "" {
		if roleValidation(roleToAdd) == true {
			options.ActiveRolesAdd = optional.NewString(roleToAdd)
		} else {
			fmt.Println("Please enter a valid role. Valid roles are: ADMIN.")
			os.Exit(1)
		}
	}
	if roleToRemove != "" {
		if roleValidation(roleToRemove) == true {
			options.ActiveRolesRemove = optional.NewString(roleToRemove)
		} else {
			fmt.Println("Please enter a valid role. Valid roles are: ADMIN.")
			os.Exit(1)
		}
	}

	updateData, response, err := client.UpdateUser(username, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if !*jsonOutput {
			fmt.Println(updateData.Message)
		} else {
			fmt.Println(string(cmdutils.OutputJson(updateData)))
		}
	}
}

func roleValidation(role string) bool {
	role = strings.ToUpper(role)
	switch role {
	case
		"ADMIN":
		return true
	}
	return false
}
