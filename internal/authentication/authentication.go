package authentication

import (
	"github.com/spf13/viper"
	"github.com/wizedkyle/cvesub/internal/encryption"
	"github.com/wizedkyle/cvesub/internal/logging"
	"os"
	"path/filepath"
)

func ConfigPath() string {
	var filePath = ".cvesub/credentials/creds.json"
	homeDirectory, _ := os.UserHomeDir()
	configFile := filepath.Join(homeDirectory, filePath)
	return configFile
}

func ReadAuthCredentials() (string, string, string) {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(ConfigPath()))
	err := viper.ReadInConfig()
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to read credentials file located at " + ConfigPath())
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
