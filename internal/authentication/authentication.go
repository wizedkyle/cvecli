package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"github.com/wizedkyle/cvecli/config"
	"github.com/wizedkyle/cvecli/internal/build"
	"github.com/wizedkyle/cvecli/internal/encryption"
	"github.com/wizedkyle/cvecli/internal/logging"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
)

func ConfirmCredentialsSet(client *cveservices_go_sdk.APIClient) {
	if client.Cfg.Authentication.APIUser == "" || client.Cfg.Authentication.APIKey == "" || client.Cfg.BasePath == "" || client.Cfg.Organization == "" {
		fmt.Println("No authentication credentials set, please run cvecli configure or set environment variables.")
		os.Exit(1)
	}
}

func GetCVEServicesSDKConfig() *cveservices_go_sdk.APIClient {
	apiUser, apiKey, organization, environment := ReadAuthCredentials()
	client := cveservices_go_sdk.APIClient{
		Cfg: &cveservices_go_sdk.Configuration{
			Authentication: cveservices_go_sdk.BasicAuth{
				APIUser: apiUser,
				APIKey:  apiKey,
			},
			BasePath:     environment,
			Organization: organization,
			UserAgent:    "cvecli " + build.Version,
			HTTPClient: &http.Client{
				Timeout: time.Second * 20,
			},
		},
	}
	return &client
}

func ReadApiUser() string {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(config.Path(true)))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No authentication credentials, please run cvecli configure.")
		return ""
	} else {
		apiUser := viper.GetString("apiUser")
		apiUserDecrypted := encryption.DecryptData(apiUser)
		return apiUserDecrypted
	}
}

func ReadAuthCredentials() (string, string, string, string) {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(config.Path(true)))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		apiUser := viper.GetString("CVE_API_USER")
		apiKey := viper.GetString("CVE_API_KEY")
		organization := viper.GetString("CVE_ORGANIZATION")
		environment := viper.GetString("CVE_ENVIRONMENT")
		return apiUser, apiKey, organization, environment
	} else {
		apiUser := viper.GetString("apiUser")
		apiKey := viper.GetString("apiKey")
		organization := viper.GetString("organization")
		environment := viper.GetString("environment")
		apiUserDecrypted := encryption.DecryptData(apiUser)
		apiKeyDecrypted := encryption.DecryptData(apiKey)
		organizationDecrypted := encryption.DecryptData(organization)
		return apiUserDecrypted, apiKeyDecrypted, organizationDecrypted, environment
	}
}

func ReadOrganization() string {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(config.Path(true)))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No authentication credentials, please run cvecli configure.")
		return ""
	} else {
		organization := viper.GetString("organization")
		organizationDecrypted := encryption.DecryptData(organization)
		return organizationDecrypted
	}
}

func WriteCredentialsFile(apiUser string, apiKey string, organization string, environment string) {
	var credentials config.CredentialFile
	credentials.APIUser = encryption.EncryptData(apiUser)
	credentials.APIKey = encryption.EncryptData(apiKey)
	credentials.Organization = encryption.EncryptData(organization)
	credentials.Environment = environment
	configFilePath := filepath.Dir(config.Path(true))
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(configFilePath, 0755)
		if err != nil {
			logging.Console().Error().Err(err).Msg("failed to create folder structure for credentials file")
			os.Exit(1)
		}
	}
	credentialsFile, err := json.MarshalIndent(credentials, "", "    ")
	if err != nil {
		logging.Console().Error().Err(err).Msg("failed to marshal credentials")
		os.Exit(1)
	}
	err = os.WriteFile(config.Path(true), credentialsFile, 0644)
	if err != nil {
		logging.Console().Error().Err(err).Msg("failed to write credentials file to " + config.Path(true))
		os.Exit(1)
	} else {
		fmt.Println("Credentials file saved to: " + config.Path(true))
		os.Exit(0)
	}
}
