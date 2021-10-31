package config

import (
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cvesub/internal/logging"
	"os"
	"path/filepath"
)

type CredentialFile struct {
	APIUser        string `json:"apiUser"`
	APIKey         string `json:"apiKey"`
	Organization   string `json:"organization"`
	GitHubUsername string `json:"githubUsername"`
	GitHubPat      string `json:"githubPat"`
}

var (
	client             *cveservices_go_sdk.APIClient
	cveListRemote      = "https://github.com/CVEProject/cvelist.git"
	cveServicesProdUrl = "https://cveawg.mitre.org/api"
	cveServicesDevUrl  = "https://cveawg-test.mitre.org/api"
	credentialFilePath = ".cvecli/credentials/creds.json"
	repoFilePath       = ".cvecli/repos"
)

func Path(credentialFile bool, repoPath bool) string {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("unable to retrieve user home directory")
	}
	if credentialFile == true {
		configFile := filepath.Join(homeDirectory, credentialFilePath)
		return configFile
	} else if repoPath == true {
		configFile := filepath.Join(homeDirectory, repoFilePath)
		return configFile
	}
	return ""
}

func GetCveListRemote() string {
	return cveListRemote
}

func GetClient() *cveservices_go_sdk.APIClient {
	return client
}

func SetClient(newClient *cveservices_go_sdk.APIClient) {
	client = newClient
}
