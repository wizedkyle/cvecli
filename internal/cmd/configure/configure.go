package configure

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/config"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/logging"
)

func NewCmdConfigure() *cobra.Command {
	var (
		showApiUser      bool
		showOrganization bool
	)
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Sets credentials for cvecli",
		Long:  "Interactively sets the cvecli API user, API key, organization information.",
		Run: func(cmd *cobra.Command, args []string) {
			if showApiUser == true && showOrganization == true {
				fmt.Println("Please select either show-api-user or show-organization.")
				os.Exit(1)
			} else if showApiUser == true {
				apiUser := authentication.ReadApiUser()
				if apiUser != "" {
					fmt.Println("The API user currently configured for authentication is: " + apiUser)
					os.Exit(0)
				} else {
					os.Exit(0)
				}
			} else if showOrganization == true {
				organization := authentication.ReadOrganization()
				if organization != "" {
					fmt.Println("The organization currently configured for authentication is: " + organization)
					os.Exit(0)
				} else {
					os.Exit(0)
				}
			}
			fmt.Println("cvecli requires and api user, api key and organization for correct functioning. " +
				"cvecli will encrypt and store the api user, api key and organization in " +
				"the following file for use by subsequent commands: " + config.Path(true))
			if userConfirmation() == true {
				setCredentials()
			} else {
				os.Exit(1)
			}
		},
	}
	cmd.Flags().BoolVarP(&showOrganization, "show-organization", "o", false, "Shows the plaintext organization name. This command "+
		"is useful for identifying which CVE CNA organization is being used. If this flag is set you cannot configure credentials.")
	cmd.Flags().BoolVarP(&showApiUser, "show-api-user", "u", false, "Shows the plain text API user. This command "+
		"is useful for identifying which API user is being used. If this flag is set you cannot configure credentials.")
	return cmd
}

func setCredentials() {
	var environment string
	promptEnvironment := promptui.Prompt{
		Label:     "Do you want to access the production CVE Servers environment? If you select no the test environment will be used.",
		IsConfirm: true,
	}
	promptApiUser := promptui.Prompt{
		Label: "Please enter your api username",
	}
	promptApiKey := promptui.Prompt{
		Label: "Please enter your api access key",
		Mask:  '*',
	}
	promptOrganization := promptui.Prompt{
		Label: "Please enter your CNA organization name",
	}
	_, err := promptEnvironment.Run()
	if err != nil {
		fmt.Println("Using the CVE Services test environment")
		environment = config.CveServicesDevUrl
	} else {
		fmt.Println("Using the CVE Services production environment")
		environment = config.CveServicesProdUrl
	}
	apiUser, err := promptApiUser.Run()
	if err != nil {
		logging.Console().Error().Err(err).Msg("failed to prompt an api user")
	}
	apiKey, err := promptApiKey.Run()
	if err != nil {
		logging.Console().Error().Err(err).Msg("failed to prompt for an api key")
	}
	organization, err := promptOrganization.Run()
	if err != nil {
		logging.Console().Error().Err(err).Msg("failed to prompt for an organization")
	}
	authentication.WriteCredentialsFile(apiUser, apiKey, organization, environment)
}

func userConfirmation() bool {
	prompt := promptui.Prompt{
		Label:     "Do you want to proceed",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Println("Configuration cancelled")
		return false
	}
	resultLower := strings.ToLower(result)
	if resultLower == "y" {
		return true
	}
	return false
}
