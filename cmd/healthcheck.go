/*
Copyright © 2026 Md Mezbah Uddin extraordinarymisbah@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// healthcheckCmd represents the healthcheck command
var healthcheckCmd = &cobra.Command{
	Use:     "healthcheck",
	Aliases: []string{"hc", "check"}, // Short hands
	Short:   "Checks whether the current directory is a git repository",
	Long: `Run this command to check whether the current directory is a git repository. 
			It will return "✅ Git repository detected." if it is a git repository. Otherwise, it will return "❌ Not a git repository."`,

	Run: func(cmd *cobra.Command, args []string) {
		if !IsGitRepo() {
			fmt.Println("❌ Not a git repository.")
			os.Exit(1)
		}
		fmt.Println("✅ Git repository detected. You can now run other commands.")
	},
}

func init() {
	rootCmd.AddCommand(healthcheckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// healthcheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// healthcheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
