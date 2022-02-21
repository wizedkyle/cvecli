package get_cve_id

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
	"os"
	"text/tabwriter"
)

func NewCmdGetCveId(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	var (
		cveId string
	)
	cmd := &cobra.Command{
		Use:   "get-cve-id",
		Short: "Retrieves a CVE ID record by the ID",
		Long:  "A CVE ID can be retrieved from a different entity if it is in either a REJECTED or PUBLISHED state.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getCveId(client, cveId, jsonOutput)
		},
	}
	cmd.Flags().StringVarP(&cveId, "cve-id", "c", "", "Specify the CVE ID to retrieve")
	cmd.MarkFlagRequired("cve-id")
	return cmd
}

func getCveId(client *cveservices_go_sdk.APIClient, cveId string, jsonOutput *bool) {
	data, response, err := client.GetCveId(cveId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if !*jsonOutput {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "CVE ID\tCVE YEAR\tSTATE\tOWNING CNA\tRESERVED DATE")
			fmt.Fprintln(writer, data.CveId+"\t"+data.CveYear+"\t"+data.State+"\t"+data.OwningCNA+"\t"+data.Reserved.String())
			writer.Flush()
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
}
