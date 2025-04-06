package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
		Long:  "This is a sample command for sample purposes:w http.ResponseWriter, r *http.Request",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("sample command...")
		},
	}

	rootCmd.AddCommand(sampleCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
