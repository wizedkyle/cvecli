package reserve_cve_id

import (
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	"github.com/wizedkyle/cvecli/internal/logging"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"time"
)

func NewCmdReserveCveId(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var (
		amount        int32
		cveIdOutput   bool
		cveYear       int32
		nonSequential bool
		sequential    bool
	)
	cmd := &cobra.Command{
		Use:   "reserve-cve-id",
		Short: "Reserves a CVE ID for the organization.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			reserveCveId(client, amount, cveYear, nonSequential, sequential, cveIdOutput)
		},
	}
	cmd.Flags().Int32VarP(&amount, "amount", "a", 1, "The number of CVE IDs to reserve.")
	cmd.Flags().BoolVarP(&cveIdOutput, "cve-id-output", "o", false, "Outputs only the CVE IDs.")
	cmd.Flags().Int32VarP(&cveYear, "cve-year", "y", 0, "The year the CVE IDs will be reserved for. If this is not set it will default to the current year.")
	cmd.Flags().BoolVarP(&nonSequential, "non-sequential", "n", false, "If the amount of CVE IDs is greater than 1 "+
		"the IDs will be assigned non sequentially.")
	cmd.Flags().BoolVarP(&sequential, "sequential", "s", false, "If the amount of CVE IDs is greater than 1 "+
		"the IDs will be assigned sequentially.")
	return cmd
}

func reserveCveId(client *cveservices_go_sdk.APIClient, amount int32, cveYear int32, nonSequential bool, sequential bool, cveIdOutput bool) {
	var options types.ReserveCveIdOpts
	if amount > 1 {
		if nonSequential == true && sequential == true {
			logging.ConsoleLogger().Error().Msg("Please select either nonSequential or sequential and try again")
		} else if nonSequential == true {
			options.BatchType = optional.NewString("non-sequential")
		} else if sequential == true {
			options.BatchType = optional.NewString("sequential")
		}
	}
	if cveYear == 0 {
		fmt.Println(cveYear)
		cveYear = int32(time.Now().Year())
		fmt.Println(cveYear)
	}
	data, response, err := client.ReserveCveId(amount, cveYear, &options)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if cveIdOutput == true {
			for _, cveId := range *data.CveIds {
				fmt.Println(cveId.CveId)
			}
		} else {
			outputData, err := json.MarshalIndent(data, "", "    ")
			if err != nil {
				logging.ConsoleLogger().Error().Err(err).Msg("failed to marshal response data")
			}
			fmt.Println(string(outputData))
		}
	}

}
