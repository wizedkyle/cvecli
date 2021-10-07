package check_id_quota

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cveservices-go-sdk"
)

func NewCmdCheckIdQuota(client *cveservices_go_sdk.APIClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check-id-quota",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			checkIdQuota(client)
		},
	}
	return cmd
}

func checkIdQuota(client *cveservices_go_sdk.APIClient) {
	data, _, err := client.CheckIdQuota()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
