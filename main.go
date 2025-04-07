package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "version",
		Long:  "print Environment Readiness Checker version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("envr: %s\n",  version)
		},
	}

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
