package update_cve_id_record

import (
	"encoding/json"
	"fmt"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
)

func NewCmdUpdateCveRecord(client *cveservices_go_sdk.APIClient, jsonOutput *bool) *cobra.Command {
	var (
		cveId string
		org   string
		path  string
		state string
	)
	cmd := &cobra.Command{
		Use:   "update-cve-record",
		Short: "Updates the specified CVE record",
		Run: func(cmd *cobra.Command, args []string) {
			updateCveRecord(client, cveId, jsonOutput, org, path, state)
		},
		Hidden: true,
	}
	cmd.Flags().StringVarP(&cveId, "cve-id", "c", "", "Specify the CVE ID to update")
	cmd.Flags().StringVarP(&path, "path", "p", "", "Specify the path of the CVE json file")
	cmd.MarkFlagRequired("cve-id")
	cmd.MarkFlagRequired("path")
	return cmd
}

func updateCveRecord(client *cveservices_go_sdk.APIClient, cveId string, jsonOutput *bool, org string, path string, state string) {
	var (
		body    types.CveJson5
		options types.UpdateCveIdRecordOpts
	)
	options.Org = optional.NewString("")
	options.State = optional.NewString("")
	fileData := cmdutils.ReadFile(path)
	json.Unmarshal(fileData, &body)
	data, response, err := client.UpdateCveRecord(body, cveId, &options)
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
