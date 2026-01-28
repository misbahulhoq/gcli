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
	Use:     "gcli",
	Version: "1.0.0",
	Short:   "An AI based Git Committer CLI tool",
	Long: `
GCLI is an AI powered Git Committer CLI tool. It uses Google's Gemini API to generate commit messages for you.
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
