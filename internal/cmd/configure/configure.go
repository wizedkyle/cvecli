package configure

import (
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvesub/config"
	"github.com/wizedkyle/cvesub/internal/authentication"
	"github.com/wizedkyle/cvesub/internal/encryption"
	"github.com/wizedkyle/cvesub/internal/logging"
	"os"
	"path/filepath"
	"strings"
)

func NewCmdConfigure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Sets cvesub credentials",
		Long:  "Interactively sets the cvesub API user, API key and organization information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cvesub requires and api user, api key and organization for correct functioning. " +
				"cvesub will encrypt and store the api user, api key and organization in " +
				"the following file for use by subsequent commands: " + authentication.ConfigPath())
			if userConfirmation() == true {
				setCredentials()
			} else {
				os.Exit(1)
			}
		},
	}
	return cmd
}

func setCredentials() {
	var credentails config.CredentialFile
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
	credentails.APIUser = encryption.EncryptData(apiUser)
	credentails.APIKey = encryption.EncryptData(apiKey)
	credentails.Organization = encryption.EncryptData(organization)
	configFilePath := filepath.Dir(authentication.ConfigPath())
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(configFilePath, 0755)
		if err != nil {
			logging.ConsoleLogger().Error().Err(err).Msg("failed to create folder structure for credentails file")
		}
	}
	credentialsFile, err := json.MarshalIndent(credentails, "", "    ")
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to marshal credentials")
	}
	err = os.WriteFile(authentication.ConfigPath(), credentialsFile, 0644)
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to write credentials file to " + authentication.ConfigPath())
	} else {
		fmt.Println("Credentials file saved to: " + authentication.ConfigPath())
	}
}

func userConfirmation() bool {
	prompt := promptui.Prompt{
		Label: "Do you want to proceed?",
	}
	result, err := prompt.Run()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to prompt for confirmation")
	}
	resultLower := strings.ToLower(result)
	if resultLower == "yes" {
		return true
	} else {
		fmt.Println("Configuration cancelled")
		return false
	}
}
