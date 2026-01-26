/*
Copyright © 2026 Md Mezbah Uddin extraordinarymisbah@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gcm",
	Version: "1.0.0",
	Short:   "An AI based Git Committer CLI tool",
	Long: `The Git Committer CLI is a command-line utility developed in Go designed to streamline the software development workflow. It automates the process of writing git commit messages by analyzing staged changes and leveraging Google's Gemini AI to generate context-aware, meaningful, and conventional commit messages.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
