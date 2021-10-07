package main

import "github.com/wizedkyle/cvesub/internal/cmd/root"

func main() {
	rootCmd := root.NewCmdRoot()
	rootCmd.Execute()
}
