package root

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvesub/config"
	NewCmdCheckIdQuota "github.com/wizedkyle/cvesub/internal/cmd/check_id_quota"
	configureCmd "github.com/wizedkyle/cvesub/internal/cmd/configure"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cvesub",
		Short: "CVE Submission CLI",
		Long: "A CLI tool that allows easy submissions of CVEs to MITREs GitHub repo (for CNAs)." +
			"This tool currently supports the ID Reservation Service.",
		TraverseChildren: true,
	}
	client := config.GetCVEServicesSDKConfig()
	cmd.AddCommand(configureCmd.NewCmdConfigure())
	cmd.AddCommand(NewCmdCheckIdQuota.NewCmdCheckIdQuota(client))
	return cmd
}
