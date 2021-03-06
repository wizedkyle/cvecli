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
	Environment  string `json:"environment"`
}

var (
	client             *cveservices_go_sdk.APIClient
	CveServicesProdUrl = "https://cveawg.mitre.org/api"
	CveServicesDevUrl  = "https://cveawg-test.mitre.org/api"
	credentialFilePath = ".cvecli/credentials/creds.json"
)

func Path(credentialFile bool) string {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("unable to retrieve user home directory")
	}
	if credentialFile == true {
		configFile := filepath.Join(homeDirectory, credentialFilePath)
		return configFile
	}
	return ""
}
