package main

import (
	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
	"github.com/wizedkyle/cvesub/config"
	"github.com/wizedkyle/cvesub/internal/authentication"
	configureCmd "github.com/wizedkyle/cvesub/internal/cmd/configure"
	"github.com/wizedkyle/cvesub/internal/cmd/root"
	"os"
)

func main() {
	if authentication.CheckCredentialsPath() == false {
		pterm.Warning.Println("There are no credentials set or accessible.")
		promptAccept := promptui.Prompt{
			Label:     "Would you like to set credentials now",
			IsConfirm: true,
		}
		promptResponse, err := promptAccept.Run()
		if err != nil {
			os.Exit(0)
		}
		if promptResponse == "y" {
			configureCmd.SetCredentials()
			client := authentication.GetCVEServicesSDKConfig()
			config.SetClient(client)
		}
	} else {
		client := authentication.GetCVEServicesSDKConfig()
		config.SetClient(client)
	}
	rootCmd := root.NewCmdRoot()
	rootCmd.Execute()
}
