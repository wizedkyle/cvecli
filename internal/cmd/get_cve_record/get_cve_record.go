package get_cve_id_record

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
)

func NewCmdGetCveRecord(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	var (
		cveId string
	)
	cmd := &cobra.Command{
		Use:   "get-cve-record",
		Short: "Retrieves the CVE record of the provided CVE ID",
		Long: "The non JSON view only provides a small subset of information. If you would like to see the entire " +
			"CVE record use the --json flag.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getCveRecord(client, cveId, jsonOutput)
		},
	}
	cmd.Flags().StringVarP(&cveId, "cve-id", "c", "", "Specify the CVE ID to retrieve")
	cmd.MarkFlagRequired("cve-id")
	return cmd
}

func getCveRecord(client *cveservices_go_sdk.APIClient, cveId string, jsonOutput *bool) {
	data, response, err := client.GetCveRecord(cveId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if !*jsonOutput {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "CVE ID\tSTATE\tASSIGNER")
			fmt.Fprintln(writer, data.CveMetadata.CveId+"\t"+data.CveMetadata.State+"\t"+data.CveMetadata.Assigner)
			writer.Flush()
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
}
