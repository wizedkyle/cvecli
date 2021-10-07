package encryption

import (
	"github.com/wizedkyle/cvesub/internal/logging"
	"golang.org/x/sys/windows/registry"
)

func getMachineId() string {
	registryKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to open registry item")
	}
	defer registryKey.Close()
	machineId, _, err := registryKey.GetStringValue("MachineGuid")
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to get MachineGuid")
	}
	return machineId
}
