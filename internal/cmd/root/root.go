package root

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvesub/internal/authentication"
	NewCmdCheckIdQuota "github.com/wizedkyle/cvesub/internal/cmd/check_id_quota"
	configureCmd "github.com/wizedkyle/cvesub/internal/cmd/configure"
	NewCmdCreateCveEntry "github.com/wizedkyle/cvesub/internal/cmd/create_cve_entry"
	NewCmdCreateUser "github.com/wizedkyle/cvesub/internal/cmd/create_user"
	NewCmdGetOrganizationInfo "github.com/wizedkyle/cvesub/internal/cmd/get_organization_info"
	NewCmdGetUser "github.com/wizedkyle/cvesub/internal/cmd/get_user"
	NewCmdListCveIds "github.com/wizedkyle/cvesub/internal/cmd/list_cve_ids"
	NewCmdListUsers "github.com/wizedkyle/cvesub/internal/cmd/list_users"
	NewCmdReserveCveId "github.com/wizedkyle/cvesub/internal/cmd/reserve_cve_id"
	NewCmdResetSecret "github.com/wizedkyle/cvesub/internal/cmd/reset_secret"
	NewCmdUpdateUser "github.com/wizedkyle/cvesub/internal/cmd/update_user"
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
	cmd.AddCommand(NewCmdCreateUser.NewCmdCreateUser(client))
	cmd.AddCommand(NewCmdCreateCveEntry.NewCmdCreateCveEntry(client))
	cmd.AddCommand(NewCmdGetOrganizationInfo.NewCmdGetOrganizationInfo(client))
	cmd.AddCommand(NewCmdGetUser.NewCmdGetUser(client))
	cmd.AddCommand(NewCmdCheckIdQuota.NewCmdCheckIdQuota(client))
	cmd.AddCommand(NewCmdListCveIds.NewCmdListCveIds(client))
	cmd.AddCommand(NewCmdListUsers.NewCmdListUsers(client))
	cmd.AddCommand(NewCmdReserveCveId.NewCmdReserveCveId(client))
	cmd.AddCommand(NewCmdResetSecret.NewCmdResetSecret(client))
	cmd.AddCommand(NewCmdUpdateUser.NewCmdUpdateUser(client))
	return cmd
}
