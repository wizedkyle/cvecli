package root

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvesub/internal/authentication"
	NewCmdCheckIdQuota "github.com/wizedkyle/cvesub/internal/cmd/check_id_quota"
	configureCmd "github.com/wizedkyle/cvesub/internal/cmd/configure"
	NewCmdCreateCveEntry "github.com/wizedkyle/cvesub/internal/cmd/create_cve_entry"
	NewCmdGetUser "github.com/wizedkyle/cvesub/internal/cmd/get_user"
	NewCmdReserveCveId "github.com/wizedkyle/cvesub/internal/cmd/reserve_cve_id"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cvesub",
		Short: "CVE Submission CLI",
		Long: "A CLI tool that allows easy submissions of CVEs to MITREs GitHub repo (for CNAs)." +
			"This tool currently supports the ID Reservation Service.",
		TraverseChildren: true,
	}
	client := authentication.GetCVEServicesSDKConfig()
	cmd.AddCommand(configureCmd.NewCmdConfigure())
	cmd.AddCommand(NewCmdCreateCveEntry.NewCmdCreateCveEntry(client))
	cmd.AddCommand(NewCmdGetUser.NewCmdGetUser(client))
	cmd.AddCommand(NewCmdCheckIdQuota.NewCmdCheckIdQuota(client))
	cmd.AddCommand(NewCmdReserveCveId.NewCmdReserveCveId(client))
	return cmd
}
