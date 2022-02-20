package check_id_quota

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
)

func NewCmdCheckIdQuota(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	var (
		available     bool
		quota         bool
		totalReserved bool
	)
	cmd := &cobra.Command{
		Use:   "check-id-quota",
		Short: "Checks the CVE ID quotas for the organization",
		Long:  "By default the command returns all values however, you can choose to return just the available, quota or totalReserved values.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			checkIdQuota(client, available, quota, totalReserved, jsonOutput)
		},
	}
	cmd.Flags().BoolVarP(&available, "available", "a", false, "Returns the available CVE IDs for the CNA")
	cmd.Flags().BoolVarP(&quota, "quota", "q", false, "Returns the quota of CVE IDs for the CNA")
	cmd.Flags().BoolVarP(&totalReserved, "total-reserved", "t", false, "Returns the total number of reserved CVE IDs for the CNA")
	return cmd
}

func checkIdQuota(client *cveservices_go_sdk.APIClient, available bool, quota bool, totalReserved bool, jsonOutput *bool) {
	data, response, err := client.CheckIdQuota()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if !*jsonOutput {
			if available == true {
				fmt.Println(data.Available)
			} else if quota == true {
				fmt.Println(data.IdQuota)
			} else if totalReserved == true {
				fmt.Println(data.TotalReserved)
			} else {
				fmt.Println("Available: " + strconv.Itoa(int(data.Available)))
				fmt.Println("ID Quota: " + strconv.Itoa(int(data.IdQuota)))
				fmt.Println("Total Reserved: " + strconv.Itoa(int(data.TotalReserved)))
			}
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
}
