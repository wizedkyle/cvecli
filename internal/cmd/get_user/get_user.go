package get_user

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
)

func NewCmdGetUser(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	var (
		username string
	)
	cmd := &cobra.Command{
		Use:   "get-user",
		Short: "Retrieves information about a user in the organization",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			getUser(client, username, jsonOutput)
		},
	}
	cmd.Flags().StringVarP(&username, "username", "u", "", "Specify the username of the user to retrieve.")
	cmd.MarkFlagRequired("username")
	return cmd
}

func getUser(client *cveservices_go_sdk.APIClient, username string, jsonOutput *bool) {
	data, response, err := client.GetUser(username)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if !*jsonOutput {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "FIRST NAME\tLAST NAME\tUSERNAME\tUUID\tACTIVE")
			fmt.Fprintln(writer, data.Name.First+"\t"+data.Name.Last+"\t"+data.Username+"\t"+data.UUID+"\t"+strconv.FormatBool(data.Active))
			writer.Flush()
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
}
