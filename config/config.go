package config

import (
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
	cveListRemote      = "https://github.com/CVEProject/cvelist.git"
	cveServicesProdUrl = "https://cveawg.mitre.org/api"
	cveServicesDevUrl  = "https://cveawg-test.mitre.org/api"
	credentialFilePath = ".cvesub/credentials/creds.json"
	repoFilePath       = "./cvesub/repos"
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
