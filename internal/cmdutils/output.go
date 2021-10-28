package cmdutils

import (
	"encoding/json"
	"github.com/wizedkyle/cvesub/internal/logging"
)

func OutputJson(data interface{}) []byte {
	dataJson, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to marshall response")
	}
	return dataJson
}
