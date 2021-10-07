package config

import (
	cveservices_go_sdk "github.com/wizedkyle/cveservices-go-sdk"
	"github.com/wizedkyle/cvesub/internal/authentication"
	"net/http"
	"time"
)

type CredentialFile struct {
	APIUser      string `json:"apiUser"`
	APIKey       string `json:"apiKey"`
	Organization string `json:"organization"`
}

var cveUrl = "https://cveawg.mitre.org/api"

func GetCVEServicesSDKConfig() *cveservices_go_sdk.APIClient {
	apiUser, apiKey, organization := authentication.ReadAuthCredentials()
	client := cveservices_go_sdk.APIClient{
		Cfg: &cveservices_go_sdk.Configuration{
			Authentication: cveservices_go_sdk.BasicAuth{
				APIUser: apiUser,
				APIKey:  apiKey,
			},
			BasePath:     "https://cveawg-test.mitre.org/api",
			Organization: organization,
			UserAgent:    "cvesub",
			HTTPClient: &http.Client{
				Timeout: time.Second * 20,
			},
		},
	}
	return &client
}
