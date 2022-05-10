package create_cve_id_record

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
)

func NewCmdCreateCveRecord(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	var (
		path string
	)
	cmd := &cobra.Command{
		Use:   "create-cve-record",
		Short: "Creates a new CVE record",
		Long:  "To create the CVE json file you can use the cvecli generate-cve-record.",
		Run: func(cmd *cobra.Command, args []string) {
			authentication.ConfirmCredentialsSet(client)
			createCveRecord(client, path, jsonOutput)
		},
	}
	cmd.Flags().StringVarP(&path, "path", "p", "", "Specify the path of the CVE json file")
	cmd.MarkFlagRequired("path")
	return cmd
}

func createCveRecord(client *cveservices_go_sdk.APIClient, path string, jsonOutput *bool) {
	var body types.CveJson5
	fileData := cmdutils.ReadFile(path)
	json.Unmarshal(fileData, &body)
	data, response, err := client.CreateCveRecord(body, body.CveMetadata.CveId)
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		if !*jsonOutput {
			fmt.Println(data.Message)
		} else {
			fmt.Println(string(cmdutils.OutputJson(data)))
		}
	}
}
