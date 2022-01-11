package authentication

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wizedkyle/cvecli/config"
	"github.com/wizedkyle/cvecli/internal/build"
	"github.com/wizedkyle/cvecli/internal/encryption"
	"github.com/wizedkyle/cveservices-go-sdk"
	"net/http"
	"os"
	"path/filepath"
	"time"
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

func ReadAuthCredentials() (string, string, string, string) {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(config.Path(true, false)))
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
