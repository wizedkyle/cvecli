package reserve_cve_id

import (
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
	"github.com/wizedkyle/cvesub/internal/cmdutils"
	"github.com/wizedkyle/cvesub/internal/logging"
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
			reserveCveId(client, amount, cveYear, nonSequential, sequential, cveIdOutput)
		},
	}
	cmd.Flags().Int32Var(&amount, "amount", 0, "The number of CVE IDs to reserve.")
	cmd.Flags().BoolVar(&cveIdOutput, "cveIdOutput", false, "Outputs only the CVE IDs.")
	cmd.Flags().Int32Var(&cveYear, "cveYear", 0, "The year the CVE IDs will be reserved for.")
	cmd.Flags().BoolVar(&nonSequential, "nonSequential", false, "If the amount of CVE IDs is greater than 1 "+
		"the IDs will be assigned non sequentially.")
	cmd.Flags().BoolVar(&sequential, "sequential", false, "If the amount of CVE IDs is greater than 1 "+
		"the IDs will be assigned sequentially.")
	cmd.MarkFlagRequired("amount")
	cmd.MarkFlagRequired("cveYear")
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
