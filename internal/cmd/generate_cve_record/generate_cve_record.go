package generate_cve_schema

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/authentication"
	"github.com/wizedkyle/cvecli/internal/cmdutils"
	"github.com/wizedkyle/cvecli/internal/logging"
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cveservices-go-sdk/types"
)

func NewCmdGenerateCveRecord(client *cveservices_go_sdk.APIClient) *cobra.Command {
	var (
		advanced               bool
		basic                  bool
		cveId                  string
		description            string
		getAssignerId          bool
		language               string
		lessThanVersion        string
		lessThanOrEqualVersion string
		path                   string
		problemTypeDescription string
		productName            string
		referenceUrl           string
		version                string
		vendorName             string
		versionType            string
	)
	cmd := &cobra.Command{
		Use:   "generate-cve-record",
		Short: "Generates a CVE record based on the JSON 5 schema",
		Long: "This command will generate and complete the bare minimum requirements for a CVE record entry. You can specify custom values for fields by using the accepted flags. " +
			"Once the CVE record has been generated it is recommended to review the contents prior to submitting using cveli create-cve-id-record.",
		Run: func(cmd *cobra.Command, args []string) {
			generateCveSchema(advanced, basic, cveId, description, getAssignerId, language, lessThanVersion, lessThanOrEqualVersion,
				path, problemTypeDescription, productName, referenceUrl, version, vendorName, versionType, client)
		},
	}
	cmd.Flags().BoolVarP(&advanced, "advanced", "a", false, "Generates a advanced CVE schema")
	cmd.Flags().BoolVarP(&basic, "basic", "b", false, "Generates a basic CVE schema")
	cmd.Flags().StringVarP(&cveId, "cve-id", "c", "", "Specify the CVE ID to update if known")
	cmd.Flags().StringVar(&description, "description", "", "Specify the details of the vulnerability. For example: OS Command Injection vulnerability parseFilename function of example.php in the Web Management Interface of Example.org Example Enterprise on Windows, MacOS and XT-4500 allows remote unauthenticated attackers to escalate privileges.")
	cmd.Flags().BoolVar(&getAssignerId, "getAssignerId", false, "If this flag is set the configured credentials will be used to obtain the organization ID and populate the fields in the CVE record")
	cmd.Flags().StringVar(&language, "language", "", "Specify the language using ")
	cmd.Flags().StringVar(&lessThanVersion, "lessThanVersion", "", "Specify the latest version that isn't affected. This is the same as < 1.0.6")
	cmd.Flags().StringVar(&lessThanOrEqualVersion, "lessThanOrEqualVersion", "", "Specify the latest version that is affected. This is the same as <= 1.0.6")
	cmd.Flags().StringVarP(&path, "path", "p", "", "Specify the path to save the json file")
	cmd.Flags().StringVar(&problemTypeDescription, "problemTypeDescription", "", "Specify the problem type description. For example: Command Injection or Remote Code Execution")
	cmd.Flags().StringVarP(&productName, "product-name", "n", "", "Specify the production name if known")
	cmd.Flags().StringVar(&referenceUrl, "referenceUrl", "", "Specify a URL to a security advisory relating to this CVE")
	cmd.Flags().StringVar(&version, "version", "", "Specify the starting version of your software that is affected. For example: 1.0.0")
	cmd.Flags().StringVarP(&vendorName, "vendor-name", "v", "", "Specify the vendor name if known")
	cmd.Flags().StringVar(&versionType, "versionType", "semver", "Specify the version type. Valid values are: custom, git, maven, python, rpm, semver")
	cmd.MarkFlagRequired("path")
	return cmd
}

func generateCveSchema(advanced bool, basic bool, cveId string, description string, getAssignerId bool, language string, lessThanVersion string, lessThanOrEqualVersion string, path string,
	problemTypeDescription string, productName string, referenceUrl string, version string, vendorName string, versionType string, client *cveservices_go_sdk.APIClient) {
	var (
		affected                 = types.CveJson5Affected{}
		affectedVersion          = types.CveJson5Versions{}
		basicJsonSchema          = types.CveJson5{}
		descriptionsArray        = []types.CveJson5Descriptions{}
		descriptions             = types.CveJson5Descriptions{}
		problemTypesArray        = []types.CveJson5ProblemTypes{}
		problemTypes             = types.CveJson5ProblemTypes{}
		problemTypesDescriptions = types.CveJson5ProblemTypeDescriptions{}
		referencesArray          = []types.CveJson5References{}
		references               = types.CveJson5References{}
	)
	filePathDirectory := filepath.Dir(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(filePathDirectory, 0755)
		if err != nil {
			logging.ConsoleLogger().Error().Err(err).Msg("failed to create folder structure for CVE record file")
			os.Exit(1)
		}
	}
	// Validate input
	validateLanguage(language)
	if referenceUrl != "" {
		validateUrl(referenceUrl)
	}
	if !validateVersionType(versionType) {
		fmt.Println("Please enter a valid version type: custom, git, maven, python, rpm, semver")
		os.Exit(1)
	}

	// Build Basic Schema
	// Assign metadata values
	basicJsonSchema.DataVersion = "5.0"
	basicJsonSchema.DataType = "CVE_RECORD"
	basicJsonSchema.CveMetadata.State = "PUBLISHED"
	if !getAssignerId {
		basicJsonSchema.CveMetadata.Assigner = "00000000-0000-0000-0000-000000000000"
		basicJsonSchema.Containers.Cna.ProviderMetadata.Id = "00000000-0000-0000-0000-000000000000"
	} else {
		authentication.ConfirmCredentialsSet(client)
		data, response, err := client.GetOrganizationRecord()
		if err != nil {
			cmdutils.OutputError(response, err)
		} else {
			basicJsonSchema.CveMetadata.Assigner = data.UUID
			basicJsonSchema.Containers.Cna.ProviderMetadata.Id = data.UUID
		}
	}
	if cveId != "" {
		basicJsonSchema.CveMetadata.CveId = cveId
	} else {
		basicJsonSchema.CveMetadata.CveId = "CVE-XXXX-XXXXX"
	}

	// Assign problem description values
	if problemTypeDescription != "" && language != "" {
		problemTypesDescriptions.Description = problemTypeDescription
		problemTypesDescriptions.Lang = language
	} else {
		problemTypesDescriptions.Description = "Problem description placeholder"
		problemTypesDescriptions.Lang = "eng"
	}
	problemTypes.Descriptions = append(problemTypes.Descriptions, problemTypesDescriptions)
	problemTypesArray = append(problemTypesArray, problemTypes)

	// Assign affected values
	if vendorName != "" {
		affected.Vendor = vendorName
	} else {
		affected.Vendor = "Vendor name placeholder"
	}
	if productName != "" {
		affected.Product = productName
	} else {
		affected.Product = "Product name placeholder"
	}
	if version != "" {
		affectedVersion.Version = version
	}
	if lessThanVersion != "" && lessThanOrEqualVersion != "" {
		fmt.Println("Please select either lessThanVersion or lessThanOrEqualVersion")
		os.Exit(1)
	} else if lessThanVersion != "" {
		affectedVersion.LessThan = lessThanVersion
	} else if lessThanOrEqualVersion != "" {
		affectedVersion.LessThanOrEqual = lessThanOrEqualVersion
	}
	affectedVersion.VersionType = versionType
	affectedVersion.Status = "affected"
	affected.DefaultStatus = "unaffected"
	affected.Versions = append(affected.Versions, affectedVersion)

	// Assign descriptions values
	if description != "" {
		descriptions.Value = description
	} else {
		descriptions.Value = "Description placeholder"
	}
	if language != "" {
		descriptions.Lang = language
	} else {
		descriptions.Lang = "eng"
	}
	descriptionsArray = append(descriptionsArray, descriptions)

	// Assign references
	if referenceUrl != "" {
		references.Url = referenceUrl
	} else {
		references.Url = "https://cve.org"
	}
	referencesArray = append(referencesArray, references)

	basicJsonSchema.Containers.Cna.Affected = append(basicJsonSchema.Containers.Cna.Affected, affected)
	basicJsonSchema.Containers.Cna.Descriptions = append(basicJsonSchema.Containers.Cna.Descriptions, descriptionsArray...)
	basicJsonSchema.Containers.Cna.ProblemTypes = append(basicJsonSchema.Containers.Cna.ProblemTypes, problemTypesArray...)
	basicJsonSchema.Containers.Cna.References = append(basicJsonSchema.Containers.Cna.References, referencesArray...)

	// Either export as a basic schema or add the advanced properties
	if basic {
		schemaFile, err := json.MarshalIndent(basicJsonSchema, "", "    ")
		if err != nil {
			logging.ConsoleLogger().Error().Err(err).Msg("failed to marshal basic schema")
			os.Exit(1)
		}
		err = os.WriteFile(path, schemaFile, 0644)
		if err != nil {
			logging.ConsoleLogger().Error().Err(err).Msg("failed to write basic schema file to " + path)
			os.Exit(1)
		} else {
			fmt.Println("Basic schema file saved to: " + path)
			os.Exit(0)
		}
	} else if advanced {

	} else {
		fmt.Println("Please select either --advanced or --basic")
		os.Exit(1)
	}
}

func validateLanguage(language string) {
	if len(language) > 3 {
		fmt.Println("Language value is to long. Maximum length is 3.")
		os.Exit(1)
	}
}

func validateUrl(referenceUrl string) {
	_, err := url.ParseRequestURI(referenceUrl)
	if err != nil {
		fmt.Println("Reference URL is not a valid URL")
		os.Exit(1)
	}
}

func validateVersionType(versionType string) bool {
	switch versionType {
	case
		"custom",
		"git",
		"maven",
		"python",
		"rpm",
		"semver":
		return true
	}
	return false
}
