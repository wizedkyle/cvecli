package authentication

import (
	"github.com/spf13/viper"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cvesub/config"
	"github.com/wizedkyle/cvesub/internal/encryption"
	"github.com/wizedkyle/cvesub/internal/logging"
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
			// TODO: Need to add support to choose which environment
			BasePath:     "https://cveawg-test.mitre.org/api",
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

func ReadGitHubCredentials(username bool, pat bool) string {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(config.Path(true, false)))
	err := viper.ReadInConfig()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to read credentials file located at " + config.Path(true, false))
		return ""
	} else if username == true {
		githubUsername := viper.GetString("githubUsername")
		githubUsernameDecrypted := encryption.DecryptData(githubUsername)
		return githubUsernameDecrypted
	} else if pat == true {
		githubPat := viper.GetString("githubPat")
		githubPatDecrypted := encryption.DecryptData(githubPat)
		return githubPatDecrypted
	}
	return ""
}
