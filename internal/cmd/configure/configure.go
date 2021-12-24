package configure

import (
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/config"
	"github.com/wizedkyle/cvecli/internal/encryption"
	"github.com/wizedkyle/cvecli/internal/logging"
	"os"
	"path/filepath"
	"strings"
)

func NewCmdConfigure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Sets credentials that cvecli requires",
		Long:  "Interactively sets the cvecli API user, API key, organization information and GitHub credentials.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cvecli requires and api user, api key, organization, GitHub Username and GitHub PAT for correct functioning. " +
				"cvecli will encrypt and store the api user, api key, organization, GitHub username and GitHub PAT in " +
				"the following file for use by subsequent commands: " + config.Path(true, false))
			if userConfirmation() == true {
				SetCredentials()
			} else {
				os.Exit(1)
			}
		},
	}
	return cmd
}

func SetCredentials() {
	var credentials config.CredentialFile
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
		credentials.Environment = config.CveServicesDevUrl
	} else {
		fmt.Println("Using the CVE Services production environment")
		credentials.Environment = config.CveServicesProdUrl
	}
	apiUser, err := promptApiUser.Run()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to prompt an api user")
	}
	apiKey, err := promptApiKey.Run()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to prompt for an api key")
	}
	organization, err := promptOrganization.Run()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to prompt for an organization")
	}
	credentials.APIUser = encryption.EncryptData(apiUser)
	credentials.APIKey = encryption.EncryptData(apiKey)
	credentials.Organization = encryption.EncryptData(organization)
	configFilePath := filepath.Dir(config.Path(true, false))
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(configFilePath, 0755)
		if err != nil {
			logging.ConsoleLogger().Error().Err(err).Msg("failed to create folder structure for credentials file")
		}
	}
	credentialsFile, err := json.MarshalIndent(credentials, "", "    ")
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to marshal credentials")
	}
	err = os.WriteFile(config.Path(true, false), credentialsFile, 0644)
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to write credentials file to " + config.Path(true, false))
		os.Exit(0)
	} else {
		fmt.Println("Credentials file saved to: " + config.Path(true, false))
		os.Exit(0)
	}
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
