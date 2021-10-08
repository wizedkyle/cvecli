package configure

import (
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvesub/config"
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
			fmt.Println("cvesub requires and api user, api key, organization, GitHub Username and GitHub PAT for correct functioning. " +
				"cvesub will encrypt and store the api user, api key, organization, GitHub username and GitHub PAT in " +
				"the following file for use by subsequent commands: " + config.Path(true, false))
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
	var credentials config.CredentialFile
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
	promptGitHubUsername := promptui.Prompt{
		Label: "Please enter your GitHub username",
	}
	promptGitHubPat := promptui.Prompt{
		Label: "Please enter a GitHub Personal Access Token (PAT)",
		Mask:  '*',
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
	githubUsername, err := promptGitHubUsername.Run()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to prompt for a GitHub username")
	}
	githubPat, err := promptGitHubPat.Run()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to prompt for a GitHub PAT")
	}
	credentials.APIUser = encryption.EncryptData(apiUser)
	credentials.APIKey = encryption.EncryptData(apiKey)
	credentials.Organization = encryption.EncryptData(organization)
	credentials.GitHubUsername = encryption.EncryptData(githubUsername)
	credentials.GitHubPat = encryption.EncryptData(githubPat)
	configFilePath := filepath.Dir(config.Path(true, false))
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(configFilePath, 0755)
		if err != nil {
			logging.ConsoleLogger().Error().Err(err).Msg("failed to create folder structure for credentails file")
		}
	}
	credentialsFile, err := json.MarshalIndent(credentials, "", "    ")
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to marshal credentials")
	}
	err = os.WriteFile(config.Path(true, false), credentialsFile, 0644)
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to write credentials file to " + config.Path(true, false))
	} else {
		fmt.Println("Credentials file saved to: " + config.Path(true, false))
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
