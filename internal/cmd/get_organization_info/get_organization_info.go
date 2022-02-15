package get_organization_info

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	"github.com/wizedkyle/cveservices-go-sdk"
	"os"
	"strconv"
	"text/tabwriter"
)

func NewCmdGetOrganizationInfo(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-organization-info",
		Short: "Retrieves information about the organization the user authenticating is apart of",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getOrganizationInfo(client, jsonOutput)
		},
	}
	return cmd
}

func getOrganizationInfo(client *cveservices_go_sdk.APIClient, jsonOutput *bool) {
	data, response, err := client.GetOrganizationRecord()
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if *jsonOutput == false {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "NAME\tSHORT NAME\tUUID\tID QUOTA")
			fmt.Fprintln(writer, data.Name+"\t"+data.ShortName+"\t"+data.UUID+"\t"+strconv.Itoa(int(data.Policies.IdQuota)))
			writer.Flush()
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
}
