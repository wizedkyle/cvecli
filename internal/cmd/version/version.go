package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/cvecli/internal/build"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Displays cvecli version",
		Long:  "Displays the version number, build number, and date the binary was created.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cvecli", build.Version, build.Build)
		},
	}
	return cmd
}
