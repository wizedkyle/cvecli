package config

import (
	"github.com/wizedkyle/cvecli/internal/logging"
	"github.com/wizedkyle/cveservices-go-sdk"
	"os"
	"path/filepath"
)

type CredentialFile struct {
	APIUser      string `json:"apiUser"`
	APIKey       string `json:"apiKey"`
	Organization string `json:"organization"`
}

var (
	client                *cveservices_go_sdk.APIClient
	CveServicesProdUrl    = "https://cveawg.mitre.org/api"
	CveServicesDevUrl     = "https://cveawg-test.mitre.org/api"
	credentialFilePath    = ".cvecli/credentials/creds.json"
	ProductionEnvironment = false
	repoFilePath          = ".cvecli/repos"
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

func GetClient() *cveservices_go_sdk.APIClient {
	return client
}

func SetClient(newClient *cveservices_go_sdk.APIClient) {
	client = newClient
}
