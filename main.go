package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool
var version = "v0.1.0"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "envr",
		Short: "Environment Readiness Checker",
		Long:  "Environment Readiness Checker (envr) helps you validate Linux environments for Go, Kubernetes, and vCluster development.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("envr: Environment Readiness Checker")
		},
	}

	var sampleCmd = &cobra.Command{
		Use:   "sample",
		Short: "Sample Command",
		Long:  "This is a sample command for sample purposes",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("sample command...")
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version Command",
		Long:  "Current version for environment checker",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Environment Readiness Checker: %s\n", version)
		},
	}

	rootCmd.AddCommand(sampleCmd)
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
