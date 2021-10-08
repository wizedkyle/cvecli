package check_id_quota

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cvesub/internal/cmdutils"
	"strconv"
)

func NewCmdCheckIdQuota(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var (
		available     bool
		quota         bool
		totalReserved bool
	)
	cmd := &cobra.Command{
		Use:   "check-id-quota",
		Short: "Checks the ID quotas for the organization.",
		Long:  "By default the command returns all values however, you can choose to return just the available, quota or totalReserved values.",
		Run: func(cmd *cobra.Command, args []string) {
			checkIdQuota(client, available, quota, totalReserved)
		},
	}
	cmd.Flags().BoolVar(&available, "available", false, "")
	cmd.Flags().BoolVar(&quota, "quota", false, "")
	cmd.Flags().BoolVar(&totalReserved, "totalReserved", false, "")
	return cmd
}

func checkIdQuota(client *cveservices_go_sdk.APIClient, available bool, quota bool, totalReserved bool) {
	data, response, err := client.CheckIdQuota()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
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
	}
}
