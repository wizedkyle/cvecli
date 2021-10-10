package get_user

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cvesub/internal/cmdutils"
	"github.com/wizedkyle/cvesub/internal/logging"
	"os"
)

func NewCmdGetUser(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var (
		username string
		output   string
	)
	cmd := &cobra.Command{
		Use:   "get-user",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			getUser(client, username, output)
		},
	}
	cmd.Flags().StringVar(&username, "username", "", "Specify the username of the user to retrieve.")
	cmd.Flags().StringVar(&output, "output", "", "Specify a specific value to output. Accepted values are: "+
		"active, activeroles, name, orguuid, username, uuid")
	cmd.MarkFlagRequired("username")
	return cmd
}

func getUser(client *cveservices_go_sdk.APIClient, username string, output string) {
	if outputValidation(output) == false {
		logging.ConsoleLogger().Error().Msg("Please select a valid output.")
		os.Exit(1)
	}
	data, response, err := client.GetUser(username)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if output == "active" {
			fmt.Println(data.Active)
		} else if output == "activerole" {
			for _, role := range data.Authority.ActiveRoles {
				fmt.Println(role)
			}
		} else if output == "name" {
			var name string
			if data.Name.Suffix != "" {
				name = name + data.Name.Suffix + " "
			}
			if data.Name.First != "" {
				name = name + data.Name.First + " "
			}
			if data.Name.Middle != "" {
				name = name + data.Name.Middle + " "
			}
			if data.Name.Last != "" {
				name = name + data.Name.Last
			}
			fmt.Println(name)
		} else if output == "orguuid" {
			fmt.Println(data.OrgUUID)
		} else if output == "username" {
			fmt.Println(data.Username)
		} else if output == "uuid" {
			fmt.Println(data.UUID)
		} else {
			dataJson, err := json.MarshalIndent(data, "", "    ")
			if err != nil {
				logging.ConsoleLogger().Error().Err(err).Msg("failed to marshall response")
			}
			fmt.Println(string(dataJson))
		}
	}
}

func outputValidation(output string) bool {
	switch output {
	case
		"active",
		"activeroles",
		"name",
		"orguuid",
		"username",
		"uuid":
		return true
	}
	return false
}
