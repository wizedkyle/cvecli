package main

import (
	"github.com/wizedkyle/cvecli/internal/cmd/root"
)

func main() {
	rootCmd := root.NewCmdRoot()
	rootCmd.Execute()
}
