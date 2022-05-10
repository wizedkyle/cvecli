package reserve_cve_id

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	"github.com/wizedkyle/cvecli/internal/logging"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
)

func NewCmdReserveCveId(client *cveservices_go_sdk.APIClient, debug *bool, jsonOutput *bool) *cobra.Command {
	var (
		amount        int32
		cveYear       int32
		nonSequential bool
		sequential    bool
	)
	cmd := &cobra.Command{
		Use:   "reserve-cve-id",
		Short: "Reserves a CVE ID for the organization",
		PreRun: func(cmd *cobra.Command, args []string) {
			logging.SetLoggingLevel(debug)
		},
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			reserveCveId(client, amount, cveYear, nonSequential, sequential, jsonOutput, debug)
		},
	}
	cmd.Flags().Int32VarP(&amount, "amount", "a", 1, "The number of CVE IDs to reserve")
	cmd.Flags().Int32VarP(&cveYear, "cve-year", "y", 0, "The year the CVE IDs will be reserved for. If this is not set it will default to the current year")
	cmd.Flags().BoolVarP(&nonSequential, "non-sequential", "n", false, "If the amount of CVE IDs is greater than 1 "+
		"the IDs will be assigned non sequentially")
	cmd.Flags().BoolVarP(&sequential, "sequential", "s", false, "If the amount of CVE IDs is greater than 1 "+
		"the IDs will be assigned sequentially")
	return cmd
}

func reserveCveId(client *cveservices_go_sdk.APIClient, amount int32, cveYear int32, nonSequential bool, sequential bool, jsonOutput *bool, debug *bool) {
	var options types.ReserveCveIdOpts
	if amount > 1 && !nonSequential && !sequential {
		logging.Console().Fatal().Msg("When amount is greater than 1 please specify either non-sequential or sequential.")
	}
	if amount > 1 {
		if nonSequential && sequential {
			logging.Console().Fatal().Msg("Please select either nonSequential or sequential and try again")
		} else if nonSequential {
			options.BatchType = optional.NewString("non-sequential")
		} else if sequential {
			options.BatchType = optional.NewString("sequential")
		}
	}
	if cveYear == 0 {
		cveYear = int32(time.Now().Year())
	}
	data, response, err := client.ReserveCveId(amount, cveYear, &options)
	if *debug {
		logging.DebugHttpResponse(response)
	}
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if !*jsonOutput {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "CVE ID\tCVE YEAR\tSTATE\tOWNING CNA\tREQUESTED BY\tRESERVED DATE")
			for i := 0; i < len(*data.CveIds); i++ {
				fmt.Fprintln(writer, (*data.CveIds)[i].CveId+"\t"+(*data.CveIds)[i].CveYear+"\t"+(*data.CveIds)[i].State+
					"\t"+(*data.CveIds)[i].OwningCna+"\t"+(*data.CveIds)[i].RequestedBy.User+"\t"+(*data.CveIds)[i].Reserved.String())
			}
			writer.Flush()
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
}
