package root

import (
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/build"
	NewCmdCheckIdQuota "github.com/wizedkyle/cvecli/internal/cmd/check_id_quota"
	configureCmd "github.com/wizedkyle/cvecli/internal/cmd/configure"
	NewCmdCreateUser "github.com/wizedkyle/cvecli/internal/cmd/create_user"
	NewCmdGetOrganizationInfo "github.com/wizedkyle/cvecli/internal/cmd/get_organization_info"
	NewCmdGetUser "github.com/wizedkyle/cvecli/internal/cmd/get_user"
	NewCmdListCveIds "github.com/wizedkyle/cvecli/internal/cmd/list_cve_ids"
	NewCmdListUsers "github.com/wizedkyle/cvecli/internal/cmd/list_users"
	NewCmdReserveCveId "github.com/wizedkyle/cvecli/internal/cmd/reserve_cve_id"
	NewCmdResetSecret "github.com/wizedkyle/cvecli/internal/cmd/reset_secret"
	NewCmdUpdateUser "github.com/wizedkyle/cvecli/internal/cmd/update_user"
)

var (
	jsonOutput bool
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cvecli",
		Short: "CVE Submission CLI",
		Long: "A CLI tool that allows CNAs to manage their organisation and CVEs. " +
			"This tool currently supports the ID Reservation Service.",
		TraverseChildren: true,
		Version:          build.GetVersion(),
	}
	client := authentication.GetCVEServicesSDKConfig()
	cmd.AddCommand(configureCmd.NewCmdConfigure())
	cmd.AddCommand(NewCmdCreateUser.NewCmdCreateUser(client, &jsonOutput))
	cmd.AddCommand(NewCmdGetOrganizationInfo.NewCmdGetOrganizationInfo(client, &jsonOutput))
	cmd.AddCommand(NewCmdGetUser.NewCmdGetUser(client, &jsonOutput))
	cmd.AddCommand(NewCmdCheckIdQuota.NewCmdCheckIdQuota(client, &jsonOutput))
	cmd.AddCommand(NewCmdListCveIds.NewCmdListCveIds(client, &jsonOutput))
	cmd.AddCommand(NewCmdListUsers.NewCmdListUsers(client, &jsonOutput))
	cmd.AddCommand(NewCmdReserveCveId.NewCmdReserveCveId(client, &jsonOutput))
	cmd.AddCommand(NewCmdResetSecret.NewCmdResetSecret(client, &jsonOutput))
	cmd.AddCommand(NewCmdUpdateUser.NewCmdUpdateUser(client, &jsonOutput))
	cmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "p", false, "Outputs the response in json")
	return cmd
}
