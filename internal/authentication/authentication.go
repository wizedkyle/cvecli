package authentication

import (
	"github.com/spf13/viper"
	"github.com/wizedkyle/cvecli/config"
	"github.com/wizedkyle/cvecli/internal/encryption"
	"github.com/wizedkyle/cvecli/internal/logging"
	"github.com/wizedkyle/cveservices-go-sdk"
	"net/http"
	"path/filepath"
	"time"
)

func CheckCredentialsPath() bool {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(config.Path(true, false)))
	err := viper.ReadInConfig()
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetCVEServicesSDKConfig() *cveservices_go_sdk.APIClient {
	apiUser, apiKey, organization := ReadAuthCredentials()
	client := cveservices_go_sdk.APIClient{
		Cfg: &cveservices_go_sdk.Configuration{
			Authentication: cveservices_go_sdk.BasicAuth{
				APIUser: apiUser,
				APIKey:  apiKey,
			},
			BasePath:     getCveServicesEnvironment(),
			Organization: organization,
			UserAgent:    "cvecli",
			HTTPClient: &http.Client{
				Timeout: time.Second * 20,
			},
		},
	}
	return &client
}

func ReadAuthCredentials() (string, string, string) {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(config.Path(true, false)))
	err := viper.ReadInConfig()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to read credentials file located at " + config.Path(true, false))
		return "", "", ""
	} else {
		apiUser := viper.GetString("apiUser")
		apiKey := viper.GetString("apiKey")
		organization := viper.GetString("organization")
		apiUserDecrypted := encryption.DecryptData(apiUser)
		apiKeyDecrypted := encryption.DecryptData(apiKey)
		organizationDecrypted := encryption.DecryptData(organization)
		return apiUserDecrypted, apiKeyDecrypted, organizationDecrypted
	}
}

func getCveServicesEnvironment() string {
	if config.ProductionEnvironment == true {
		return config.CveServicesProdUrl
	} else {
		return config.CveServicesDevUrl
	}
}
